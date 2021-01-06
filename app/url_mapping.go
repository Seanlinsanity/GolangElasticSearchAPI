package app

import (
	"net/http"

	"github.com/Seanlinsanity/GolangElasticSearchAPI/controllers"
)

func mapUrls() {
	router.HandleFunc("/items", controllers.Create).Methods(http.MethodPost)
}
