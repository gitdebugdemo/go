package test

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHttp(t *testing.T) {
	resp, _ := http.Get("http://127.0.0.1:8080")
	body, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, "OK\n", string(body))
}
