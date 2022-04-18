package routes

import (
	"net/http"

	"github.com/BernardoDeveloper/learngoapifromzero/src/controllers"
)

func Routes() {
	// fruits
	http.HandleFunc("/", controllers.ShowFruits)
	http.HandleFunc("/fruits", controllers.CreateFruit)
}
