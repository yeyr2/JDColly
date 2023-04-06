package api

import (
	"fmt"
	"github.com/jonreiter/govader"
	"math"
	"reptile-test-go/api/cmd"
	"reptile-test-go/api/sql"
	"reptile-test-go/config"
)

// 按照出现频率排序
type kv struct {
	Key   string
	Value int
}

func GetCommentBySql(id string, lastTime int64) *[]cmd.Comments {
	comments := sql.GetComments(id, lastTime)

	return comments
}

func AnalyzeGetComments(comment *[]cmd.Comments, analyze *cmd.AnalyzeComment) bool {
	if len(*comment) == 0 {
		return false
	}
	var sum [5]float64
	var count [5]int
	counts, sums := 0, 0

	sia := govader.NewSentimentIntensityAnalyzer()
	for _, text := range *comment {
		scores := sia.PolarityScores(text.EnContent)
		if scores.Compound < -0.6 {
			sum[0] += scores.Compound + 1
			count[0]++
		} else if scores.Compound < -0.2 {
			sum[1] += scores.Compound + 1
			count[1]++
		} else if scores.Compound < 0.2 {
			sum[2] += scores.Compound + 1
			count[2]++
		} else if scores.Compound < 0.6 {
			sum[3] += scores.Compound + 1
			count[3]++
		} else if scores.Compound < 1.0 {
			sum[4] += scores.Compound + 1
			count[4]++
		}
	}

	(*analyze).Interval = make([]cmd.Interval, 5)
	for i := 0; i < 5; i++ {
		(*analyze).Interval[i].ScoreRange = fmt.Sprintf("[%d~%d]", i*20, 20*(i+1))
		if count[i] != 0 {
			sum[i] = sum[i] * 50
			sums += int(sum[i])
			if sum[i]-math.Floor(sum[i]) > 0.5 {
				sum[i] += 1
			}
			counts += count[i]
			(*analyze).Interval[i].Interval = count[i]
		}
	}
	(*analyze).Fraction = sums / counts

	return true
}

func WordCloudAnalysis(comment *[]cmd.Comments, analyze *cmd.AnalyzeComment, id string) {
	if len(*comment) == 0 {
		analyze.AnalyzeWord = ""
		return
	}

	response := wordCloudRpc(comment, id)
	response = "http://" + config.Host + "/wordcloud/" + response
	analyze.AnalyzeWord = response
	//fmt.Println(response)
}
