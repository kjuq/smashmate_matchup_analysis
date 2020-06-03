package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"github.com/PuerkitoBio/goquery"
)

func ScrapeMatch(roomId int) (Player, Player, error) {
	url := baseURL + strconv.Itoa(roomId)
	var winner, loser Player
	nonPlayer := Player{rate: -1}

	res, err := http.Get(url)
	if err != nil {
		return nonPlayer, nonPlayer, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nonPlayer, nonPlayer, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nonPlayer, nonPlayer, err
	}

	title := doc.Find(".col-xs-6")
	title.Each(func(i int, s *goquery.Selection) {
		characterSel := s.Find(".smash-icon").First()
		characterImg, exists := characterSel.Attr("src")
		var character string
		if exists {
			characterSplitted := strings.Split(characterImg, "/")
			characterPNG := characterSplitted[len(characterSplitted)-1]
			character = characterPNG[:len(characterPNG)-4]
		} else {
			character = "not found"
		}

		var playerName string
		var rate int
		var matchStatus string

		s.Find(".col-xs-7").Each(func(j int, t *goquery.Selection) {
			content := strings.TrimSpace(t.Text())
			switch j {
			case 0: // name
				playerName = sqlEscape(content)
			case 1: // rate
				rateStr := content
				rate, _ = strconv.Atoi(rateStr)
			case 2: // fighter
			case 3: // win or lose
				matchStatus = content

			}
		})

		switch matchStatus {
		case "勝ち":
			isCanceled := false
			winner = Player{playerName, character, rate, isCanceled}
		case "負け":
			isCanceled := false
			loser = Player{playerName, character, rate, isCanceled}
		case "対戦中止":
			winner = nonPlayer
			loser = nonPlayer
		}

	})

	return winner, loser, nil
}

func scrapeMultiplePages() error {
	db, errSql := sqlConnect()
	if errSql != nil {
		return errSql
	}
	defer db.Close()

	cnt := 0
	for roomId := minRoom; roomId <= maxRoom; roomId++ {
		winner, loser, errScrp := ScrapeMatch(roomId)
		if errScrp != nil {
			return errScrp
		}

		errInst := insertPlayerData(db, roomId, winner, loser)
		if errInst != nil {
			return errInst
		}

		if roomId%100 == 0 {
			cnt++
			fmt.Printf("%d done\n", cnt*100)
		}
	}

	fmt.Println("success")
	return nil
}
