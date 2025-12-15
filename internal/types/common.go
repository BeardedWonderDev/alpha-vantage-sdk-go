package types

import "net/url"

type Client interface {
	Do(string, url.Values) ([]byte, error)
}
