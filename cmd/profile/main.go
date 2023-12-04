package main

import (
	"task-5-pbi-btpns-deianearra/internal/app"
	"task-5-pbi-btpns-deianearra/internal/config/database"
)

func main() {
	database.Connect()
	app.Start()
}
