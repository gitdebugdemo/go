package test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

//存储结构 - 只有大写开头的属性才会被序列号成功,小写的对外不可见
type jsont struct {
	Id   int
	id   int
	Name string
	name string
	Ok   bool
	ok   bool
}

func TestJson(t *testing.T) {
	//生成一个新的对象
	jsono_obj := jsont{Id: 1234, id: 222, Name: "okname中文", name: "xname", ok: false, Ok: false}
	//开始序列化对象
	b, _ := json.Marshal(jsono_obj)
	assert.Equal(t, b, []byte(`{"Id":1234,"Name":"okname中文","Ok":false}`))
}
