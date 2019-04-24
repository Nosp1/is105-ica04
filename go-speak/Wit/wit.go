package speech

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var witKey string

type SpeechToText struct {
	Text     string `json:"_text"`
	Outcomes []struct {
		Confidence interface{} `json:"confidence"`
		Intent     string      `json:"intent"`
		Text       string      `json:"_text"`
		Entities   struct {
		} `json:"entities"`
	} `json:"outcomes"`
	WARNING string `json:"WARNING"`
	MsgID   string `json:"msg_id"`
}

/** SetWitKey
*   witKey must be set prior to executing any wit commands
**/
func SetWitKey(key string) string {
	witKey = key
	return witKey
}

/** PrintWitKey
*   Returns the current wit key if set, otherwise returns nil
**/
func PrintWitKey() string {
	return witKey
}

/** convert
* converts a message with spaces into one suitable to passing to wit
**/

func convert(message string) string {
	arrString := strings.Split(message, " ")
	var ret string
	for x := 0; x < len(arrString); x++ {
		ret += arrString[x] + "%20"
	}
	return ret
}

func SendWitMessage(message string) string {
	url := "https://api.wit.ai/message?v=20160225&q=" + convert(message)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer "+witKey)
	client := &http.Client{}
	resp, _ := client.Do(req)
	contents, _ := ioutil.ReadAll(resp.Body)
	return string(contents)
}

/**
*Sends an audio file to wit.ai, wit key must have been set prior to calling
*@param filename the full path to the file that is to be sent
*@return a string with the json data received
**/

func SendWitVoice(fileRef string) string {
	audio, err := ioutil.ReadFile(fileRef)
	if err != nil {
		log.Fatal("Error reading file:\n%v\n", err)

	}

	reader := bytes.NewReader(audio)

	url := "https://api.wit.ai/speech?v=20141022"
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, reader)

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+witKey)
	req.Header.Set("Content-Type", "audio/wav")
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, _ := ioutil.ReadAll(res.Body)

	var speechtotext SpeechToText
	jsontxt := json.Unmarshal(body, &speechtotext)
	if jsontxt != nil {
		fmt.Println("There was an error:", err)
	}

	return string(speechtotext.Text)
}
func SendWitBuff(buffer *bytes.Buffer) string {
	url := "https://api.wit.ai/speech?v=20141022"
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, buffer)

	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Authorization", "Bearer "+witKey)
	req.Header.Set("Content-Type", "audio/wav")
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}
