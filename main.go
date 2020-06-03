package main

// TODO: implement tests

import (
	"fmt"
)

func main() {
	db, errSql := sqlConnect()
	if errSql != nil { panic(errSql) }
	defer db.Close()

	winner, err := getWinnerInfo(db, 5)
	if err != nil { panic(err) }

	pDict := playerToDict(winner)

	fmt.Println(pDict)
}

