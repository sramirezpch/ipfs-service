package main

import (
	"github.com/sramirezpch/ipfs-service/internal/adapters"
	s "github.com/sramirezpch/ipfs-service/internal/service"
	"github.com/sramirezpch/ipfs-service/pkg/config"

	// p "github.com/sramirezpch/ipfs-service/pkg/ipfs-provider/pinata"
	p "github.com/sramirezpch/ipfs-service/pkg/ipfs-provider"
)

func main() {
	config := config.NewConfig()
	// pinataProvider := p.NewPinataProvider()
	defaultProvider := p.NewDefaultProvider(config)
	ipfsService := s.NewIPFSService(defaultProvider)

	router := adapters.NewRouter(ipfsService)

	router.EnableCORS().AttachHandlers().Run(config)
}
