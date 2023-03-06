package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	model "github.com/sramirezpch/ipfs-writer/src/writer/model"
	p "github.com/wabarc/ipfs-pinner/pkg/pinata"
)

type IPFSWriter struct {
	pinata *p.Pinata
}

func NewIPFSWriter() *IPFSWriter {
	return &IPFSWriter{pinata: &p.Pinata{Apikey: os.Getenv("PINATA_API_KEY"), Secret: os.Getenv("PINATA_SECRET_KEY")}}
}

func (w *IPFSWriter) PinJSON(data model.Metadata) ([]byte, error) {
	url := "https://api.pinata.cloud/pinning/pinJSONToIPFS"

	json, jsonErr := json.Marshal(data)
	if jsonErr != nil {
		return nil, jsonErr
	}

	req, reqErr := http.NewRequest("POST", url, strings.NewReader(string(json)))
	if reqErr != nil {
		return nil, reqErr
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("PINATA_JWT")))
	req.Header.Add("Content-Type", "application/json")

	client := http.Client{}

	resp, respErr := client.Do(req)
	if respErr != nil {
		return nil, respErr
	}

	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return nil, readErr
	}

	return body, nil
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
	resp, respErr := client.Do(req)
	if respErr != nil {
		fmt.Println(err.Error())
		return "", respErr
	}
	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return "", readErr
	}

	return string(body), nil
}

func (w *IPFSWriter) ListPinnedFiles() ([]byte, error) {
	url := "https://api.pinata.cloud/data/pinList?status=pinned"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("PINATA_JWT")))

	client := http.Client{}

	resp, reqErr := client.Do(req)
	if reqErr != nil {
		fmt.Println(reqErr.Error())
		return nil, reqErr
	}
	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		fmt.Println(readErr.Error())
		return nil, readErr
	}

	return body, nil
}
