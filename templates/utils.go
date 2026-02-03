package templates

import "fmt"

func CommonGo() string {
	return `package utils

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

const (
	CodeOK    = 0
	MessageOK = "success"
)

// MD5 calculates MD5 hash of a string
func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// TodayLastSecond returns remaining seconds until end of today
func TodayLastSecond() time.Duration {
	now := time.Now()
	tomorrow := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
	return tomorrow.Sub(now)
}

// InArray checks if an element exists in array
func InArray[T comparable](item T, arr []T) bool {
	for _, v := range arr {
		if v == item {
			return true
		}
	}
	return false
}
`
}

func JwtGo(moduleName string) string {
	return fmt.Sprintf(`package utils

import (
	"time"

	"%s/conf"
	"github.com/golang-jwt/jwt/v5"
)

type JWTClaims struct {
	UID string `+"`json:\"uid\"`"+`
	jwt.RegisteredClaims
}

// GenerateToken generates a JWT token
func GenerateToken(uid string) (string, error) {
	claims := JWTClaims{
		UID: uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(conf.Conf.JWT.Expire) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(conf.Conf.JWT.Secret))
}

// ParseToken parses a JWT token
func ParseToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.Conf.JWT.Secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}
`, moduleName)
}
