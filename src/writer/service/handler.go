package service

import "net/http"

type Handler interface {
	HandlePinFile(w http.ResponseWriter, r *http.Request)
}
