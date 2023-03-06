package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	controller "github.com/sramirezpch/ipfs-writer/src/writer/controller"
	pinata "github.com/sramirezpch/ipfs-writer/src/writer/service/pinata"
)

func NewRouter() *mux.Router {
	pinataIPFS := pinata.NewIPFSWriter()
	ipfsWriterHandler := &controller.IPFSWriterHandler{Writer: pinataIPFS}

	r := mux.NewRouter()
	r.HandleFunc("/pin", ipfsWriterHandler.HandlePinFile).Methods("POST")
	r.HandleFunc("/unpin/{cid}", ipfsWriterHandler.HandleUnpinFile).Methods("DELETE")
	r.HandleFunc("/pin", ipfsWriterHandler.HandleListPinnedFiles).Methods("GET")
	return r
}

func main() {
	router := NewRouter()
	c := cors.AllowAll()

	err := http.ListenAndServe(":8080", c.Handler(router))
	if err != nil {
		log.Fatalf("Couldn't start the server: %s\n", err.Error())
		os.Exit(0)
	}

	log.Println("Server started in port 8080")
}
