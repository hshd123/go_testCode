package main

import (
	//"encoding/json"
	"fmt"
	//"github.com/pquerna/ffjson/ffjson"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/json-iterator/go"
	 _ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"os"
	"time"
)

type FlowerItem struct {
	ImageName   string `json:"imageName"`
	Title       string `json:"title"`
	SubTitle    string `json:"subTitle"`
	ContextText string `json:"contextText"`
}

type Flower struct {
	Header  string       `json:"header"`
	Flowers []FlowerItem `json:"flowers"`
}

func main() {
	stars := time.Now().Nanosecond()
	fmt.Println("stars -- ", stars)
	fd, err := os.Open("./Flowers.json")
	if err != nil {
		fmt.Println("error -- ", err.Error())
	}
	b1, err := ioutil.ReadAll(fd)
	if err != nil {
		fmt.Println(err.Error())
	}
	l1 := make([]Flower, 0)
	err = jsoniter.Unmarshal(b1, &l1)
	if err != nil {
		fmt.Println("err  ", err.Error())
	}
	end := time.Now().Nanosecond()
	fmt.Println("en -- ", end)
	fmt.Println("ffjson unmal执行时间 -- ", (end-stars)/1000000.0, "  毫秒")
	engine := xormInit()

	isExit , err := engine.IsTableExist(Flower{})
	if isExit == false {
		err := engine.CreateTables(new(Flower))
		if err != nil {
			fmt.Println("err - " , err.Error())
		}
	}
	ss := engine.NewSession()

	for i := 0; i < len(l1); i++ {
		f1 := l1[i]
		_ ,err = engine.Insert(&f1)
		if err != nil  {
			ss.Rollback()
			fmt.Println(" err - ",err.Error())
		}
	}

	ss.Commit()

}

/*xorm 初始化*/
func xormInit() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:123456@/test_db?charset=utf8")
	if err != nil {
		fmt.Println(err.Error())
	}
	e := engine.Ping()
	if e != nil {
		fmt.Println(e.Error())
	}
	engine.ShowExecTime(true)
	engine.ShowSQL(true)
	engine.SetMaxIdleConns(10)
	engine.SetMaxOpenConns(10)
	engine.SetMapper(core.SameMapper{})

	fmt.Println(engine.DBMetas())
	engine.Charset("utf8mb4")

	return engine
}
