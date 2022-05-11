package security

import (
	"douyin/pkg/com"
	"errors"
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
	config       JwtConfig
	ignoreRoutes = make(map[string]struct{})

	JwtExpired = errors.New("jwt expired")
)

func Setup(jwtConfig JwtConfig) {
	config = jwtConfig
}

func Bind(engine *gin.Engine, ignore ...string) {
	for _, val := range ignore {
		ignoreRoutes[val] = struct{}{}
	}
	engine.Use(securityMiddleware)
}

func GenerateJwt(userId int64) (string, error) {
	return SignJwt(userId, config.Issuer,
		config.ExpiresIn, []byte(config.Secret))
}

func GetUserId(ctx *gin.Context) int64 {
	return ctx.GetInt64(userIdKey)
}

// securityMiddleware filters unauthorized requests
func securityMiddleware(ctx *gin.Context) {
	// ignore unmatched routes
	if ctx.FullPath() == "" {
		return
	}
	// ignore routes that configured in ignoreRoutes
	if _, contains := ignoreRoutes[ctx.FullPath()]; contains {
		return
	}
	// get token from query
	var token = ctx.Query(tokenKey)
	if token == "" {
		// get bearer token if query token doesn't exist
		token = GetBearerToken(ctx)
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, com.Response{
				StatusCode: http.StatusUnauthorized,
				StatusMsg:  "token required",
			})
			return
		}
	}
	// verify token and get user id
	userId, err := getUserIdFromToken(token)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// attach user id to context
	ctx.Set(userIdKey, userId)
}

// getUserIdFromToken verifies token and returns user id
func getUserIdFromToken(token string) (int64, error) {
	jwt, err := ParseJwt(token, []byte(config.Secret))
	if err != nil {
		return 0, err
	}
	if IsJwtExpired(jwt) {
		return 0, JwtExpired
	}
	return jwt.UserId, nil
}
