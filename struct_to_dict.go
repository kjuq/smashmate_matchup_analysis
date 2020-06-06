package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age int
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

func main() {
	john := Person{"john", 5}
	a := structToDict(john)
	fmt.Println(a)
}

