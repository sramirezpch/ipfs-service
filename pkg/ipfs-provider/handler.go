package ipfsprovider

import (
	"github.com/sramirezpch/ipfs-service/internal/entity"
	"github.com/sramirezpch/ipfs-service/pkg/config"
)

type Handler interface {
	PinJSON(data entity.Form) (bool, error)
	UnpinJSON(cid string) (bool, error)
	ListPinnedFiles() (bool, error)
}

type DefaultProvider struct {
	config *config.Config
}

func NewDefaultProvider(config *config.Config) *DefaultProvider {
	return &DefaultProvider{
		config: config,
	}
}

func (p *DefaultProvider) PinJSON(data entity.Form) (bool, error) {
	return true, nil
}

func (p *DefaultProvider) UnpinJSON(cid string) (bool, error) {
	return true, nil
}

func (p *DefaultProvider) ListPinnedFiles() (bool, error) {
	return true, nil
}
