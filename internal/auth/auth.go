package auth

import (
  "log"
  "time"
  "github.com/upDJ/Image-Processing-Service/internal/secret"
  "github.com/golang-jwt/jwt/v5"
)

func SignToken(token *jwt.Token) (string, error) {
  var keyManager secret.SecretManager
  keyManager = secret.JwtSecretManager{}

  secretKey, err := keyManager.GetSecret("JWTSECRETKEY")
  tokenString, err := token.SignedString(secretKey)

  if err != nil {
    return "", err
  }
  return tokenString, nil
}

func getRole(username string) string {
  if username == "dj" {
    return "admin"
  }

  return "user"
}

func MapClaims(username string) jwt.Claims {
  claimsMap := jwt.MapClaims{
    "sub": username,
    "iss": "image-processing-service",                 
		"aud": getRole(username),           
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),                 
	}

  return claimsMap
}

func CreateToken(username string) (string, error) {
  claimsMap := MapClaims(username)
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsMap)
  log.Printf("Token claims added: %+v\n", token)
  tokenString, err := SignToken(token)
  
  return tokenString, err
}

