package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reptile-test-go/api/cmd"
	"strings"
)

func translation(comments *cmd.JDComment) {
	var text strings.Builder

	for i, comment := range (*comments).Comments {
		text.WriteString(comment.Content)
		if i != len((*comments).Comments)-1 {
			text.WriteString("|")
		}
	}

	result, _ := Translate(text.String(), "zh-CN", "en")
	com := strings.Split(result, "|")

	length := len((*comments).Comments)
	if len(com) < length {
		length = len(com)
	}

	for i := 0; i < len((*comments).Comments); i++ {
		//fmt.Println(i, comments.Comments[i].Content, com[i])
		(*comments).Comments[i].EnContent = com[i]
	}
}

func Translate(source, sourceLang, targetLang string) (string, error) {
	var text []string
	var result []interface{}

	encodedSource := url.PathEscape(source)
	url := "https://translate.googleapis.com/translate_a/single?client=gtx&sl=" + sourceLang + "&tl=" + targetLang + "&dt=t&q=" + encodedSource

	r, err := http.Get(url)
	if err != nil {
		return "err", errors.New("Error getting translate.googleapis.com")
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "err", errors.New("Error reading response body")
	}

	bReq := strings.Contains(string(body), `<title>Error 400 (Bad Request)`)
	if bReq {
		return "err", errors.New("Error 400 (Bad Request)")
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return "err", errors.New("Error unmarshaling data")
	}

	if len(result) > 0 {
		inner := result[0]
		for _, slice := range inner.([]interface{}) {
			for _, translatedText := range slice.([]interface{}) {
				text = append(text, fmt.Sprintf("%v", translatedText))
				break
			}
		}
		cText := strings.Join(text, "")

		return cText, nil
	} else {
		return "err", errors.New("No translated data in responce")
	}
}
