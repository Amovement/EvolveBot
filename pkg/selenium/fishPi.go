package selenium

import (
	"fmt"
	"github.com/emolve/EvolveBot/pkg/config"
	"github.com/tebeka/selenium"
	"time"
)

type fishPiEvolve struct {
	webDriver selenium.WebDriver
}

func (f fishPiEvolve) Login(config config.Config) {
	wd := f.webDriver
	wd.Get(config.User.LoginPath)
	we, _ := wd.FindElement(selenium.ByID, "nameOrEmail")
	we.SendKeys(config.User.Username)
	we, _ = wd.FindElement(selenium.ByID, "loginPassword")
	we.SendKeys(config.User.Password)
	// login button
	we, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[3]/div/div[1]/div/button[1]")
	we.Click()
	if err != nil {
		fmt.Println("[Error] Login FindElement error: ", err)
	}
}

// LoadCloudArchive load user cloud documents
func (f fishPiEvolve) LoadCloudArchive(config config.Config) {
	wd := f.webDriver
	// goto setting page
	set, _ := wd.FindElement(selenium.ByXPATH, "/html/body/div[2]/div/div[2]/div/div/nav/ul/li[8]/a")
	set.Click()

	time.Sleep(time.Second * 5) // wait css load

	// goto PlayFab (启用云存档)
	os, err := wd.FindElement(selenium.ByXPATH, "/html/body/div[2]/div/div[2]/div/div/section/div[8]/div[8]")
	if err != nil {
		fmt.Println(err)
	}
	check, err := os.FindElement(selenium.ByXPATH, "/html/body/div[2]/div/div[2]/div/div/section/div[8]/div[8]/div/div[1]/label/span[1]")
	if err != nil {
		fmt.Println(err)
	}
	check.Click()

	// user login
	user, _ := wd.FindElement(selenium.ByID, "playfab-username")
	user.SendKeys(config.User.Username)
	pwd, _ := wd.FindElement(selenium.ByID, "playfab-password")
	pwd.SendKeys(config.User.Password)
	login, _ := wd.FindElement(selenium.ByXPATH, "/html/body/div[2]/div/div[2]/div/div/section/div[8]/div[8]/div/div[2]/div/div[1]/div/section/div[1]/div/button")
	login.Click()

	time.Sleep(time.Second * 5) // wait css load
	st, _ := wd.FindElement(selenium.ByXPATH, "/html/body/div[2]/div/div[2]/div/div/section/div[8]/div[8]/div/div[2]/div/div[2]/p/button[1]")
	st.Click()

}

func (f fishPiEvolve) ClickMetalA() {
	wd := f.webDriver
	for true {
		we, _ := wd.FindElement(selenium.ByID, "incSheet_MetalA")
		we.Click()
		we, _ = wd.FindElement(selenium.ByID, "incWrought_IronA")
		we.Click()
		we, _ = wd.FindElement(selenium.ByID, "incBrickA")
		we.Click()

		time.Sleep(time.Second * 10)
	}

}
