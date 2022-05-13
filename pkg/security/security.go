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

	ErrJwtExpired    = errors.New("jwt expired")
	ErrTokenRequired = errors.New("token required")
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
	// if there is a valid token, attach it to the context
	// otherwise, abort the request if the route are not configured in ignoreRoutes
	_, ignore := ignoreRoutes[ctx.FullPath()]
	userId, err := getUserIdFromContext(ctx)
	if err != nil && !ignore {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, com.Response{
			StatusCode: http.StatusUnauthorized,
			StatusMsg:  err.Error(),
		})
		return
	}
	if err == nil {
		// attach user id to the context
		ctx.Set(userIdKey, userId)
	}
}

// getUserIdFromContext returns user id parsed from token, or error if there is not a valid token
func getUserIdFromContext(ctx *gin.Context) (int64, error) {
	// get token from query
	var token = ctx.Query(tokenKey)
	if token == "" {
		// get bearer token if query token doesn't exist
		token = GetBearerToken(ctx)
		if token == "" {
			return 0, ErrTokenRequired
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
