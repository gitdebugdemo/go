package env

import "sync"

// Envinfo 项目的环境配置
type Envinfo struct {
	Dockername string
}

var m *Envinfo
var once sync.Once

//Getenv 获取当前编译的环境配置
func Getenv() *Envinfo {
	once.Do(func() {
		m = &Envinfo{Dockername: "demo"}
	})
	return m
}
