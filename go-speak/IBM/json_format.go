package IBM

import(
	"encoding/json"
	"fmt"
)

var result FormattedJSON

func JSONFormat(s string) (string, string) {
	json.Unmarshal([]byte(s), &result)

	return result.Results[0].Alternatives[0].Transcript,
		fmt.Sprintf("confidence: %f%%", result.Results[0].Alternatives[0].Confidence*100)
}

type FormattedJSON struct {
	Results []struct {
		Final        bool `json:"final"`
		Alternatives []struct {
			Transcript string          `json:"transcript"`
			Confidence float64         `json:"confidence"`
			Timestamps [][]interface{} `json:"timestamps"`
		} `json:"alternatives"`
		KeywordsResult struct {
			Hello []struct {
				NormalizedText string  `json:"normalized_text"`
				StartTime      float64 `json:"start_time"`
				EndTime        float64 `json:"end_time"`
				Confidence     float64 `json:"confidence"`
			} `json:"hello"`
		} `json:"keywords_result"`
		WordAlternatives []struct {
			StartTime    float64 `json:"start_time"`
			EndTime      float64 `json:"end_time"`
			Alternatives []struct {
				Confidence float64 `json:"confidence"`
				Word       string  `json:"word"`
			} `json:"alternatives"`
		} `json:"word_alternatives"`
	} `json:"results"`
	ResultIndex int `json:"result_index"`
}
