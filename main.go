package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	w "github.com/sramirezpch/ipfs-writer/src/writer"
)

func NewRouter() *mux.Router {
	ipfsWriter := w.NewIPFSWriter()

	ipfsWriterHandler := &w.IPFSWriterHandler{Writer: ipfsWriter}

	r := mux.NewRouter()
	r.HandleFunc("/pin-file", ipfsWriterHandler.HandlePinFile).Methods("POST")

	return r
}

func main() {
	router := NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedHeaders: []string{"*"},
	})

	log.Fatal(http.ListenAndServe(":8080", c.Handler(router)))
}
