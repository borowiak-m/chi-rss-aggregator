package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Get APIKey extracts an API key from
// the headers of an HTTP request
// Expected auth header:
// Authorisation: APiKey {insert paikey here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorisation")
	if val == "" {
		return "", errors.New("No authentication value found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("Incorrect authorisation header data")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("Incorrect first part of auth header")
	}
	if len(vals[1]) != 64 {
		return "", errors.New("Incorrect second part of auth header")
	}
	return vals[1], nil
}
