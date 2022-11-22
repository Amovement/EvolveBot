package app_core

import (
	"fmt"
	config "github.com/emolve/EvolveBot/pkg/config"
	"os"
)

func getMode(_ Core, options ...interface{}) {
	cfg := config.GetConfig()
	fmt.Printf("App Mode: %s\n", cfg.App.Mode)
}

func getAppHelpInfo(_ Core, options ...interface{}) {
	fmt.Printf("%s\n", GetAppHelpInfo())
}

func loginFishPi(app Core, options ...interface{}) {
	cfg := config.GetConfig()
	app.Sel.Login(cfg)
}

func exitApp(app Core, options ...interface{}) {
	app.Exit <- os.Interrupt
}

func clickMetalA(app Core, options ...interface{}) {
	go app.Sel.ClickMetalA()
}
