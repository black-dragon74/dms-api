package utils

import (
	"net/http"
	"net/url"
)

func NewRequest(method string, path string, cookies *map[string]string, headers *map[string][]string) *http.Request {
	req := &http.Request{}
	req.Method = method
	req.URL = getURL(path)
	req.Header = *headers

	for key, val := range *cookies {
		req.AddCookie(newCookie(key, val))
	}

	return req
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
