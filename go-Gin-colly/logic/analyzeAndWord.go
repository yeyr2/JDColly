package logic

import (
	"github.com/jonreiter/govader"
	"math"
	"reptile-test-go/cmd"
	"reptile-test-go/config"
	"reptile-test-go/model"
)

// 按照出现频率排序
type kv struct {
	Key   string
	Value int
}

func GetCommentBySql(id string, startTime, lastTime int64) *[]cmd.Comments {
	comments := sql.GetComments(id, startTime, lastTime)

	return comments
}
func AnalyzeGetComments(comment *[]cmd.Comments, analyze *cmd.AnalyzeComment, atype string, flag chan bool) {
	if len(*comment) == 0 {
		flag <- false
		return
	}

	if atype == "Chinese NLP" {
		flag <- AnalyzeGetCommentsNLP(comment, analyze)
		return
	}
	if atype == "JieBa" {
		flag <- AnalyzeGetCommentsJieBa(comment, analyze)
		return
	}
	flag <- false
}

func WordCloudAnalysis(comment *[]cmd.Comments, analyze *cmd.AnalyzeComment, id string, result chan bool) {
	if len(*comment) == 0 {
		analyze.AnalyzeWord = ""
		result <- true
		return
	}

	response := wordCloudRpc(comment, id)
	response = "http://" + config.Host + "/wordcloud/" + response
	analyze.AnalyzeWord = response

	result <- true
}

func AnalyzeGetCommentsJieBa(comment *[]cmd.Comments, analyze *cmd.AnalyzeComment) bool {
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

	for i := 0; i < 5; i++ {
		if count[i] != 0 {
			sum[i] = sum[i] * 50
			sums += int(sum[i])
			if sum[i]-math.Floor(sum[i]) > 0.5 {
				sum[i] += 1
			}
			counts += count[i]
			(*analyze).Interval[i] = int32(count[i])
		}
	}
	(*analyze).Fraction = int32(sums / counts)

	return true
}

func AnalyzeGetCommentsNLP(comment *[]cmd.Comments, analyze *cmd.AnalyzeComment) bool {
	if len(*comment) == 0 {
		analyze.AnalyzeWord = ""
		return false
	}

	fraction, interval := AnalysisByNLPRpc(comment)
	analyze.Interval = interval
	analyze.Fraction = fraction
	return true
}
