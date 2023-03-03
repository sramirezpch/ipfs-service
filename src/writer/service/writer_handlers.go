package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	controller "github.com/sramirezpch/ipfs-writer/src/writer/controller"
	dao "github.com/sramirezpch/ipfs-writer/src/writer/dao"
)

type IPFSWriterHandler struct {
	Writer *controller.IPFSWriter
}

func (h *IPFSWriterHandler) HandlePinFile(w http.ResponseWriter, r *http.Request) {
	var m dao.Metadata

	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to decode the body\n%s", err.Error()), http.StatusBadRequest)
		return
	}

	hash, pinErr := h.Writer.PinJSON(m)

	if pinErr != nil {
		http.Error(w, fmt.Sprintf("Failed to pin the file to Pinata\n%s", pinErr.Error()), http.StatusBadRequest)
		return
	}

	var pinnedFile dao.PinnedFile
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
