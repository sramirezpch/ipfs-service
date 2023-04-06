package adapters

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sramirezpch/ipfs-service/internal/adapters/http"
	s "github.com/sramirezpch/ipfs-service/internal/service"
	"github.com/sramirezpch/ipfs-service/pkg/config"
)

type Router struct {
	Router      *gin.Engine
	IPFSService *s.IPFSService
}

func (r *Router) EnableCORS() *Router {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	r.Router.Use(cors.New(corsConfig))

	return r
}

func (r *Router) AttachHandlers() *Router {
	http.CreateHandlers(r.Router, r.IPFSService)

	return r
}

func (r *Router) Run(c *config.Config) {
	if err := r.Router.Run(c.Port); err != nil {
		panic(err)
	}
}

func NewRouter(ipfsService *s.IPFSService) *Router {
	r := gin.Default()

	return &Router{Router: r, IPFSService: ipfsService}
}
