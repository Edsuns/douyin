package security

import (
	"douyin/pkg/com"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type JwtConfig struct {
	Issuer    string
	Secret    string
	ExpiresIn time.Duration `yaml:"expires-in"`
}

const (
	tokenKey  = "token"
	userIdKey = "userId"
)

var (
	config JwtConfig

	ErrJwtExpired    = errors.New("jwt expired")
	ErrTokenRequired = errors.New("token required")
)

func Setup(jwtConfig JwtConfig) {
	config = jwtConfig
}

func GenerateJwt(userId int64) (string, error) {
	return SignJwt(userId, config.Issuer,
		config.ExpiresIn, []byte(config.Secret))
}

func GetUserId(ctx *gin.Context) int64 {
	if id, ok := ctx.Get(userIdKey); ok {
		return id.(int64)
	}
	panic(fmt.Sprintf("GetUserId from unguarded route: %s", ctx.FullPath()))
}

// Middleware filters unauthorized requests
func Middleware(ctx *gin.Context) {
	// ignore unmatched routes
	if ctx.FullPath() == "" {
		return
	}
	userId, err := getUserIdFromContext(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, com.Response{
			StatusCode: http.StatusUnauthorized,
			StatusMsg:  err.Error(),
		})
		return
	}
	// attach user id to the context
	ctx.Set(userIdKey, userId)
}

// getUserIdFromContext returns user id parsed from token, or error if there is not a valid token
func getUserIdFromContext(ctx *gin.Context) (int64, error) {
	// get token from query
	var token = ctx.Query(tokenKey)
	if token == "" {
		// get token from form
		token = ctx.PostForm(tokenKey)
		if token == "" {
			// get bearer token
			token = GetBearerToken(ctx)
			if token == "" {
				return 0, ErrTokenRequired
			}
		}
	}
	// verify token and get user id
	userId, err := getUserIdFromToken(token)
	if err != nil {
		return 0, err
	}
	return userId, nil
}

// getUserIdFromToken verifies token and returns user id
func getUserIdFromToken(token string) (int64, error) {
	jwt, err := ParseJwt(token, []byte(config.Secret))
	if err != nil {
		return 0, err
	}
	if IsJwtExpired(jwt) {
		return 0, ErrJwtExpired
	}
	return jwt.UserId, nil
}
