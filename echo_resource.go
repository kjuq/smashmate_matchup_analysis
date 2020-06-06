package main

import (
	"github.com/labstack/echo"
	"strconv"
)

func parseRoomInfoFromParams(c echo.Context) (RoomInfo, error) {

	nonRoom := RoomInfo{roomId: -1}

	var winnerRate, loserRate int
	var err error
	if rateStr := c.FormValue("winnerRate"); rateStr != "" {
		winnerRate, err = strconv.Atoi(c.FormValue("winnerRate"))
		if err != nil { return nonRoom, err }
	} else {
		winnerRate = -1
	}
	if rateStr := c.FormValue("loserRate"); rateStr != "" {
		loserRate, err = strconv.Atoi(c.FormValue("loserRate"))
		if err != nil { return nonRoom, err }
	} else {
		loserRate = -1
	}

	roomInfo := RoomInfo{}
	roomInfo.winnerName = c.FormValue("winnerName")
	roomInfo.winnerFighter = c.FormValue("winnerFighter")
	roomInfo.winnerRate = winnerRate
	roomInfo.loserName = c.FormValue("loserName")
	roomInfo.loserFighter = c.FormValue("loserFighter")
	roomInfo.loserRate = loserRate
	
	return roomInfo, nil
}


