package main

import (
	"fmt"
	"github.com/cla-github/file_manager/api"
	"log"
)

func main() {

	r := api.SetupRouter()

	log.Fatal(r.Run(":8080"))
	fmt.Println("Hello world!")
}
