package routers

import (
	"net/http"
	"story/handlers"
	"story/pkg/database"
)

// Setup sets routers
func Setup() *http.ServeMux {
	api := &handlers.APIEnv{
		DB: database.GetDB(),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/stories/", api.StoriesHandler)
	mux.HandleFunc("/event/", api.EventHandler)
	return mux

	// --------------------------------------------------
	// firstly developed routers with gin
	/*
		router := gin.Default()
		router.GET("/stories/:app_token",api.StoriesHandlerGin)
		return router
	*/
}
