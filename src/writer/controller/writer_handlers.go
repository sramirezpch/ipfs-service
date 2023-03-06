package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	model "github.com/sramirezpch/ipfs-writer/src/writer/model"
	"github.com/sramirezpch/ipfs-writer/src/writer/service"
)

type IPFSWriterHandler struct {
	Writer service.Handler
}

func (h *IPFSWriterHandler) HandlePinFile(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}

	log.Printf("Content-Length: %s", r.Header.Get("Content-Length"))

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	title := r.FormValue("title")
	description := r.FormValue("description")

	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var pinnedFile model.PinnedFile
	formData := model.FormData{
		Title:       title,
		Description: description,
		File:        fileData,
	}

	hash, err := h.Writer.PinJSON(formData)

	unmarshalErr := json.Unmarshal(hash, &pinnedFile)
	if unmarshalErr != nil {
		http.Error(w, fmt.Sprintf("Something happened\n%s", unmarshalErr.Error()), http.StatusInternalServerError)
		return
	}

	log.Printf("Pinata file hash: %s", pinnedFile.IpfsHash)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Metadata pinned successfully!", "ipfs_hash": pinnedFile.IpfsHash})
}

func (h *IPFSWriterHandler) HandleUnpinFile(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.URL.Path)
	params := mux.Vars(r)

	cid := params["cid"]

	_, err := h.Writer.UnpinJSON(cid)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *IPFSWriterHandler) HandleListPinnedFiles(w http.ResponseWriter, r *http.Request) {
	res, err := h.Writer.ListPinnedFiles()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(res), &jsonMap)
	json.NewEncoder(w).Encode(jsonMap)
}
