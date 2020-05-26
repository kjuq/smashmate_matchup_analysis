package main

import (
	"fmt"
	"log"
	"strings"
	"strconv"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const base_url string = "https://smashmate.net/rate/"
const room_id int = 114157 //72751

var url string = base_url + strconv.Itoa(room_id)

func ScrapeMatch() {
	fmt.Println(url)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	title := doc.Find(".col-xs-6")
	title.Each(func(i int, s *goquery.Selection) {
		character_sel := s.Find(".smash-icon").First()
		character_img, _ := character_sel.Attr("src")
		character_splitted := strings.Split(character_img, "/")
		character_png := character_splitted[len(character_splitted) - 1]
		character := character_png[:len(character_png) - 4]

		var player_name string
		var rate int
		//var character = "hoge"
		var match_status string

		s.Find(".col-xs-7").Each(func(j int, t *goquery.Selection) {
			content := strings.TrimSpace(t.Text())
			switch j {
			case 0: //name
				player_name = content
			case 1: // rate
				rate_str := content
				rate, _ = strconv.Atoi(rate_str)
			case 2: // fighter
			case 3: // win or lose
				match_status = content

			}
		})

		/*
		var isCanceled, isWin bool

		switch match_status {
		case "対戦中止":
			isCanceled = true
		case "勝ち":
			isWin = true
		case "負け":
			isWin = false
		}
		*/

		fmt.Println(player_name, rate, character, match_status)
	})
}

func main() {
	ScrapeMatch()
}

