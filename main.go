package main

import (
	_ "github.com/lib/pq"
	"github.com/thakkarkaushal/http-service-clean-architecture/config"
	"github.com/thakkarkaushal/http-service-clean-architecture/services"
	"github.com/thakkarkaushal/http-service-clean-architecture/store"
	"github.com/thakkarkaushal/http-service-clean-architecture/web"
)

func main() {
	config.Connection()
	var userStore = store.New()
	var service = services.New(userStore)

	server := web.New(service)
	server.Run()
}
