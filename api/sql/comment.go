package sql

import (
	"log"
	"reptile-test-go/api/cmd"
	"time"
)

type sqlComment struct {
	ProductId       int64  `gorm:"column:product_id"`
	Context         string `gorm:"column:context"`
	EnContext       string `gorm:"column:en_context"`
	OldScore        int    `gorm:"column:old_score"`
	ReferenceTime   int64  `gorm:"colum:reference_time"`
	UsefulVoteCount int    `gorm:"colum:useful_vote_count"`
}

func (s sqlComment) TableName() string {
	return "comments"
}

func CommentsLastTime(productId int64) (lastTime int64) {
	var lastComment = new(sqlComment)
	result := db.Where("product_id = ?", productId).Order("reference_time DESC").Find(lastComment)
	if result.RowsAffected != 0 {
		lastTime = lastComment.ReferenceTime
	}
	return lastTime
}

func SaveComment(comments cmd.JDComment, lastTime int64) {
	if len(comments.Comments) == 0 {
		return
	}
	sqlCom := make([]*sqlComment, 0, len(comments.Comments))

	for _, comment := range comments.Comments {
		tmpReferenceTime, _ := time.Parse("2006-01-02 15:04:05", comment.ReferenceTime)
		referenceTime := tmpReferenceTime.Unix()

		if referenceTime <= lastTime {
			continue
		}

		tmp := &sqlComment{
			ProductId:       comments.ProductCommentSummary.ProductID,
			Context:         comment.Content,
			EnContext:       comment.EnContent,
			OldScore:        comment.Score,
			ReferenceTime:   referenceTime,
			UsefulVoteCount: comment.UsefulVoteCount,
		}
		sqlCom = append(sqlCom, tmp)
	}
	res := db.Create(sqlCom)
	if res.RowsAffected == 0 {
		log.Println(res.Error)
	}
}

func GetComments(id string, lastTime int64) *[]cmd.Comments {
	sqlComments := make([]sqlComment, 0)

	db.Where("reference_time > ? and product_id = ?", lastTime, id).Find(&sqlComments)

	comments := make([]cmd.Comments, len(sqlComments))
	for i, comment := range sqlComments {
		comments[i].UsefulVoteCount = comment.UsefulVoteCount
		comments[i].Score = comment.OldScore
		comments[i].EnContent = comment.EnContext
		comments[i].Content = comment.Context
	}

	return &comments
}
