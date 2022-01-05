package security_test

import (
	"cargo-rest-api/pkg/security"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateSecret(t *testing.T) {
	_, genErr := security.GenerateSecret()
	assert.Nil(t, genErr)
}
