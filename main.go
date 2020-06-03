package main

// TODO: implement tests

import (
	"fmt"
)

const baseURL string = "https://smashmate.net/rate/"
const tableName string = "all_data"
const minRoom = 82462
const maxRoom = 173297

func main() {
	db, errSql := sqlConnect()
	if errSql != nil { panic(errSql) }
	defer db.Close()

	winner, err := getWinnerInfo(db, 5)
	if err != nil { panic(err) }

	fmt.Print(winner)
}

