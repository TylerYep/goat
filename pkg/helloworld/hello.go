package helloworld

import "fmt"

func SayHello(name string) (string, error) {
	result := fmt.Sprintf("Hello %s!", name)
	// fmt.Println(result)
	return result, nil
}
