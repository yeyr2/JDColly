package sql

import (
	"reptile-test-go/api"
)

type sqlComment struct {
	productId int64  `gorm:"column:id"`
	context   string `gorm:"column:context"`
	enContext string `gorm:"column:EnContext"`
	oldScore  int    `gorm:"column:oldScore"`
}

func (s sqlComment) TableName() string {
	return "comments"
}

func SaveComment(comments api.JDComment) {
	sqlCom := make([]sqlComment, 0, len(comments.Comments))

	for _, comment := range comments.Comments {
		if comment.UsefulVoteCount == 0 {
			continue
		}
		tmp := sqlComment{
			productId: comments.ProductCommentSummary.ProductID,
			context:   comment.Content,
			enContext: comment.EnContent,
			oldScore:  comment.Score,
		}
		sqlCom = append(sqlCom, tmp)
	}

	db.Create(&sqlCom)
}
