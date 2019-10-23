package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	var uri, expected string
	uri = "/?hoge=fug"
	b, err := get(uri)
	if err != nil {
		t.Errorf("%v", err)
	}
	expected = "hoge=fuga\n"
	if s := string(b); s != expected {
		t.Errorf("unexpected response: %s", s)
	}
	uri = "/?hoge=fuga&x=1&hoge=FUGA"
	b, err = get(uri)
	if err != nil {
		t.Errorf("%v", err)
	}
	expected = "hoge=FUGA\nhoge=fuga\nx=1\n"
	if s := string(b); s != expected {
		t.Errorf("unexpected response: %s", s)
	}
}

func get(url string) ([]byte, error) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", url, nil)
	handler(w, r)
	rw := w.Result()
	defer rw.Body.Close()
	if rw.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code")
	}
	b, err := ioutil.ReadAll(rw.Body)
	if err != nil {
		return b, fmt.Errorf("unexpected error: %w", err)
	}
	return b, nil
}
