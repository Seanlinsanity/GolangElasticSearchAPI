package app

import (
	"net/http"
	"time"

	"github.com/Seanlinsanity/GolangElasticSearchAPI/clients/elasticSearch"
	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	elasticSearch.Init()
	mapUrls()

	srv := &http.Server{
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      router,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
