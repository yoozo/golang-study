package main

import (
	"fmt"
	"greetings"
	"log"
)

func main() {
	log.SetPrefix("greetings:")
	log.SetFlags(0)

	// 一个名字切片.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// 请求姓名的问候消息.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	// 如果没有返回错误，则打印返回的map
	// 消息到控制台.
	fmt.Println(messages)

	fmt.Println("main: ", c())
}

func c() int {
	i := 1
	defer func() {
		i++
		fmt.Println("defer", i)
	}()
	return i
}
