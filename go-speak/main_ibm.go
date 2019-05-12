package main

import (
	"./IBM"
	"fmt"
	"os"
)

func main() {

	args := os.Args
	if len(args) != 2 {
		panic("Remember to add the authentication key.")
	}

	key := args[1]
	result := IBM.GetSpeech("test.wav", key)
	formattedResult, confidence := IBM.JSONFormat(result)
	fmt.Println(formattedResult)
	fmt.Println(confidence)
}