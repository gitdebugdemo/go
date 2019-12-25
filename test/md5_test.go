package test

import (
	"crypto/md5"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMd5(t *testing.T) {
	sum := md5.Sum([]byte("axaxaxaxsa"))
	//转换成字符串
	md5str := fmt.Sprintf("%x", sum)
	assert.Equal(t, md5str, "010ac6044ba8ab01cfb87fbbb3e38399")
}
