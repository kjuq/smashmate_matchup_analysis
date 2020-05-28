package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
//	"errors"

	"github.com/PuerkitoBio/goquery"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const baseURL string = "https://smashmate.net/rate/"
const roomId int = 114157 //72751

var url string = baseURL + strconv.Itoa(roomId)

func ScrapeMatch() (Player, Player, error) {
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
				playerName = content
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

		//fmt.Println(playerName, rate, character, matchStatus)
	})

	return winner, loser, nil
}

func sqlConnect() (*sql.DB, error) {
    DBMS := "mysql"
    USER := "root"
    PROTOCOL := "tcp(localhost:3306)"
    DBNAME := "mydb"

    CONNECT := USER + "@" + PROTOCOL + "/" + DBNAME // + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	db, err := sql.Open(DBMS, CONNECT)
    return db, err
}

func insertPlayerData(db *sql.DB, roomId int, winner Player, loser Player) (error) {
	_, err := db.Exec("insert into `user` values (999996, 'ファイターさん', 'mario', 1500, 'ファイターに', 'samus', 1600);")
	
	return err
}

func insertExample(db *sql.DB) (error) {
	_, err := db.Exec("insert into `user` values (999996, 'ファイターさん', 'mario', 1500, 'ファイターに', 'samus', 1600);")
	
	return err
}

type Player struct {
	name string
	fighter string
	rate int
	isCanceled bool
}

func main() {
	if winner, loser, err := ScrapeMatch(); err != nil {
		fmt.Print("error occered")
		panic(err.Error())
	} else {
		fmt.Println(winner)
		fmt.Println(loser)
	}

	db, err := sqlConnect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	if err := insertExample(db); err != nil {
		panic(err.Error())
	}

}

