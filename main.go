package main

import (
	"mmulyana/todolist-be/config"
	"mmulyana/todolist-be/database"
	"mmulyana/todolist-be/routes"
)

func main() {
	config.LoadEnv()

	database.InitDB()

	r := routes.SetupRouter()

	r.Run(":" + config.GetEnv("APP_PORT", "3000"))
}
