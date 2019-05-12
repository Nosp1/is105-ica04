package main

import (
	"fmt"
	"./Wit"
)

func main(){
	speech.SetWitKey("AGBTMHDDRJUD7FBTFSY3MJLTJS4SEY73")
	print, confidence := speech.SendWitVoice("test.wav")
	fmt.Println("The text from the audio file: ", print)
	fmt.Println(confidence)
}

