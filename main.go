package main

import (
	"fmt"
	"log"

	"github.com/sclevine/agouti"
)

func main() {
	options := agouti.ChromeOptions(
		"args", []string{
			"--headless",
			"--window-size=1280,800",
			"no-sandbox",
			"--disable-dev-shm-usage",
			"--disable-gpu",
		})

	driver := agouti.ChromeDriver(options)
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

	page.Navigate("http://localhost:3000/")
	fmt.Println(page.Title())
	page.Screenshot("./screen.jpg")
}
