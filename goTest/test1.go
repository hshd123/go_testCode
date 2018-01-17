package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sync"
)

func main() {
	//test1()
	TestDoit()
}

func test1() {
	jsonStr := []byte(`{"age":1}`)
	var value map[string]interface{}
	json.Unmarshal(jsonStr, &value)
	age := value["age"]
	fmt.Println(reflect.TypeOf(age))
	fmt.Println(value)
}

type UserAges struct {
	ages map[string]int
	sync.Mutex
}

func (u *UserAges) Add(name string, age int) {
	u.Lock()
	defer u.Unlock()
	u.ages[name] = age
}

func (u *UserAges) Get(name string) int {
	if age, ok := u.ages[name]; ok {
		return age
	}
	return -1
}

func TestDoit() {
	doit := func(arg int) interface{} {
		var result *struct{} = nil
		if arg > 0 {
			result = &struct{}{}
		}
		return result
	}(3)
	fmt.Println(doit)
}
