package auth

import (
	"testing"
	"github.com/stretchr/testify/assert"
  "github.com/upDJ/Image-Processing-Service/internal/secret"
)

func TestJwtSecretManger(t *testing.T) {
  envKey := "TestJWTSECRETKEY"
  jwtKey := "test-key"
  t.Setenv(envKey, jwtKey)
  
  var keyManager secret.SecretManager
  keyManager = secret.JwtSecretManager{}

  secretKey, err := keyManager.GetSecret("TestJWTSECRETKEY")

  assert.Equal(t, secretKey, jwtKey)
  assert.Equal(t, err, nil)
}
