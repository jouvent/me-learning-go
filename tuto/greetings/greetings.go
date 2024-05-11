package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return "", errors.New("empty name")
	}
	// If a message was received, return a value that embeds the name
	// in a greeting message.
	message := fmt.Sprintf("Hi, %v. Wellcome!", name)
	return message, nil
}
