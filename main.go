package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
//	"errors"

	"github.com/PuerkitoBio/goquery"
)

const baseURL string = "https://smashmate.net/rate/"
const roomId int = 114157 //72751

var url string = baseURL + strconv.Itoa(roomId)

func ScrapeMatch() (error) {
	fmt.Println(url)

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}

	title := doc.Find(".col-xs-6")
	title.Each(func(i int, s *goquery.Selection) {
		characterSel := s.Find(".smash-icon").First()
		characterImg, exists := characterSel.Attr("src")
		if !exists {
			return //errors.New("no image src found error")
		}
		characterSplitted := strings.Split(characterImg, "/")
		characterPNG := characterSplitted[len(characterSplitted)-1]
		character := characterPNG[:len(characterPNG)-4]

		var playerName string
		var rate int
		var matchStatus string

		s.Find(".col-xs-7").Each(func(j int, t *goquery.Selection) {
			content := strings.TrimSpace(t.Text())
			switch j {
			case 0: // name
				playerName = content
			case 1: // rate
				rateStr := content
				rate, _ = strconv.Atoi(rateStr)
			case 2: // fighter
			case 3: // win or lose
				matchStatus = content

			}
		})

		fmt.Println(playerName, rate, character, matchStatus)
	})

	return nil
}

func main() {
	if err := ScrapeMatch(); err != nil {
		fmt.Print("error occered")
		log.Fatal(err)
	}
}

