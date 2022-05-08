package security

import (
	"douyin/app/errs"
	"douyin/pkg/util"
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
	ignoreRoutes = make(map[string]bool)
)

func Setup(jwtConfig JwtConfig) {
	config = jwtConfig
}

func Bind(engine *gin.Engine, ignore ...string) {
	for _, val := range ignore {
		ignoreRoutes[val] = true
	}
	engine.Use(securityMiddleware)
}

func GenerateJwt(userId int64) (string, error) {
	return util.GenerateJwt(userId, config.Issuer,
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
	if ignoreRoutes[ctx.FullPath()] {
		return
	}
	// get token from query
	var token = ctx.Query(tokenKey)
	if token == "" {
		// get bearer token if query token doesn't exist
		token = util.GetBearerToken(ctx)
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, "bearer token not found")
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
	jwt, err := util.ParseJwt(token, []byte(config.Secret))
	if err != nil {
		return 0, err
	}
	if util.IsJwtExpired(jwt) {
		return 0, errs.JwtExpired
	}
	return jwt.UserId, nil
}
