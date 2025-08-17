package web

import "net/http"

var (
	cfg = config{}
)

func CreateMuxServe(dbURL string) *http.ServeMux {
	mux := http.NewServeMux()
	return mux
}
