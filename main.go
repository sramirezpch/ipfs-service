package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	w "github.com/sramirezpch/ipfs-writer/src/writer"
)

func contentTypeApplicationJsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}
func NewRouter() *mux.Router {
	ipfsWriter := w.NewIPFSWriter()

	ipfsWriterHandler := &w.IPFSWriterHandler{Writer: ipfsWriter}

	r := mux.NewRouter()
	r.HandleFunc("/pin-file", ipfsWriterHandler.HandlePinFile).Methods("POST")
	r.Use(contentTypeApplicationJsonMiddleware)

	return r
}

func main() {
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
