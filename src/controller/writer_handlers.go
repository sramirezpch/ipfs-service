package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sramirezpch/ipfs-writer/src/model"
	"github.com/sramirezpch/ipfs-writer/src/service"
	"go.uber.org/zap"
)

type IPFSWriterHandler struct {
	Writer service.Handler
	Logger *zap.Logger
}

func (h *IPFSWriterHandler) HandlePinFile(w http.ResponseWriter, r *http.Request) {
	defer h.Logger.Sync()

	h.Logger.Info("received a request", zap.String("request uri", r.RequestURI), zap.String("host", r.URL.Host))
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.Logger.Error(err.Error())
		return

	}

	file, _, err := r.FormFile("File")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.Logger.Error(err.Error())
		return
	}
	defer file.Close()

	title := r.FormValue("Title")
	description := r.FormValue("Description")

	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.Logger.Error(err.Error())
		return
	}

	h.Logger.Debug("Reading body of the request", zap.String("title", title), zap.String("description", description))

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
		h.Logger.Error(unmarshalErr.Error())
		return
	}

	h.Logger.Debug("Successfully pinned file", zap.String("pinata file hash", pinnedFile.IpfsHash), zap.String("timestamp", pinnedFile.Timestamp.String()))

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Metadata pinned successfully!", "ipfs_hash": pinnedFile.IpfsHash})
}

func (h *IPFSWriterHandler) HandleUnpinFile(w http.ResponseWriter, r *http.Request) {
	defer h.Logger.Sync()
	h.Logger.Info("received a request", zap.String("request uri", r.RequestURI), zap.String("host", r.URL.Host))

	params := mux.Vars(r)

	cid := params["cid"]

	_, err := h.Writer.UnpinJSON(cid)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.Logger.Error(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *IPFSWriterHandler) HandleListPinnedFiles(w http.ResponseWriter, r *http.Request) {
	res, err := h.Writer.ListPinnedFiles()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.Logger.Error(err.Error())
		return
	}

	var jsonMap map[string]interface{}
	json.Unmarshal([]byte(res), &jsonMap)
	json.NewEncoder(w).Encode(jsonMap)
}
