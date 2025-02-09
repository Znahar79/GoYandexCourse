package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Person struct {
	Name        string `json:"Имя"`
	Email       string `json:"Почта"`
	DateOfBirth time.Time
}

func main() {
	man := Person{
		Name:        "Alex",
		Email:       "alex@yandex.ru",
		DateOfBirth: time.Now(),
	}

	result, err := json.Marshal(man)
	if err != nil {
		log.Fatalln("unable marshal to json")
	}
	fmt.Println(string(result))
}
