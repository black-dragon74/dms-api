package utils

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func NewRequest(method string, path string, cookies *map[string]string, headers *map[string][]string, body *io.ReadCloser) *http.Request {
	req := &http.Request{}
	req.Method = method
	req.URL = getURL(path)
	req.Header = *headers

	for key, val := range *cookies {
		req.AddCookie(newCookie(key, val))
	}

	if body != nil {
		req.Body = *body
	}

	return req
}

// GetSessionFromResponse returns the value of `SessionCookie` from a http response
func GetSessionFromResponse(resp *http.Response) string {
	if len(resp.Cookies()) == 0 {
		return ""
	} else {
		for _, v := range resp.Cookies() {
			if v.Name == SessionCookie {
				return v.Value
			}
		}
	}

	return ""
}

// ParseArgs reads variables passed as URL query string in the request
func ParseArgs(request *http.Request, values *[]string) map[string]string {
	resp := make(map[string]string)
	query := request.URL.Query()

	for _, v := range *values {
		if tVal := query.Get(v); tVal != "" {
			resp[v] = tVal
		}
	}

	return resp
}

// ValidateArgs checks if each item from `keys` is present in the `store` with a non-zero value
func ValidateArgs(keys *[]string, store *map[string]string) error {
	for _, k := range *keys {
		if (*store)[k] == "" {
			return errors.New(fmt.Sprintf("missing var %s", k))
		}
	}

	return nil
}

// WriteJSONError writes error to the HTTP response as a JSON object
func WriteJSONError(writer http.ResponseWriter, error error) {
	_, _ = writer.Write(errorToJSON(error.Error()))
}

// errorToJSON returns a JSON formatted string with the `msg` as error
func errorToJSON(msg string) []byte {
	return []byte(`{"error":"` + msg + `"}`)
}

func getURL(path string) *url.URL {
	u, _ := url.Parse(path)
	return u
}

func newCookie(name string, val string) *http.Cookie {
	return &http.Cookie{
		Name:  name,
		Value: val,
	}
}
