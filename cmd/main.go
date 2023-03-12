package main

import (
	"fmt"

	"github.com/tyleryep/goat/pkg/helloworld"
)

func main() {
	result, err := helloworld.SayHello("Tyler")
	if err != nil {
		fmt.Println("Error!")
	}
	fmt.Println(result)
}
