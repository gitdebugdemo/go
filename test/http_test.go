package test

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

//模拟post - 提交json整段数据
func TestPOSTjson(t *testing.T) {
	bodys := "{\"xx\":\"1不知道\"}"
	resp, _ := http.Post("http://127.0.0.1:8080/json", "application/json;charset=utf-8", bytes.NewBuffer([]byte(bodys)))
	//程序在使用完回复后必须关闭回复的主体。
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, string(body), bodys)
}
