package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}

	// 在问候消息中.
	message := fmt.Sprintf(randomFormat(), name)
	// message := fmt.Sprint(randomFormat())

	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	msgMap := make(map[string]string)

	for _, name := range names {
		msg, err := Hello(name)
		if err != nil {
			return nil, err
		}

		msgMap[name] = msg
	}
	return msgMap, nil
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}
