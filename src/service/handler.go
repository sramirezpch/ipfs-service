package service

import (
	"github.com/sramirezpch/ipfs-writer/src/model"
)

type Handler interface {
	PinJSON(data model.FormData) ([]byte, error)
	UnpinJSON(cid string) (string, error)
	ListPinnedFiles() ([]byte, error)
}
