package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	model "github.com/sramirezpch/ipfs-writer/src/model"
	service "github.com/sramirezpch/ipfs-writer/src/service"
	image "github.com/sramirezpch/ipfs-writer/src/service/image"
)

type PinataHandler struct {
	IpfsService  service.Handler
	ImageService *image.ImageService
}

func (h *PinataHandler) HandlePinFile(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}

	log.Printf("Content-Length: %s", r.Header.Get("Content-Length"))

	file, header, err := r.FormFile("file")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Send image to another microservice to handle the S3 upload
	h.ImageService.SaveImage(file, header)

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

	hash, err := h.IpfsService.PinJSON(formData)

	unmarshalErr := json.Unmarshal(hash, &pinnedFile)
	if unmarshalErr != nil {
		http.Error(w, fmt.Sprintf("Something happened\n%s", unmarshalErr.Error()), http.StatusInternalServerError)
		return
	}

	log.Printf("Pinata file hash: %s", pinnedFile.IpfsHash)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Metadata pinned successfully!", "ipfs_hash": pinnedFile.IpfsHash})
}

func (h *PinataHandler) HandleUnpinFile(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s", r.Method, r.URL.Path)
	params := mux.Vars(r)

	cid := params["cid"]

	_, err := h.IpfsService.UnpinJSON(cid)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *PinataHandler) HandleListPinnedFiles(w http.ResponseWriter, r *http.Request) {
	res, err := h.IpfsService.ListPinnedFiles()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(res), &jsonMap)
	json.NewEncoder(w).Encode(jsonMap)
}
