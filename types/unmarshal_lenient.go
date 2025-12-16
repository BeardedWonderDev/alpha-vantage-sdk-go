package types

import (
	"bytes"
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
)

// UnmarshalLenient unmarshals JSON into v, tolerating placeholder strings like "n/a"
// for numeric fields by treating them as zero values.
//
// This is primarily to handle Alpha Vantage responses that sometimes emit "n/a"
// for numeric-looking fields that are otherwise encoded as JSON strings.
func UnmarshalLenient(data []byte, v any) error {
	origErr := json.Unmarshal(data, v)
	if origErr == nil {
		return nil
	}

	if !containsNumericNA(data) {
		// Retry logic is only intended for the "n/a" numeric placeholder case; if
		// there is no such placeholder in the payload, return the original error.
		return origErr
	}

	sanitized, ok := sanitizeNumericNA(data, reflect.TypeOf(v))
	if !ok {
		return origErr
	}

	if err := json.Unmarshal(sanitized, v); err != nil {
		return origErr
	}
	return nil
}

func containsNumericNA(data []byte) bool {
	// Fast pre-check to avoid reflection + decode cost on normal payloads.
	lower := bytes.ToLower(data)
	return bytes.Contains(lower, []byte(`"n/a"`)) || bytes.Contains(lower, []byte(`"na"`))
}

type jsonFieldInfo struct {
	name         string
	stringEncode bool
	hasTag       bool
	skip         bool
}

func parseJSONFieldInfo(f reflect.StructField) jsonFieldInfo {
	tag, ok := f.Tag.Lookup("json")
	if !ok {
		return jsonFieldInfo{name: f.Name, hasTag: false}
	}
	if tag == "-" {
		return jsonFieldInfo{skip: true, hasTag: true}
	}

	parts := strings.Split(tag, ",")
	name := parts[0]
	if name == "" {
		name = f.Name
	}

	stringEncode := false
	for _, opt := range parts[1:] {
		if opt == "string" {
			stringEncode = true
			break
		}
	}

	return jsonFieldInfo{name: name, stringEncode: stringEncode, hasTag: true}
}

func sanitizeNumericNA(data []byte, t reflect.Type) ([]byte, bool) {
	if t == nil {
		return nil, false
	}

	var root any
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	if err := dec.Decode(&root); err != nil {
		return nil, false
	}

	root = sanitizeValue(root, t, false)

	out, err := json.Marshal(root)
	if err != nil {
		return nil, false
	}

	return out, true
}

func sanitizeValue(v any, t reflect.Type, stringEncoded bool) any {
	if t == nil {
		return v
	}

	for t.Kind() == reflect.Pointer {
		if v == nil {
			return nil
		}
		t = t.Elem()
	}

	switch t.Kind() {
	case reflect.Struct:
		m, ok := v.(map[string]any)
		if !ok {
			return v
		}
		sanitizeStructMap(m, t)
		return m
	case reflect.Slice, reflect.Array:
		arr, ok := v.([]any)
		if !ok {
			return v
		}
		elem := t.Elem()
		for i := range arr {
			arr[i] = sanitizeValue(arr[i], elem, false)
		}
		return arr
	case reflect.Map:
		// For maps with numeric values, handle "n/a" generically (no tag options exist for map values).
		m, ok := v.(map[string]any)
		if !ok {
			return v
		}
		elem := t.Elem()
		for k, vv := range m {
			m[k] = sanitizeValue(vv, elem, false)
		}
		return m
	case reflect.Float32, reflect.Float64:
		return sanitizeNumericLeaf(v, t.Kind(), stringEncoded)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return sanitizeNumericLeaf(v, t.Kind(), stringEncoded)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return sanitizeNumericLeaf(v, t.Kind(), stringEncoded)
	default:
		return v
	}
}

func sanitizeStructMap(m map[string]any, t reflect.Type) {
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.PkgPath != "" && !f.Anonymous {
			continue
		}

		info := parseJSONFieldInfo(f)
		if info.skip {
			continue
		}

		// Handle embedded structs that are flattened by encoding/json.
		if f.Anonymous && !info.hasTag {
			ft := f.Type
			for ft.Kind() == reflect.Pointer {
				ft = ft.Elem()
			}
			if ft.Kind() == reflect.Struct {
				sanitizeStructMap(m, ft)
				continue
			}
		}

		raw, ok := m[info.name]
		if !ok {
			continue
		}

		m[info.name] = sanitizeValue(raw, f.Type, info.stringEncode)
	}
}

func sanitizeNumericLeaf(v any, kind reflect.Kind, stringEncoded bool) any {
	switch vv := v.(type) {
	case nil:
		if stringEncoded {
			return "0"
		}
		return numericZero(kind)
	case string:
		s := strings.TrimSpace(vv)
		if isNAString(s) {
			if stringEncoded {
				return "0"
			}
			return numericZero(kind)
		}
		if stringEncoded {
			return s
		}
		// Best-effort numeric conversion for non-,string numeric fields that arrive as strings.
		if n, ok := parseNumberStringAsKind(s, kind); ok {
			return n
		}
		return v
	case json.Number:
		if stringEncoded {
			return vv.String()
		}
		return v
	case float64:
		if stringEncoded {
			return strconv.FormatFloat(vv, 'f', -1, 64)
		}
		return v
	default:
		return v
	}
}

func isNAString(s string) bool {
	s = strings.TrimSpace(strings.ToLower(s))
	return s == "n/a" || s == "na"
}

func numericZero(kind reflect.Kind) any {
	switch kind {
	case reflect.Float32, reflect.Float64:
		return float64(0)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return int64(0)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return uint64(0)
	default:
		return float64(0)
	}
}

func parseNumberStringAsKind(s string, kind reflect.Kind) (any, bool) {
	switch kind {
	case reflect.Float32, reflect.Float64:
		f, err := strconv.ParseFloat(s, 64)
		if err != nil {
			return nil, false
		}
		return f, true
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		i, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil, false
		}
		return i, true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		u, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			return nil, false
		}
		return u, true
	default:
		return nil, false
	}
}
