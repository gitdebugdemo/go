package test

import (
	"github.com/gitdebugdemo/go/src/config"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestConnect(t *testing.T) {
	_, err := new(config.Demodb).Connect()
	assert.Nil(t, err)
}

func TestFetch(t *testing.T) {
	db, _ := new(config.Demodb).Connect()
	//只取出一行,所以不需要句柄
	rows := db.QueryRow("select 'abc' as t ")
	var tt1 string
	if err := rows.Scan(&tt1); err != nil {
		log.Fatal(err)
	}
	assert.Equal(t, "abc", tt1)
}
