package main

import (
	"connecting-server/api"
	"connecting-server/app"
	"connecting-server/lib"
)

func init() {
	app.NewEnv()
}

func main() {
	port := lib.GetEnvWithKey("PORT")
	server := app.New()

	api.ApplyRoutes(server)
	server.Logger.Fatal(server.Start(":" + port))
}
