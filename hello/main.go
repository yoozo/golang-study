package main

import (
	"fmt"
	"greetings"
	"log"
)

func main() {
	log.SetPrefix("greetings:")
	log.SetFlags(0)
	msg, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(msg)
}
