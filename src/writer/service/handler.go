package service

import (
	"github.com/sramirezpch/ipfs-writer/src/writer/model"
)

type Handler interface {
	PinJSON(data model.Metadata) ([]byte, error)
	UnpinJSON(cid string) (string, error)
	ListPinnedFiles() ([]byte, error)
	Hello()
}
