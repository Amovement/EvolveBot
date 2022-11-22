package app_core

import (
	"bufio"
	"fmt"
	"github.com/emolve/EvolveBot/pkg/config"
	"github.com/emolve/EvolveBot/pkg/selenium"
	"os"
	"strings"
)

type Command map[string]func(Core, ...interface{})

type Core struct {
	Cmd  Command
	Exit chan os.Signal
	Sel  selenium.SeleniumControl
}

func NewApp() (Core, error) {
	config := config.GetConfig()
	app := Core{
		Cmd: Command{
			"help":  getAppHelpInfo,
			"exit":  exitApp,
			"mode":  getMode,
			"login": loginFishPi,
			"metal": clickMetalA,
		},
		Exit: make(chan os.Signal),
		Sel:  selenium.NewSeleniumControl(config),
	}

	return app, nil
}

func (app Core) Run() {

	app.Cmd["help"](app, "")
	fmt.Print("/")

	buf := bufio.NewScanner(os.Stdin)
	for buf.Scan() {
		input := strings.Split(buf.Text(), " ")
		inputCommand := input[0]
		inputOption := input[1:]
		if _, ok := app.Cmd[inputCommand]; !ok {
			fmt.Printf("[Error] Unknown Command: %s\n", inputCommand)
			app.Cmd["help"](app, "")
		} else {
			app.Cmd[inputCommand](app, inputOption)
		}
		fmt.Print("/")
	}
}
