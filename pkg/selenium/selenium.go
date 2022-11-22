package selenium

import (
	"github.com/emolve/EvolveBot/pkg/config"
	"github.com/tebeka/selenium"
	"time"
)

type SeleniumControl interface {
	Login(config.Config)
	LoadCloudArchive(config.Config)
	ClickMetalA()
}

func NewSeleniumControl(config config.Config) SeleniumControl {
	wd := initWebDriver(config)
	return fishPiEvolve{webDriver: wd}
}

func initWebDriver(config config.Config) selenium.WebDriver {
	// TODO
	// cancel service after exit
	ops := []selenium.ServiceOption{}
	selenium.NewChromeDriverService(config.Selenium.DriverPath, config.Selenium.Port, ops...)

	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	wd, _ := selenium.NewRemote(caps, "http://127.0.0.1:9515/wd/hub")
	wd.SetImplicitWaitTimeout(time.Second * 100)
	return wd
}
