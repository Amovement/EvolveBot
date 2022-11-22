package app_core

var help_info string

func init() {
	help_info = `

A robot that helps play games.

COMMAND:
	help		View the help documentation
	exit		Exit the program
	mode		Show the app mode information
	login		Login game with your config.account
	metal		click metal button

To get more help with this demo, visit https://github.com/emolve/EvolveBot

	`
}

func GetAppHelpInfo() string {
	return help_info
}
