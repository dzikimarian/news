package main

import (
	"fmt"
	"time"
	"net/http"
	"log"
	"strings"
	"strconv"
	"github.com/PuerkitoBio/goquery"
)

func main() {
    
	var newsId int
	fmt.Print("\033[H\033[2J") //clear console

	for {
		document := get("https://wiadomosci.onet.pl/swiat/wojna-rosja-ukraina-nie-jestem-ikona-ukraina-jest-relacja-na-zywo/ztv0qk4")
		
		oldNewsId := newsId

		newsItem := document.Find(".commentatorLiveClient .messageItem").First()
		newsId = getNewsItemId(newsItem)

		if newsId > oldNewsId {
	
			fmt.Println(strings.TrimSpace(newsItem.Text()))
		} 

		time.Sleep(10 * time.Second)
	}



}

func getNewsItemId(newsItem *goquery.Selection) int{
	stringId, _ := newsItem.Attr("rel")
	id, _ := strconv.Atoi(stringId)
	return id
}

func get(url string) *goquery.Document{
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return doc
}