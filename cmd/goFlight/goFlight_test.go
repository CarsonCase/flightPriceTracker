package main

import (
	"net/http"
	"testing"
)

// This passes...
// Main may not because of waiting for api key or improperly formating it?
func TestPostRoute(t *testing.T) {
	want := http.StatusOK
	got, err := postRoute("413519a5d8613e574660a848fbc7c661bf1c72a368721a14823c0fdea72a817f", "JFK", "LAX")
	if err != nil {
		t.Errorf(err.Error())
	}
	if want != got {
		t.Errorf("Expected %v, got %v", want, got)
	}
}
