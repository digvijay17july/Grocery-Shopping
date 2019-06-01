package main

import (
	"Grocery-Shopping-Category-Module/src/app/api"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	)
func main() {
	config := api.GetConfig()

	app := &api.App{}
	app.Initialize(config)
	app.Run(":3002")
}