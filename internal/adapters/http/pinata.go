package http

import (
	"github.com/gin-gonic/gin"
	s "github.com/sramirezpch/ipfs-service/internal/service"
)

type Routes struct {
	Method  string
	Path    string
	Handler func(c *gin.Context)
}

func CreateHandlers(r *gin.Engine, service *s.IPFSService) {
	routes := []Routes{
		{
			Method: "POST", Path: "/pin", Handler: service.PinJSON,
		},
	}

	for i := 0; i < len(routes); i++ {
		r.Handle(routes[i].Method, routes[i].Path, routes[i].Handler)
	}
}
