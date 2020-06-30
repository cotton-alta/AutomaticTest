package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/PuerkitoBio/goquery"
	"github.com/sclevine/agouti"
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
	keywords.Fill("電気")
	err = page.FindByID("skwr_search").Click()
	if err != nil {
		log.Fatal(err)
	}
	// list := page.FindByClass("list")

	content, err := page.HTML()
	// content, err := list.HTML()
	if err != nil {
			log.Printf("Failed to get html: %v", err)
	}

	reader := strings.NewReader(content)
	doc, _ := goquery.NewDocumentFromReader(reader)
	href_list := []string{}
	body_list := []string{}
	doc.Find("table .column_odd a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		href = os.Getenv("URL_HEAD") + href
		href_list = append(href_list, href)

		body := s.Text()
		body_list = append(body_list, body)
	})
	doc.Find("table .column_even a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		href = os.Getenv("URL_HEAD") + href
		href_list = append(href_list, href)
		
		body := s.Text()
		body_list = append(body_list, body)
	})
	
	for i, href := range href_list {
		page.Navigate(href)
		name := "./" + body_list[i] + ".jpg"
		page.Screenshot(name)
		content, err = page.HTML()
		reader = strings.NewReader(content)
		doc, _ = goquery.NewDocumentFromReader(reader)
		fmt.Printf("content: %s\n", content)	
	}

	fmt.Printf("href_list: %s\n", href_list)
	fmt.Printf("body_list: %s\n", body_list)
	page.Screenshot("./screen.jpg")
}
