package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// TODO
// Update tests to work with a loop of structs instead of many tests
func TestCreateUrl(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "/", nil)
	check(err)

	query := request.URL.Query()
	query.Add("url", "http://example.com")
	request.URL.RawQuery = query.Encode()

	response := httptest.NewRecorder()

	handler(response, request)

	got := response.Body.String()
	want := "8OamqXB"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestGetUrl(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "/8OamqXB", nil)
	check(err)

	response := httptest.NewRecorder()

	handler(response, request)

	got := response.Body.String()
	want := "http://example.com"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestGetRootMessage(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "/", nil)
	check(err)

	response := httptest.NewRecorder()

	handler(response, request)

	got := response.Body.String()
	want := "404 Not Found"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestGetRootMethod(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "/", nil)
	check(err)

	response := httptest.NewRecorder()

	handler(response, request)

	got := response.Result().StatusCode
	want := 404

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
