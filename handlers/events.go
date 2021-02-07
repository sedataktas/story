package handlers

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"net/http"
	"story/models"
	"story/pkg/database"
	"strings"
)

// EventHandler is a POST handler
// takes app_token as parameter
// takes event infos from request body
// if there is a record according to these, increment count field 1 in events table
// else add event to events table
func (a *APIEnv) EventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Error("request method must be post")
		handlerResponse(w, http.StatusBadRequest, ResponseObj{Msg: MustBePost})
		return
	}

	appToken := strings.Replace(r.URL.Path, "/event/", "", 1)
	appID, err := database.GetAppIDByAppToken(a.DB, appToken)
	if err != nil {
		log.Error(err)
		handlerResponse(w, http.StatusInternalServerError, ResponseObj{Msg: WentWrong})
		return
	}
	if appID == 0 {
		handlerResponse(w, http.StatusNotFound, ResponseObj{Msg: TokenNotFound})
		return
	}

	eventFromBody, err := getEventFromBody(r.Body)
	if err != nil {
		log.Error(err)
		handlerResponse(w, http.StatusBadRequest, ResponseObj{Msg: "error occured when reading body"})
		return
	}

	storyIDExists, err := database.CheckStoryIDExists(a.DB, appID, eventFromBody.StoryID)
	if err != nil {
		log.Error(err)
		handlerResponse(w, http.StatusInternalServerError, ResponseObj{Msg: WentWrong})
		return
	}
	if !storyIDExists {
		handlerResponse(w, http.StatusBadRequest, ResponseObj{Msg: "story_id not exists in db"})
		return
	}

	eventFromBody.AppID = appID
	if err := addEventToDB(a.DB, eventFromBody); err != nil {
		log.Error(err)
		handlerResponse(w, http.StatusInternalServerError, ResponseObj{Msg: TokenNotFound})
		return
	}

	handlerResponse(w, http.StatusOK, ResponseObj{Msg: "event successfully added to db"})
}

func getEventFromBody(body io.ReadCloser) (models.Event, error) {
	var event models.Event
	b, err := ioutil.ReadAll(body)
	if err != nil {
		return event, err
	}

	err = json.Unmarshal(b, &event)
	if err != nil {
		return event, err
	}
	return event, nil
}

func addEventToDB(db *gorm.DB, eventFromBody models.Event) error {
	eventFromDB, err := database.GetEvent(db, eventFromBody)
	if err != nil {
		return err
	}

	var nilEvent models.Event
	// there is no event matching with request body fields in db
	// store incoming event from request to db
	if eventFromDB == nilEvent {
		eventFromBody.Count = 1
		db.Save(&eventFromBody)
	} else {
		// increments 1 event count in db
		eventFromDB.Count++
		db.Save(&eventFromDB)
	}
	return nil
}
