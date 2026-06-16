package main

import (
	"log"

	"github.com/rajeshbond/smart/application"
	"github.com/rajeshbond/smart/internal/common/utils"
)

func main() {
	utils.InitValidator()
	app := application.NewApp()
	defer app.DB.Close()
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
