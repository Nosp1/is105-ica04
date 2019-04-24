package main

import (
	"fmt"
	"github.com/Henrikohlsen/Is-105/go-speak/Wit"
)

func main(){
	speech.SetWitKey("AGBTMHDDRJUD7FBTFSY3MJLTJS4SEY73")
	print := speech.SendWitVoice("test.wav")
	fmt.Println("The text from the audio file: ", print)
}

