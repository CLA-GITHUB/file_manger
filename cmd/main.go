package main

import (
	"log"
	"os"

	"github.com/cla-github/file_manager/api"
	"github.com/cla-github/file_manager/db"
	"github.com/cla-github/file_manager/internal"
)

func init() {
	// Load env variables
	internal.LoadEnvValues()

	// connect  to db
	db.Connect()
}

func main() {
	port := os.Getenv("PORT")

	r := api.SetupRouter()

	log.Fatal(r.Run(port))
}
