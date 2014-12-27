package main

import (
	"net/http"
	"net/http/httptest"
	//	"reflect"
	"fmt"
	"testing"
)

func TestSearchHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(searchHandler))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Error("cant get %v")
	}

	want := 200
	if got := res.StatusCode; got != want {
		t.Fatalf("status_code = $v; want = %v", got, want)
	}
	fmt.Println(res)
}
