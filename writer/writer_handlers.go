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
	h.Writer.PinJSON(m)
}
