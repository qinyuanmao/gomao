package security

import (
	"fmt"
	"strconv"
	"time"

	"github.com/pkg/errors"

	jwt "github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	jwt.RegisteredClaims
}

// GetUserId 获取 userId
func (c *Claims) GetUserId() (int64, error) {
	userId, err := strconv.ParseInt(c.ID, 10, 64)
	if err != nil {
		return 0, errors.New("parse userId error")
	}
	return userId, nil
}

func CreateClaims(userId int64, exp time.Time) Claims {
	return Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{Time: exp},
			ID:        fmt.Sprintf("%d", userId),
		},
	}
}

// jwt 加密
func (p *Parser) JwtEncode(claims Claims) (token string, err error) {
	t := jwt.New(jwt.GetSigningMethod("RS256"))
	t.Claims = claims
	return t.SignedString(p.privateKey)
}

// jwt 解密
func (p *Parser) JwtDecode(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return p.publicKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims.(*Claims), nil
}
