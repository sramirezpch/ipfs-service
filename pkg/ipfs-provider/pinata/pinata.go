package pinata

import (
	"github.com/sramirezpch/ipfs-service/internal/entity"
	"github.com/sramirezpch/ipfs-service/pkg/config"
)

type PinataProvider struct {
	config *config.Config
}

func NewPinataProvider() *PinataProvider {
	return &PinataProvider{}
}

func (p *PinataProvider) PinJSON(data entity.Form) (bool, error) {
	return true, nil
}

func (p *PinataProvider) UnpinJSON(cid string) (bool, error) {
	return true, nil
}

func (p *PinataProvider) ListPinnedFiles() (bool, error) {
	return true, nil
}
