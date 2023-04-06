package service

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/sramirezpch/ipfs-service/internal/entity"
	ipfsprovider "github.com/sramirezpch/ipfs-service/pkg/ipfs-provider"
)

type IPFSService struct {
	Writer ipfsprovider.Handler
}

func NewIPFSService(writer ipfsprovider.Handler) *IPFSService {
	return &IPFSService{Writer: writer}
}

func (ipfsService *IPFSService) PinJSON(c *gin.Context) {
	form, _ := c.MultipartForm()

	image := form.File["Image"][0]
	fmt.Println(image)
	title := form.Value["Title"][0]
	description := form.Value["Description"][0]

	fmt.Println()
	log.Println(title, description)
	ipfsService.Writer.PinJSON(entity.Form{
		Image:       image,
		Title:       title,
		Description: description,
	})
}
