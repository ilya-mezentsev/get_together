package code

import (
  "fmt"
  "github.com/dgrijalva/jwt-go"
  "utils"
)

type Coder struct {
  secret string
}

func NewCoder(key string) Coder {
  return Coder{secret: utils.GetHash(key)}
}

func (c Coder) Encrypt(tokenData map[string]interface{}) (string, error) {
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims(tokenData))

  return token.SignedString([]byte(c.secret))
}

func (c Coder) Decrypt(tokenString string) (map[string]interface{}, error) {
  token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
      return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
    }

    return []byte(c.secret), nil
  })
  if err != nil {
    return nil, err
  } else if token == nil {
    return nil, CannotParseToken
  }

  if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
    return claims, nil
  }

  return nil, CannotParseToken
}               
