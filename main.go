package main

import (
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"log"
	"time"
)

//func RequestHandler(ctx context.Context, event events.CloudWatchEvent) error {
//	var err error
//
//	return err
//}

//type Selenium struct {
//
//}
//
//func NewSelenium() () {
//	s := Selenium{
//
//	}
//	return s
//}

// func (s *Selenium) Start

func main() {
	//// Run Chrome browser
	//cmd := exec.Command("bash", "killall -9 'Google Chrome'")
	//cmd.Stdout = os.Stdout
	//err := cmd.Run()
	//if err != nil {
	//	log.Fatal(err)
	//}

	service, err := selenium.NewChromeDriverService("./chromedriver", 4444)
	if err != nil {
		log.Printf("failed at 1")
		panic(err)
	}
	defer service.Stop()

	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{
		Path: "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome",
		Args: []string{
			//"window-size=1920x1080",
			"--no-sandbox",
			"--disable-dev-shm-usage",
			"disable-gpu",
			//"--headless", // comment out this line to see the browser
		},
	})

	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		log.Print("failed at 2")
		panic(err)
	}

	driver.Get("https://www.google.com")

	time.Sleep(5 * time.Minute)

}
