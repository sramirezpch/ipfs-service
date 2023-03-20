package pinata

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	config "github.com/sramirezpch/ipfs-writer/config"
	model "github.com/sramirezpch/ipfs-writer/src/writer/model"
)

type IPFSWriter struct {
	Config *config.Config
}

func NewIPFSWriter(c *config.Config) *IPFSWriter {
	return &IPFSWriter{Config: c}
}

func (w *IPFSWriter) PinJSON(data model.FormData) ([]byte, error) {
	url := "https://api.pinata.cloud/pinning/pinJSONToIPFS"

	json, jsonErr := json.Marshal(data)
	if jsonErr != nil {
		return nil, jsonErr
	}

	req, reqErr := http.NewRequest("POST", url, strings.NewReader(string(json)))
	if reqErr != nil {
		return nil, reqErr
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", w.Config.PinataJWT))
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
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", w.Config.PinataJWT))

	if err != nil {
		return "", err
	}

	client := http.Client{}
	resp, respErr := client.Do(req)
	if respErr != nil {
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
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", w.Config.PinataJWT))

	client := http.Client{}

	resp, reqErr := client.Do(req)
	if reqErr != nil {
		return nil, reqErr
	}
	defer resp.Body.Close()

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		return nil, readErr
	}

	return body, nil
}
