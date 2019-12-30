package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnect(t *testing.T) {
	var q测试下中文变命名 string = "china命名"
	assert.Equal(t, q测试下中文变命名, "china命名")
}
