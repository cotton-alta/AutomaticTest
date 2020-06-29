package main

import (
	"fmt"
	"log"
	// "time"
	"github.com/sclevine/agouti"
	"os"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	options := agouti.ChromeOptions(
		"args", []string{
			"--headless",
			"--window-size=1440,900",
			"no-sandbox",
			"--disable-dev-shm-usage",
			"--disable-gpu",
		})

	driver := agouti.ChromeDriver(options)
	// driver := agouti.ChromeDriver(agouti.Browser("chrome"))
	defer driver.Stop()

	if err := driver.Start(); err != nil {
		fmt.Println("err1")
		log.Fatal(err)
	}

	page, err := driver.NewPage()
	if err != nil {
		fmt.Println("err2")
		log.Fatal(err)
	}

	page.Navigate(os.Getenv("URL"))
	keywords := page.FindByID("keywords")
	keywords.Fill("a")
	err = page.FindByID("skwr_search").Click()
	if err != nil {
		log.Fatal(err)
	}
	page.Screenshot("./screen.jpg")
}
