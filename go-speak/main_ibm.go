package main

import(
	"./IBM"
	"fmt"
)

func main() {

	result := IBM.GetSpeech("test.wav", "Use your own key here")
	formattedResult, confidence := IBM.JSONFormat(result)
	fmt.Println(formattedResult)
	fmt.Println(confidence)
}