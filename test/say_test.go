package test

import (
	"github.com/gitdebugdemo/go/src/funcs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMd5(t *testing.T) {
	assert.Equal(t,"hello v1",funcs.Say())
}
