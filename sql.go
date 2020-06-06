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

func getRoomInfo(db *sql.DB, roomId int) (RoomInfo, error) {
	var roomInfo RoomInfo
	statemnt := fmt.Sprintf(`select * from %s where roomId=%d`,
							 tableName, roomId)

	err := db.QueryRow(statemnt).Scan(&roomInfo.roomId,
									  &roomInfo.winnerName,
									  &roomInfo.winnerFighter,
									  &roomInfo.winnerRate,
									  &roomInfo.loserName,
									  &roomInfo.loserFighter,
									  &roomInfo.loserRate)

	return roomInfo, err
}

func insertRoomInfo(db *sql.DB, roomInfo RoomInfo) error {
	statemnt := fmt.Sprintf("insert into `%s` values (%d, '%s', '%s', %d, '%s', '%s', %d);",
							 tableName,
							 roomInfo.roomId,
							 roomInfo.winnerName,
							 roomInfo.winnerFighter,
							 roomInfo.winnerRate,
							 roomInfo.loserName,
							 roomInfo.loserFighter,
							 roomInfo.loserRate)

	_, err := db.Exec(statemnt)

	return err
}

func deleteRoomInfo(db *sql.DB, roomId int) error {
	statemnt := fmt.Sprintf("delete from %s where roomId = %d", tableName, roomId)

	_, err := db.Exec(statemnt)

	return err
}

func updateRoomInfo(db *sql.DB, roomInfo RoomInfo) error {
	roomDict := structToDict(roomInfo)
	for key, val := range roomDict {
		isUpdatable := true
		if val == "" { isUpdatable = false }
		if key == "winnerRate" && val == "-1" { isUpdatable = false }
		if key == "loserRate" && val == "-1" { isUpdatable = false }
		if isUpdatable {
			statemnt := fmt.Sprintf("update %s set %s='%s' where roomId=%d",
									 tableName, key, val, roomInfo.roomId)
			_, err := db.Exec(statemnt)
			if err != nil { return err }
		}
	}

	return nil
}

