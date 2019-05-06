// URL	   https://gateway-lon.watsonplatform.net/speech-to-text/api
package IBM

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/go-sdk-core/core"
	"github.com/watson-developer-cloud/go-sdk/speechtotextv1"
	"io"
	"os"
)

const url = "https://gateway-lon.watsonplatform.net/speech-to-text/api"

func GetSpeech(inputFile string, key string) string {

	var resultString string

	speechToText, speechToTextErr :=
		speechtotextv1.NewSpeechToTextV1(&speechtotextv1.SpeechToTextV1Options{
			IAMApiKey: key,
			URL:       url,
		})
	if speechToTextErr != nil {
		panic(speechToTextErr)
	}

	files := [1]string{inputFile}
	for _, file := range files {
		var audioFile io.ReadCloser
		var audioFileErr error
		audioFile, audioFileErr = os.Open("./" + file)
		if audioFileErr != nil {
			panic(audioFileErr)
		}
		response, responseErr := speechToText.Recognize(
			&speechtotextv1.RecognizeOptions{
				Audio:						&audioFile,
				ContentType:
					core.StringPtr(speechtotextv1.RecognizeOptions_ContentType_AudioWav),
				Timestamps:					core.BoolPtr(true),
				WordAlternativesThreshold:	core.Float32Ptr(0.9),
				Keywords: 					[]string{"hello", "group"},
				KeywordsThreshold:			core.Float32Ptr(0.5),

			},
		)
		if responseErr != nil {
			panic(responseErr)
		}

		result := speechToText.GetRecognizeResult(response)
		b, _ := json.MarshalIndent(result, "", "  ")
		resultString = string(b)
	}

	fmt.Println(files)

	return resultString
}
