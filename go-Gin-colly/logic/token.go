package logic

import (
	"github.com/dgrijalva/jwt-go"
	"reptile-test-go/cmd"
	"strings"
	"time"
)

var jwtSecret = []byte("iyeyr2.Token")

// GenerateToken 根据用户的用户名和密码产生token
func GenerateToken(id int64, username, password string) (string, error) {
	//设置token有效时间
	nowTime := time.Now()
	expireTime := nowTime.Add(6 * time.Hour)

	claims := cmd.Claims{
		Id:       id,
		Username: username,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: expireTime.Unix(),
			// 指定token发行人
			Issuer: "yeyr2",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//该方法内部生成签名字符串，再用于获取完整、已签名的token
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// ParseToken 根据传入的token值获取到Claims对象信息，（进而获取其中的用户名和密码）
func ParseToken(token string) (*cmd.Claims, error) {

	//用于解析鉴权的声明，方法内部主要是具体的解码和校验的过程，最终返回*Token
	tokenClaims, err := jwt.ParseWithClaims(token, &cmd.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		// 从tokenClaims中获取到Claims对象，并使用断言，将该对象转换为我们自己定义的Claims
		// 要传入指针，项目中结构体都是用指针传递，节省空间。
		if claims, ok := tokenClaims.Claims.(*cmd.Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

// Substring 从开头截取token的长为size内容
func Substring(token string, size int) string {
	var str strings.Builder

	for i, x := range token {
		if i == size {
			break
		}
		str.WriteByte(byte(x))
	}

	return str.String()
}

func Trim(token *string) {
	length := len(*token)
	if length == 0 {
		return
	}
	if (*token)[0] == '"' {
		*token = strings.TrimLeft(*token, "\"")
	}
	if (*token)[length-2] == '"' {
		*token = strings.TrimRight(*token, "\"")
	}
}
