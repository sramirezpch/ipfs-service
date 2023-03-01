package writer

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	p "github.com/wabarc/ipfs-pinner/pkg/pinata"
)

type IPFSWriter struct {
	pinata *p.Pinata
}

func NewIPFSWriter() *IPFSWriter {
	return &IPFSWriter{pinata: &p.Pinata{Apikey: os.Getenv("PINATA_API_KEY"), Secret: os.Getenv("PINATA_SECRET_KEY")}}
}

func (w *IPFSWriter) PinJSON(data Metadata) (string, error) {

	json, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err.Error())
		return "nil", errors.New("couldn't serialize the data")
	}

	fmt.Println(json)
	hash, pinErr := w.pinata.PinWithBytes(json)
	if pinErr != nil {
		fmt.Println(pinErr.Error())
		return "", errors.New("failed to pin the data")
	}

	return hash, nil
}
