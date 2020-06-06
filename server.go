package main

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"fmt"
)

func main() {
	e := echo.New()

	e.GET("/api", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"hello": "world"})
	})

	e.GET("/api/room/:roomId", func(c echo.Context) error {
		roomId, err := strconv.Atoi(c.Param("roomId"))
		if err != nil { panic(err) }
		db, err := sqlConnect()
		if err != nil { panic(err) }
		defer db.Close()

		roomInfo, err := getRoomInfo(db, roomId)
		if err != nil { panic(err) }

		result := structToDict(roomInfo)

		return c.JSON(http.StatusOK, result)
	})

	e.POST("/api/room", func(c echo.Context) error {
		db, errSql := sqlConnect()
		if errSql != nil { panic(errSql) }
		defer db.Close()

		roomId, errStr := strconv.Atoi(c.FormValue("roomId"))
		if errStr != nil { panic(errStr) }

		roomInfo, err := parseRoomInfoFromParams(c)
		if err != nil { panic(err) }
		roomInfo.roomId = roomId

		insertRoomInfo(db, roomInfo)

		result := structToDict(roomInfo)
		return c.JSON(http.StatusOK, result)
	})

	e.DELETE("/api/room/:roomId", func(c echo.Context) error {
		roomId, err := strconv.Atoi(c.Param("roomId"))
		if err != nil { panic(err) }
		db, err := sqlConnect()
		if err != nil { panic(err) }
		defer db.Close()

		err = deleteRoomInfo(db, roomId)
		if err != nil { panic(err) }

		return c.JSON(http.StatusOK, map[string]string{})
	})

	e.PUT("/api/room/:roomId", func(c echo.Context) error {
		roomId, err := strconv.Atoi(c.Param("roomId"))
		if err != nil { panic(err) }
		db, err := sqlConnect()
		if err != nil { panic(err) }
		defer db.Close()

		roomInfo, err := parseRoomInfoFromParams(c)
		if err != nil { panic(err) }
		roomInfo.roomId = roomId

		err = updateRoomInfo(db, roomInfo)
		if err != nil { panic(err) }

		roomInfo, err = getRoomInfo(db, roomId)
		if err != nil { panic(err) }

		return c.JSON(http.StatusOK, roomInfo)
	})

	e.POST("/debug", func(c echo.Context) error {
		fmt.Println(c.FormValue("roomId"))
		fmt.Println(c.FormValue("winnerName"))
		fmt.Println(c.FormValue("winnerFighter"))
		fmt.Println(c.FormValue("winnerRate"))
		fmt.Println(c.FormValue("loserName"))
		fmt.Println(c.FormValue("loserFighter"))
		fmt.Println(c.FormValue("loserRate"))

		return c.JSON(http.StatusOK, map[string]string{})
	})

	e.Start(":1313")
}


