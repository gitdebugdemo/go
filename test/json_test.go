package test

import (
	"encoding/json"
	"fmt"
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

//反序列化成对象
func TestJson(t *testing.T) {
	//生成一个新的对象
	jsono_obj := jsont{Id: 1234, id: 222, Name: "okname中文", name: "xname", ok: false, Ok: true}
	//开始序列化对象
	b, _ := json.Marshal(jsono_obj)
	s := `{"Id":1234,"Name":"okname中文","Ok":true}`
	assert.Equal(t, b, []byte(s))

	//这里开始本实验,反序列化回去
	var abc_jsont jsont
	fmt.Printf("%+v\n", abc_jsont)
	json.Unmarshal([]byte(s), &abc_jsont)
	assert.Equal(t, 1234, abc_jsont.Id)
	assert.Equal(t, 0, abc_jsont.id)
	assert.Equal(t, "", abc_jsont.name)
	assert.Equal(t, "okname中文", abc_jsont.Name)
	assert.Equal(t, true, abc_jsont.Ok)
	assert.Equal(t, false, abc_jsont.ok)
}

//反序列化成对象数组
func TestJsonArray(t *testing.T) {
	var abc_jsont []jsont
	s := `[{"Id":111,"Name":"okname中文111","Ok":false},{"Id":1234,"Name":"okname中文","Ok":true}]`
	json.Unmarshal([]byte(s), &abc_jsont)
	fmt.Printf("%+v\n", abc_jsont)
	//目标预期对象
	news := []jsont{
		{Id: 111, id: 0, Name: "okname中文111", name: "", Ok: false, ok: false},
		{Id: 1234, id: 0, Name: "okname中文", name: "", Ok: true, ok: false},
	}
	assert.Equal(t, news, abc_jsont)
}
