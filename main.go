package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	config "github.com/sramirezpch/ipfs-writer/config"
	controller "github.com/sramirezpch/ipfs-writer/src/controller"
	pinata "github.com/sramirezpch/ipfs-writer/src/service/pinata"
	"go.uber.org/zap"
)

func NewRouter(logger *zap.Logger) *mux.Router {
	config := config.NewConfig()

	pinataIPFS := pinata.NewIPFSWriter(config)

	ipfsWriterHandler := &controller.IPFSWriterHandler{Writer: pinataIPFS, Logger: logger}

	r := mux.NewRouter()
	applyJSONToResponseHeader := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	}
	r.Use(applyJSONToResponseHeader)
	r.HandleFunc("/pin", ipfsWriterHandler.HandlePinFile).Methods("POST")
	r.HandleFunc("/unpin/{cid}", ipfsWriterHandler.HandleUnpinFile).Methods("DELETE")
	r.HandleFunc("/pin", ipfsWriterHandler.HandleListPinnedFiles).Methods("GET")
	return r
}

func main() {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("Couldn't initialize logger: %s\n", err)
	}

	router := NewRouter(logger)
	c := cors.AllowAll()

	err = http.ListenAndServe(":8080", c.Handler(router))
	if err != nil {
		log.Fatalf("Couldn't start the server: %s\n", err.Error())
		os.Exit(0)
	}

	log.Println("Server started in port 8080")
}
