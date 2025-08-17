package main

import "net/http"

type APIConfig struct {
	port    string
	dbURL   string
	handler *http.ServeMux
}
