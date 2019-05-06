package main

import(
	"./IBM"
	"fmt"
)

func main() {
	result := IBM.GetSpeech("test.wav")
	fmt.Println(result)
}