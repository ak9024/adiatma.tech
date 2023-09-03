package env_test

import (
	"os"
	"testing"

	"github.com/ak9024/adiatma.tech/api/internal/app/env"
	"github.com/stretchr/testify/assert"
)

func TestGenerateEnv(t *testing.T) {
	fileName := ".env.test"
	os.Setenv("PORT", "8001")
	errGenerateEnv := env.GenerateEnv(fileName)
	assert.Nil(t, errGenerateEnv)
}
