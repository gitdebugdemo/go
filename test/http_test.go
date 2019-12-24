package test

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

//模拟get
func TestGeta(t *testing.T) {
	resp, _ := http.Get("http://127.0.0.1:8080/h1?a=ok")
	body, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, string(body), "ok")
}
//模拟post
func TestPOSTa(t *testing.T) {
	resp, _ := http.PostForm("http://127.0.0.1:8080/h1", url.Values{"a": {"ok2"}})
	body, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, string(body), "ok2")
}
