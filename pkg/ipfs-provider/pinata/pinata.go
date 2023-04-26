package pinata

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/sramirezpch/ipfs-service/internal/entity"
	"github.com/sramirezpch/ipfs-service/pkg/config"
)

type PinataProvider struct {
	config *config.Config
}

func NewPinataProvider() *PinataProvider {
	return &PinataProvider{}
}

func (p *PinataProvider) PinJSON(data entity.Form) ([]byte, error) {
	json, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/pinning/pinJSONToIPFS", p.config.PinataUrl), strings.NewReader(string(json)))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", p.config.PinataJWT))
	req.Header.Add("Content-Type", "application/json")

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (p *PinataProvider) UnpinJSON(cid string) (bool, error) {
	return true, nil
}

func (p *PinataProvider) ListPinnedFiles() (bool, error) {
	return true, nil
}
