package writer

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type IPFSWriterHandler struct {
	Writer *IPFSWriter
}

func (h *IPFSWriterHandler) HandlePinFile(w http.ResponseWriter, r *http.Request) {
	var m Metadata

	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	hash, pinErr := h.Writer.PinJSON(m)
	if pinErr != nil {
		fmt.Println(pinErr)
		http.Error(w, "An error ocurred, please check the logs", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Metadata pinned successfully!", "ipfs_hash": hash})
}
