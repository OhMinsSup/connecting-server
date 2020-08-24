package main

import (
	"connecting-server/api"
	"connecting-server/app"
)

func init() {
	app.NewEnv()
}

func main() {
	port := app.GetEnvWithKey("PORT")
	server := app.New()

	api.ApplyRoutes(server)
	server.Logger.Fatal(server.Start(":" + port))
}
