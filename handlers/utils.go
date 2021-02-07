package handlers

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const (
	MustBePost    = "request method must be post"
	MustBeGet     = "request method must be get"
	WentWrong     = "something went wrong"
	TokenNotFound = "app_token not found"
)

// ResponseObj stores message for handler response
type ResponseObj struct {
	Msg string
}

func handlerResponse(w http.ResponseWriter, status int, resp interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	byteArray, err := json.Marshal(resp)
	if err != nil {
		log.Error(err)
	}
	_, err = w.Write(byteArray)
	if err != nil {
		log.Error(err)
	}
}

// APIEnv stores handlers dependent objects
type APIEnv struct {
	DB *gorm.DB
}
