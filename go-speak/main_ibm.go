package main

import(
	"./IBM"
	"fmt"
)

func main() {

	result := IBM.GetSpeech("test.wav", "Use your own API key here")
	fmt.Println(result)
}