package security

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type Claims struct {
	UserId int64 `json:"usr"`
	jwt.StandardClaims
}

// SignJwt generate tokens used for auth
func SignJwt(userId int64, issuer string, expiresIn time.Duration, secret []byte) (string, error) {
	nowTime := time.Now()
	expiresAt := nowTime.Add(expiresIn)

	claims := Claims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: expiresAt.Unix(),
			Issuer:    issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(secret)

	return token, err
}

// ParseJwt parsing token
func ParseJwt(token string, secret []byte) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func IsJwtExpired(claims *Claims) bool {
	return claims.ExpiresAt <= time.Now().Unix()
}
