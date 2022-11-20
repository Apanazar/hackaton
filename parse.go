package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	ColorRed    = "\u001b[31m"
	ColorGreen  = "\u001b[32m"
	ColorYellow = "\u001b[33m"
	ColorReset  = "\u001b[0m"
)

type Article struct {
	Id               int
	Source           string
	Author           string
	Title            string
	Href             string
	Publication_time string
	Category         string
	Text             string
	Mood             string
	Fake_detector    string
}

func write_file(text string) {
	file, err := os.OpenFile("db.json", os.O_APPEND|os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0777)

	if err != nil {
		fmt.Println(ColorRed, err, ColorReset)
	}
	defer file.Close()

	file.WriteString(text + "\n")
}

func send_request(url string) *goquery.Document {
	defer func() {
		if err := recover(); err != nil {
			log.Println(ColorRed, err, ColorReset)
		}
	}()
	response, err := http.Get(url)

	if err != nil {
		log.Println(ColorRed, err, ColorReset)
	}
	defer response.Body.Close()

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Println(ColorRed, err, ColorReset)
	}

	return document
}

func get_article(url string) []Article {
	log.Println(ColorYellow, "Expect a connection with ", ColorGreen, url, ColorReset)

	list := []Article{}

	doc := send_request(url)
	doc.Find(".postcards .grid h2").Each(func(i int, elem *goquery.Selection) {
		id := i
		title := strings.TrimSpace(elem.Text())
		href, _ := elem.Find("a").Attr("href")

		new_doc := send_request(href)

		source, _ := new_doc.Find(".custom-logo-link img").Attr("alt")
		publication_time := new_doc.Find(".entry-date.published.updated").Text()
		author := new_doc.Find(".vcard.author").Text()
		category := new_doc.Find(".category-link").Text()
		text := new_doc.Find(".entry-content p").Text()
		mood := "Neutral"

		detektions := []string{"Reliable", "Unreliable"}
		fake_detektion := detektions[rand.Intn(2)]

		instance := Article{
			Id:               id,
			Source:           source,
			Author:           author,
			Title:            title,
			Href:             href,
			Publication_time: publication_time,
			Category:         category,
			Text:             text,
			Mood:             mood,
			Fake_detector:    fake_detektion,
		}

		list = append(list, instance)
	})

	data, _ := json.Marshal(list)
	write_file(string(data))

	log.Println(ColorYellow, "Done.", ColorReset)
	return list
}

func main() {
	target := flag.String("url", "[empty]", "specify the site address")

	flag.Usage = func() {
		w := flag.CommandLine.Output()
		fmt.Fprintln(w, "Usage of crawler:")
		flag.PrintDefaults()
	}
	flag.Parse()

	rand.Seed(time.Now().UnixNano())
	get_article(*target)
}
