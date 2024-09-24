package main

import (
	"net/http"
)

func HandlePing(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

func HandlePostURL(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("return shortened url"))
}

func HandleGetURL(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("http://google.com"))
}
