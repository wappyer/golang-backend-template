package jwt

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type Client struct {
	SigningKey []byte
}

var client *Client

func Initialize(signingKey []byte) {
	client = &Client{
		SigningKey: signingKey,
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
func (c *Client) GenerateToken(userId string) (string, error) {
	claims := Claims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 60*60*24*7, //7天的token
			Issuer:    "wappyer",
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
