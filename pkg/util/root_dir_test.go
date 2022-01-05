package util_test

import (
	"cargo-rest-api/pkg/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootDir(t *testing.T) {
	rootDir := util.RootDir()
	assert.NotEmpty(t, rootDir)
}
