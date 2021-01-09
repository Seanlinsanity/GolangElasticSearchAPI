package app

import (
	"net/http"

	"github.com/Seanlinsanity/GolangElasticSearchAPI/controllers"
)

func mapUrls() {
	router.HandleFunc("/ping", controllers.PingController.Ping).Methods(http.MethodGet)
	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)
	router.HandleFunc("/item", controllers.ItemsController.Get).Methods(http.MethodGet)
}
