package handler

import (
	"crypto/sha256"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"reptile-test-go/cmd"
	"reptile-test-go/logic"
	"reptile-test-go/model"
)

type UserResponse struct {
	StatusCode int32  `json:"status_code"`          // 0为正常,其他为异常
	StatusMsg  string `json:"status_msg,omitempty"` // 传给前端的信息
	Id         int64  `json:"id,omitempty"`
	Token      string `json:"token,omitempty"`
	Value      any    `json:"value,omitempty"`
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(c.PostForm("password"))))
	log.Println("username:", username)

	user, err := sql.CheckLogin(username, password)
	if err != nil {
		c.JSON(http.StatusOK, cmd.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	token, _ := logic.GenerateToken(user.Id, username, password)

	c.JSON(http.StatusOK, UserResponse{
		StatusCode: 0,
		Token:      token,
		Id:         user.Id,
	})
}

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(c.PostForm("password"))))
	log.Println("username:", username)

	result := sql.CheckUserExist(username)
	if result {
		c.JSON(http.StatusOK, cmd.Response{
			StatusCode: 1,
			StatusMsg:  "该账号已存在",
		})
		return
	}

	user, err := sql.CreateUser(username, password)
	if err != nil {
		c.JSON(http.StatusOK, cmd.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, UserResponse{
		StatusCode: 0,
		Id:         user.Id,
	})
}

func Userinfo(c *gin.Context) {
	//token, _ := c.Cookie("token")
	token := c.Query("token")
	logic.Trim(&token)

	claims, err := logic.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusOK, cmd.Response{
			StatusCode: 2,
			StatusMsg:  err.Error(),
		})
		return
	}

	user, err := sql.FindUserById(claims.Id)
	if err != nil {
		c.JSON(http.StatusOK, cmd.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, UserResponse{
		StatusCode: 0,
		Value:      user,
	})
}
