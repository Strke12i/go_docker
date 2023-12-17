package main

import (
	"github.com/Strke12i/go_docker/database"
	"github.com/Strke12i/go_docker/routes"
)

func main() {
	DB, err := database.ConnectToDatabase()
	if err != nil {
		panic(err)
	}

	err = routes.Router(DB).Run(":8080")
	if err != nil {
		return
	}

}
