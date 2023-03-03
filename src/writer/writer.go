package writer

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
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

	hash, pinErr := w.pinata.PinWithBytes(json)
	if pinErr != nil {
		fmt.Println(pinErr.Error())
		return "", errors.New("failed to pin the data")
	}

	return hash, nil
}

func (w *IPFSWriter) UnpinJSON(cid string) (string, error) {
	url := "https://api.pinata.cloud/pinning/unpin/"

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s%s", url, cid), nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("PINATA_JWT")))

	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	client := http.Client{}
	resp, reqError := client.Do(req)
	if reqError != nil {
		fmt.Println(err.Error())
		return "", reqError
	}
	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return "", readErr
	}

	return string(body), nil
}
