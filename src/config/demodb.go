package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
	"sync"
)

//Demodb 数据库的配置结构
type Demodb struct {
	Tns      string
	User     string
	Password string
	DB       string
	Port     int
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
func (t Demodb) New() Demodb {
	return Demodb{
		Tns:      "127.0.0.1",
		User:     "demo",
		Password: "demo",
		DB:       "demo",
		Port:     3306,
	}
}

var db *sql.DB
var once sync.Once

//Connect 链接数据库
func (t Demodb) Connect() (*sql.DB, error) {
	once.Do(func() {
		var err error
		newlink := t.New()
		db, err = sql.Open("mysql", newlink.User+":"+newlink.Password+"@tcp("+newlink.Tns+":"+strconv.Itoa(newlink.Port)+")/"+newlink.DB)
		if err != nil {
			log.Println("链接数据库错误")
			log.Printf("%+v", newlink)
			log.Println(err)
		}
	})
	return db, db.Ping()
}
