package main

import (
	"log"

	"github.com/mitchellh/mapstructure"
)

type User struct {
	ID   int    `mapstructure:"id"`
	Name string `mapstructure:"name"`
}

func main() {
	var mp map[string]interface{}
	u := &User{
		ID:   1,
		Name: "yukpiz",
	}

	err := mapstructure.Decode(u, &mp)
	if err != nil {
		panic(err)
	}
	log.Printf("%+v\n", mp)
}
