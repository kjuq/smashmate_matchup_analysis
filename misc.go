package main

import (
	"strconv"
)

type Player struct {
	name       string
	fighter    string
	rate       int
	isCanceled bool
}

func playerToDict(p Player) map[string]string {
	result := map[string]string{}
	result["name"] = p.name
	result["fighter"] = p.fighter
	result["rate"] = strconv.Itoa(p.rate)

	return result
}

