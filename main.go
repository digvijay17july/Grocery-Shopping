package main

import (
	"Grocery-Shopping-Category-Module/src/app/api"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)
func main() {
	fmt.Println("Starting Category Module.... ")
	config := api.GetConfig()

	app := &api.App{}
	app.Initialize(config)

	port := os.Getenv("PORT")
	fmt.Println("Port No. is :"+port)
	app.Run(":"+port)

	fmt.Println("Started Category Module.... ")
}