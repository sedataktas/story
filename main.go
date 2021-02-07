package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"story/pkg/config"
	"story/pkg/database"
	"story/routers"
)

func init() {
	config.Setup()
	database.Setup()
}

func main() {
	mux := routers.Setup()
	log.Fatal(http.ListenAndServe(":8080", mux))
}
