package main

import (
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"hello": "world"})
	})

	e.GET("/winner", func(c echo.Context) error {
		db, errSql := sqlConnect()
		if errSql != nil { panic(errSql) }
		defer db.Close()

		winner, err := getWinnerInfo(db, 5)
		if err != nil { panic(err) }

		result := playerToDict(winner)

		return c.JSON(http.StatusOK, result)
	})

	e.Logger.Fatal(e.Start(":1313"))
}

