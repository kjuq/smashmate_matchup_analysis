package main

import (
	"reflect"
	"strconv"
	"fmt"
)

type Player struct {
	name       string
	fighter    string
	rate       int
	isCanceled bool
}

type RoomInfo struct {
	roomId int
	winnerName string
	winnerFighter string
	winnerRate int
	loserName string
	loserFighter string
	loserRate int
}

func playerToDict(p Player) map[string]string {
	result := map[string]string{}
	result["name"] = p.name
	result["fighter"] = p.fighter
	result["rate"] = strconv.Itoa(p.rate)

	return result
}

func structToDict(strct interface{}) map[string]string {
	dict := map[string]string{}

	fields := reflect.TypeOf(strct)
	values := reflect.ValueOf(strct)

	for i := 0; i < fields.NumField(); i++ {
		field := fields.Field(i).Name
		value := fmt.Sprint(values.Field(i))
		dict[field] = value
	}

	return dict

}

