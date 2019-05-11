package main

import (
	"fmt"
	"github.com/Henrikohlsen/Is-105/is105-ica04/go-speak/Wit"
)

func main(){
	speech.SetWitKey(UseYourWitKeyHere)
	print := speech.SendWitVoice("test.wav")
	fmt.Println("The text from the audio file: ", print)
}

