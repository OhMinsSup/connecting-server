package main

import "connecting-server/app"

func init() {
	app.NewEnv()
}

func main() {
	port := app.GetEnvWithKey("PORT")
	server := app.New()

	server.Logger.Fatal(server.Start(":" + port))
}
