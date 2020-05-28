package main

import (
	"fmt"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// SQLConnect DB??
func sqlConnect() (*sql.DB, error) {
    DBMS := "mysql"
    USER := "root"
    PROTOCOL := "tcp(localhost:3306)"
    DBNAME := "mydb"

    CONNECT := USER + "@" + PROTOCOL + "/" + DBNAME // + "?charset=utf8&parseTime=true&loc=Asia%2FTokyo"
	db, err := sql.Open(DBMS, CONNECT)
    return db, err
}

func insert(db *sql.DB) (error) {
	_, err := db.Exec("insert into `user` values (999997, 'ファイターさん', 'mario', 1500, 'ファイターに', 'samus', 1600);")
	
	return err
}

func main() {
	db, err := sqlConnect()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = insert(db)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("success")
}


