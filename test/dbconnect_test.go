package test

import (
	"database/sql"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"xltxlm/go/config"
)

var (
	db  *sql.DB
	err error
)

//ORM存储的数据结构
type userinfo struct {
	id       int
	name     string
	add_time string
}

func TestQueryRow(t *testing.T) {
	db, err = config.NewDemoConnect()
	if err != nil {
		fmt.Print("数据库链接错误")
		fmt.Printf("%+v\n		", err)
	}
	//初始化值
	var myuserinfo userinfo = userinfo{}
	//一个一个字段敲-蛋疼
	row := db.QueryRow("select id,name,add_time from userinfo where id=? and name=?", 1, "新来的")
	//一个一个属性再敲一次,蛋疼得晕过去了
	err := row.Scan(&myuserinfo.id, &myuserinfo.name, &myuserinfo.add_time)
	if err != nil {
		fmt.Print("Scan错误")
		fmt.Printf("%+v\n", err)
	}
	//校验
	assert.Equal(t, myuserinfo.id, 1)
	assert.Equal(t, myuserinfo.name, "新来的")
	assert.Equal(t, myuserinfo.add_time, "2019-12-30 15:24:27")
	fmt.Printf("%+v", myuserinfo)
}
