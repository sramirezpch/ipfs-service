package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	w "github.com/sramirezpch/ipfs-writer/writer"
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

	log.Fatal(http.ListenAndServe(":8080", router))
}
