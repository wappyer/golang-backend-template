package jwt

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type Client struct {
	SigningKey    []byte
	Issuer        string
	ExpiresSecond int64
}

var client *Client

func Initialize(signingKey, issuer string, expiresSecond int64) {
	client = &Client{
		SigningKey:    []byte(signingKey),
		Issuer:        issuer,
		ExpiresSecond: expiresSecond,
	}
}

func GetClientIns() *Client {
	return client
}

type Claims struct {
	LoginId string `json:"loginId"`
	jwt.StandardClaims
}

// GenerateToken 生成Token
func (c *Client) GenerateToken(loginId string) (string, error) {
	claims := Claims{
		loginId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + c.ExpiresSecond, //7天的token
			Issuer:    c.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(c.SigningKey)
}

func (c *Client) VerifyToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return c.SigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	if token != nil {
		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, err
}
