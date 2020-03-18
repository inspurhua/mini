package middleware

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"huage.tech/mini/app/config"
	"huage.tech/mini/app/util"
	"time"
)

// JWTAuth 中间件，检查token
func JWTToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			queryToken := c.DefaultQuery("token", "")
			if queryToken == "" {
				util.AbortNewResultErrorOfClient(c, errors.New("未提供token"))
				return
			}
			token = queryToken
		}

		j := NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == TokenExpired {
				util.AbortNewResultErrorOfClient(c, errors.New("token已过期"))
				return
			}
			util.AbortNewResultErrorOfClient(c, errors.New("错误的token"))
			return
		}
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("claims", claims)
		c.Set("token", token)
		c.Set("UID", claims.ID)
		c.Set("ORG_ID", claims.Org)
		c.Set("ROLE_ID", claims.Role)
		c.Next()
	}
}

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

// 一些常量
var (
	TokenExpired     error  = errors.New("Token is expired")
	TokenNotValidYet error  = errors.New("Token not active yet")
	TokenMalformed   error  = errors.New("That's not even a token")
	TokenInvalid     error  = errors.New("Couldn't handle this token:")
	SignKey          string = config.JwtSecret
)

// 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	ID   int64 `json:"id"`
	Role int64 `json:"role"`
	Org  int64 `json:"org"`
	jwt.StandardClaims
}

// 新建一个jwt实例
func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

// 获取signKey
func GetSignKey() string {
	return SignKey
}

// 这是SignKey
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

// CreateToken 生成一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 解析Token
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

// 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(2 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}
