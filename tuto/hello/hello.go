package main


import (
	"fmt"
	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Bod")
	fmt.Println(message)
}
