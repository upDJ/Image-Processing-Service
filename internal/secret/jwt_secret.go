package secret

import (
  "os"
  "errors"
)

type JwtSecretManager struct {}

func (jsm JwtSecretManager) GetSecret(key string) (string, error) {
  jwtKey := os.Getenv(key)
  if jwtKey == "" {
    return "", errors.New("environment variable not set: " + key)
  }
  return jwtKey, nil
}
