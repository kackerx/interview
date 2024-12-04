package middleware

import (
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"github.com/kackerx/interview/common/code"
	"github.com/kackerx/interview/common/log"
	"github.com/kackerx/interview/common/resp"
	"github.com/kackerx/interview/internal/conf"
)

type JWT struct {
	key []byte
}

type MyCustomClaims struct {
	UserId string
	jwt.RegisteredClaims
}

func NewJwt(conf *conf.Conf) *JWT {
	return &JWT{key: []byte(conf.Security.Jwtkey)}
}

func (j *JWT) GenToken(userId string, expiresAt time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyCustomClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "",
			Subject:   "",
			ID:        "",
			Audience:  []string{},
		},
	})

	tokenString, err := token.SignedString(j.key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j *JWT) ParseToken(tokenString string) (*MyCustomClaims, error) {
	re := regexp.MustCompile(`(?i)Bearer `)
	tokenString = re.ReplaceAllString(tokenString, "")
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.key, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func StrictAuth(j *JWT) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.Request.Header.Get("Authorization")
		if tokenString == "" {
			log.New(ctx).Warn("请求未携带token，无权限访问", "data", map[string]interface{}{
				"url":    ctx.Request.URL,
				"params": ctx.Params,
			})
			resp.HandleErr(ctx, code.ErrUserTokenNotFound)
			ctx.Abort()
			return
		}

		claims, err := j.ParseToken(tokenString)
		if err != nil {
			log.New(ctx).Error("token error", "data", map[string]interface{}{
				"url":    ctx.Request.URL,
				"params": ctx.Params,
			})
			resp.HandleErr(ctx, code.ErrUserTokenInvalid.WithArgs(tokenString))
			ctx.Abort()
			return
		}

		ctx.Set("claims", claims)
		recoveryLoggerFunc(ctx)
		ctx.Next()
	}
}

func recoveryLoggerFunc(ctx *gin.Context) {
	userInfo := ctx.MustGet("claims").(*MyCustomClaims)
	log.New(ctx).Info("request log", "userID", userInfo.UserId)
}
