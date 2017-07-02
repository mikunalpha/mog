package main

import (
	"log"

	"github.com/mikunalpha/mog/api/router"
	"github.com/mikunalpha/mog/api/services/store"
	"github.com/mikunalpha/mog/api/shared/utils"
)

func main() {
	var err error

	// Load env variables from file.
	utils.LoadEnvs("./.env")

	// Initialize database store.
	err = store.Init()
	if err != nil {
		log.Println(err)
		return
	}

	// Start API service.
	router.Start()
}
