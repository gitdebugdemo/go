package test

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHttp404(t *testing.T) {
	//没有对应的处理函数
	resp, _ := http.Get("http://127.0.0.1:8080/")
	assert.Equal(t, resp.StatusCode, 404)
}

func TestH1(t *testing.T) {
	resp, _ := http.Get("http://127.0.0.1:8080/h1")
	body, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, string(body), "h1")
}

func TestH2(t *testing.T) {
	resp, _ := http.Get("http://127.0.0.1:8080/h2")
	body, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, string(body), "h2")
}
