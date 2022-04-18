package config

import (
	"fmt"
	"log"
	"net/http"
)

func StartServer() {
	fmt.Printf("Starting server on port 3333 🚀 \n")
	if err := http.ListenAndServe(":3333", nil); err != nil {
		log.Fatal(err)
	}
}
