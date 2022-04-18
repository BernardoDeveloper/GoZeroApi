package main

import (
	"github.com/BernardoDeveloper/learngoapifromzero/config"
	"github.com/BernardoDeveloper/learngoapifromzero/config/database"
	"github.com/BernardoDeveloper/learngoapifromzero/src/routes"
)

func main() {
	database.StartConnect()

	routes.Routes()
	config.StartServer()
}
