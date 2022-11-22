package main

import (
	"github.com/emolve/EvolveBot/pkg/app_core"
)

func main() {
	app, _ := app_core.NewApp()
	go app.Run()
	<-app.Exit
}
