package greetings

import "fmt"

func Hello(name string) string {
	// 在问候消息中.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
