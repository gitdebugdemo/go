package test

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

//存储结构 - 只有大写开头的属性才会被序列号成功,小写的对外不可见
type jsont struct {
	Id int
	id int
	//重定向解析的字段
	Name string `json:"namexx"`
	name string
	//重定向解析的字段
	Ok bool `json:"myok"`
	ok bool
}

//反序列化成对象
func TestJson(t *testing.T) {

	//生成一个新的对象
	jsono_obj := jsont{Id: 1234, id: 222, Name: "okname中文", name: "xname", ok: false, Ok: true}
	//开始序列化对象 - 输出的时候,字段也被格式化成指定名字,不再是属性名称了!!!!!!
	b, _ := json.Marshal(jsono_obj)
	outs := `{"Id":1234,"namexx":"okname中文","myok":true}`
	os.Stdout.Write(b)
	os.Stdout.WriteString("\n")
	assert.Equal(t, b, []byte(outs))

	s := `{"Id":1234,"namexx":"okname中文","Ok":true,"myok":false}`

	//这里开始本实验,反序列化回去
	var abc_jsont jsont
	fmt.Printf("%+v\n", abc_jsont)
	json.Unmarshal([]byte(s), &abc_jsont)
	assert.Equal(t, 1234, abc_jsont.Id)
	assert.Equal(t, 0, abc_jsont.id)
	assert.Equal(t, "", abc_jsont.name)
	assert.Equal(t, "okname中文", abc_jsont.Name)
	assert.Equal(t, false, abc_jsont.Ok)
	assert.Equal(t, false, abc_jsont.ok)
}
