package main

import (
	//	"encoding/json"
	"fmt"
	"log"
	"mysite/models"
	"net"
	"net/http"
	"net/http/fcgi"
	"net/url"
)

const (
	APIVersionPrefix = "/api/v0"
	APISearch        = "/search"
)

type result struct {
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	//rq := url.QueryEscape(r.URL.RawQuery)
	rq := r.URL.RawQuery
	//q := r.URL.Query().Get("q")
	v, err := url.ParseQuery(rq)
	if err != nil {
		log.Panic(err)
	}
	q := v.Get("q")
	defer func() {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		d, _ := models.Connect()
		r := d.Search(q)
		fmt.Fprint(w, r)
	}()
}

/*
func writeResponse(w http.ResponseWriter, code int, msg string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriterHeader(code)
	json.NewEncoder(w).Encode(map[string]interface{} {
		"status_code": code,
		"msg":        msg,
	})
}
*/
func main() {
	l, err := net.Listen("tcp", "127.0.0.1:9001")
	if err != nil {
		fmt.Println(err)
	}
	http.HandleFunc(APIVersionPrefix+APISearch, searchHandler)
	fcgi.Serve(l, nil)
}