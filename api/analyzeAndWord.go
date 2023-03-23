package api

import (
	"fmt"
	gt "github.com/bas24/googletranslatefree"
	"github.com/jonreiter/govader"
)

func GetWordCloudAndAnalyzeRating(analyze *AnalyzeComment, id string, lastTime int64) {

}

func AnalyzeGetComments(comment *JDComment, analyze *AnalyzeComment) {
	var enString []string
	for _, text := range comment.Comments {
		result, _ := gt.Translate(text.Content, "zh-CN", "en")
		enString = append(enString, result)
	}

	sia := govader.NewSentimentIntensityAnalyzer()
	for _, text := range enString {
		scores := sia.PolarityScores(text)
		fmt.Println("for line:", text, "score:", scores)
		fmt.Println("Compound score:", scores.Compound)
		fmt.Println("Positive score:", scores.Positive)
		fmt.Println("Neutral score:", scores.Neutral)
		fmt.Println("Negative score:", scores.Negative)
	}
}

func WordCloudAnalysis(comment *JDComment, analyze *AnalyzeComment) {

}
