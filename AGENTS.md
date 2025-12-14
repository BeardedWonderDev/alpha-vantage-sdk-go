# AGENTS – alpha-vantage-go-wrapper

Repo-specific operating guide for Codex when working in this Go library.

## Global Rules and Required Reading
- Always read and follow the global rules at `~/.codex/AGENTS.md` before applying any repo-local guidance here.
- This repo is plain Go; also read and follow `rules/Go_General.md`.

## Repository Snapshot
- Language: Go 1.21.1 (`go.mod`).
- Module path: `github.com/masonJamesWheeler/alpha-vantage-go-wrapper`.
- Packages: `client` (API surface) and `models` (data types, JSON parsing, string formatters).
- Dependencies: standard library only; keep it that way unless absolutely necessary.
- No tests or CI config are present yet; additions should be standard Go tooling friendly.

## Folder and Naming Conventions
- `client/`: exported `Client` type plus methods for each Alpha Vantage function. Helper functions: `getTimeSeriesData`, `getCryptoData`, `getIndicator`, `GetIndicatorData`. Keep new methods grouped with peers (time series, crypto, indicators).
- `models/`: structs with JSON tags mirroring Alpha Vantage responses, plus custom `UnmarshalJSON` helpers that normalize map-based payloads into ordered slices and add `String()` formatters. Add new response shapes here, not in `client/`.
- File naming is descriptive by domain (`time_series.go`, `crypto.go`, `indicators.go`). Follow this pattern for new files.
- Keep exports Go-idiomatic: package names lower_case, exported symbols UpperCamelCase matching Alpha Vantage function names (e.g., `GetBBANDS`, `TimeSeriesDailyAdjusted`).

## Architecture and Boundaries
- `Client` only holds the API key; no global state. New functionality should stay instance-bound.
- HTTP behavior: simple `net/http` GET against `alphaVantageURL` with query params assembled via `url.Values`. Reuse this pattern; do not introduce other HTTP clients unless a concrete requirement emerges.
- Parameter structs (`TimeSeriesParams`, `IndicatorParams`, `CryptoParams`, etc.) model Alpha Vantage inputs. Preserve existing optional-field handling (string or pointer/interface) so public API remains backward compatible.
- Parsing: keep custom unmarshalers and post-parse sorting to produce deterministic slices. When adding new responses, prefer map-to-slice normalization plus stable sort.
- Formatting: `String()` methods present table-like output. Extend consistently if you add new data shapes.
- Error handling: propagate API and parse errors; do not swallow or log them here.

## Testing Expectations
- Add Go tests alongside code (`client/client_test.go`, `models/..._test.go`). Use table-driven tests and golden fixtures under `testdata/` instead of live API calls; never require real API keys in tests.
- For new parsing logic, include sample JSON covering success and edge cases (missing fields, numeric parsing). Validate slice ordering and key field values.
- Run `go test ./...` before delivery. Include benchmarks only if useful for parsing changes.

## CI / Tooling
- Run `gofmt` and `go vet` on touched files. If dependencies ever change, run `go mod tidy` and keep `go.sum` in sync.
- No existing CI config; keep additions minimal and Go-native (e.g., `go test ./...`).

## Extending the Library
- New time series/quote variants: add a typed method in `client` that calls `getTimeSeriesData` (or a small helper if a new Alpha Vantage function requires one) and unmarshal into a new/extended model type.
- New indicators: prefer wiring through `getIndicator(indicatorName, params)`; add the indicator-specific method that forwards the correct function name.
- New crypto endpoints: mirror `getCryptoData` pattern; keep market/interval handling consistent.
- Validate new params early (non-empty symbols, intervals) and document required/optional fields in code comments for exported APIs.
- Maintain backwards compatibility of public signatures; if breaking changes are unavoidable, note them prominently in doc comments and README.
