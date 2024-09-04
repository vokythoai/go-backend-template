package main

import "qropen-backend/internal/app"

func main() {
	app := app.NewApp()
	app.Run(":8080")
}
