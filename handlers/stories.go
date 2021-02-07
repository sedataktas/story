package handlers

import (
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"net/http"
	"story/models"
	"story/pkg/database"
	"strings"
)

// StoriesHandler is a GET handler
// takes app_token as parameter
// returns metadatas
func (a *APIEnv) StoriesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Error("request method must be get")
		handlerResponse(w, http.StatusBadRequest, ResponseObj{Msg: MustBeGet})
		return
	}

	appToken := strings.Replace(r.URL.Path, "/stories/", "", 1)

	// In the first implementation this block used, but now not used for better performance
	/*
		exists, err := database.CheckAppTokenExists(a.DB, appToken)
		if err != nil {
			log.Print(err)
			handlerResponse(w, http.StatusInternalServerError,ResponseObj{Msg: "something went wrong"})
			return
		}
		if !exists {
			handlerResponse(w, http.StatusNotFound,ResponseObj{Msg: "app_token not found"})
			return
		}
	*/
	stories, err := database.GetMetadatasByAppToken(a.DB, appToken)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Error(err)
		handlerResponse(w, http.StatusInternalServerError, ResponseObj{Msg: WentWrong})
		return
	}

	if len(stories) == 0 {
		handlerResponse(w, http.StatusNotFound, ResponseObj{Msg: "stories metadatas not found"})
		return
	}

	dto := models.StoriesToDTO(stories)
	handlerResponse(w, http.StatusOK, dto)
}

// ---------------------------------------------------------------------------------------------------------
// Firstly developed handler with gin
/*
func (a *APIEnv)StoriesHandlerGin(c *gin.Context){
	appToken := c.Param("app_token")

	exists, err := database.CheckAppTokenExists(a.DB, appToken)
	if err != nil {
		log.Print(err)
		c.JSON(http.StatusInternalServerError,ResponseObj{Msg: "something went wrong"})
		return
	}
	if !exists {
		c.JSON( http.StatusNotFound,ResponseObj{Msg: "app_token not found"})
		return
	}

	stories, err := database.GetMetadatasByAppToken(a.DB,appToken)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Print(err)
		c.JSON(http.StatusInternalServerError,ResponseObj{Msg: "something went wrong"})
		return
	}

	if stories == nil {
		c.JSON(http.StatusNotFound, ResponseObj{Msg: "stories metadatas not found"})
		return
	}

	dto := models.StoriesToDTO(stories)
	c.JSON(http.StatusOK, dto)
}
*/
