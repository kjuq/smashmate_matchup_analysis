package main

import (
	"strings"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func sqlConnect() (*sql.DB, error) {
	DBMS := "mysql"
	USER := "root"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "mydb"

	CONNECT := USER + "@" + PROTOCOL + "/" + DBNAME
	db, err := sql.Open(DBMS, CONNECT)
	return db, err
}

func sqlEscape(str string) string {
	result := strings.Replace(str, "'", "''", -1)
	return result
}

func getWinnerInfo(db *sql.DB, roomId int) (Player, error) {
	var winner Player
	statemnt := fmt.Sprintf(`select winner_player, winner_fighter, winner_rate 
							 from %s where room_id=%d`,
							 tableName, roomId)

	err := db.QueryRow(statemnt).Scan(&winner.name, &winner.fighter, &winner.rate)

	return winner, err
}

func insertPlayerData(db *sql.DB, roomId int, winner Player, loser Player) error {
	statemnt := fmt.Sprintf("insert into `%s` values (%d, '%s', '%s', %d, '%s', '%s', %d);",
							 tableName,
							 roomId,
							 winner.name,
							 winner.fighter,
							 winner.rate,
							 loser.name,
							 loser.fighter,
							 loser.rate)

	_, err := db.Exec(statemnt)

	return err
}

