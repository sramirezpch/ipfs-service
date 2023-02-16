package writer

import (
	"encoding/json"
	"fmt"
	"os"

	p "github.com/wabarc/ipfs-pinner/pkg/pinata"
)

type IPFSWriter struct {
	writer *p.Pinata
}

func NewIPFSWriter() *IPFSWriter {
	return &IPFSWriter{writer: &p.Pinata{Apikey: os.Getenv("PINATA_API_KEY"), Secret: os.Getenv("PINATA_SECRET_KEY")}}
}

func (w *IPFSWriter) PinJSON(data Metadata) {

	json, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
	}

	w.writer.PinWithBytes(json)

	if err != nil {
		fmt.Println(err.Error())
	}
}
