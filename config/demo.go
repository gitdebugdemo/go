package config

import (
	"database/sql"
    "log"
    "strconv"
    "sync"
    "strings"
    "xltxlm/go/env"
	_ "github.com/go-sql-driver/mysql"
)

//Demo 数据库的配置结构
type Demo struct {
    Tns string
    User string
    Password string
    DB string
    Port int
}

func stringsContains(array []string, val string) (index int) {
    index = -1
    for i := 0; i < len(array); i++ {
        if array[i] == val {
            index = i
            return
        }
    }
    return
}



//New 返回数据库连接配置
func  NewDemo() Demo {
    dockername := env.NewEnvinfo().Dockername
    log.Println("环境:"+dockername)
        dockernames := []string{"dev_go_demo"}
        //测试环境配置(后台:,扩展:["dev_go_demo"])
        if(stringsContains(dockernames,dockername)!=-1 && strings.Index(dockername,"dev_") == 0){
        return Demo{
            Tns : "192.168.31.101",
            User : "demo",
            Password : "demo",
            DB : "demo",
            Port : 3306,
        };
    }

    //默认连接测试环境的数据库
    return Demo{
        Tns : "127.0.0.1",
        User : "demo",
        Password : "demo",
        DB : "demo",
        Port : 3306,
    };
}

var db *sql.DB
var once sync.Once

//Connect 链接数据库
func  NewDemoConnect() (*sql.DB,error) {
    once.Do(func() {
        var err error
        newlink := NewDemo()
        db, err = sql.Open("mysql", newlink.User+":"+newlink.Password+"@tcp("+newlink.Tns+":"+strconv.Itoa(newlink.Port)+")/"+newlink.DB+"?charset=utf8")
        if err != nil {
            log.Println("链接数据库错误")
            log.Printf("%+v",newlink)
            log.Println(err)
        }
        // 设置最大连接数
        db.SetMaxOpenConns(10)
        // 设置最大空闲连接数
        db.SetMaxIdleConns(2)
    })
    return db,db.Ping()
}

//Connect 链接数据库
func (newlink Demo) DemoConnect() (*sql.DB,error) {
    var err error
    db, err = sql.Open("mysql", newlink.User+":"+newlink.Password+"@tcp("+newlink.Tns+":"+strconv.Itoa(newlink.Port)+")/"+newlink.DB+"?charset=utf8")
    if err != nil {
        log.Println("链接数据库错误")
        log.Printf("%+v",newlink)
        log.Println(err)
    }
    return db,db.Ping()
}