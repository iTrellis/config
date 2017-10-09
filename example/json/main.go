// GNU GPL v3 License

// Copyright (c) 2017 github.com:go-trellis

package main

import (
	"fmt"

	"github.com/go-trellis/config"
)

func main() {
	c, e := config.NewConfig("../example.json")
	if e != nil {
		fmt.Println(e)
		return
	}

	fmt.Println(c)
	// get key value
	printT("get a Easy!\t", c.GetString("a"))
	printT("copy b.c.cn.a test!", c.GetInterface("b.c.cn.a"))
	c.SetKeyValue("b.c.cn.a", "value")
	printT("copy b.c.cn.a value!", c.GetInterface("b.c.cn.a"))
	printT("copy n.a value!\t", c.GetInterface("n.a"))
	printT("copy b.d to b.c.cbd!", c.GetIntList("b.c.cbd"))
	printT("get a.b.c def:example", c.GetString("a.b.c", "example"))
	printT("get b.c.e\t", c.GetString("b.c.e"))
	printT("get h is 1.01\t", c.GetFloat("h"))
	printT("get b.c.f def:100,but 2", c.GetInt("b.c.f", 100))
	printT("get b.c.e not exist", c.GetInt("b.c.e"))
	printT("get b.c.g ON,return T", c.GetBoolean("b.c.g"))
	printT("get b.c.x def:true", c.GetBoolean("b.c.x", true))
	printT("get b config\t", c.GetConfig("b"))
	printT("get b.d list 3->4", c.GetList("b.d"))
	printT("get b.c.t time:1day", c.GetTimeDuration("b.c.t"))

	// set key value
	printT("set a.b.c Correct", c.SetKeyValue("a.b.c", "Correct"))
	printT("set b.c.e Correct", c.SetKeyValue("b.c.e", "Correct"))
	printT("set b.c.d d\t", c.SetKeyValue("b.c.d", "d"))
	printT("set b.c.g false ", c.SetKeyValue("b.c.g", false))
	printT("set b.d list 1->4", c.SetKeyValue("b.d", []int{1, 2, 3, 4}))

	// get key value
	printT("get a def:example", c.GetString("a", "example"))
	printT("get a interface\t", c.GetInterface("a", "example"))
	printT("get a.b.c set Correct", c.GetString("a.b.c", "example"))
	printT("get b.c.e set Correct", c.GetString("b.c.e", "example"))
	printT("get b.c.g set false", c.GetBoolean("b.c.g", true))
	printT("get b.c.d set d\t", c.GetString("b.c.d", "example"))

	// set key value
	printT("set a Difficult!", c.SetKeyValue("a", "Difficult!"))
	printT("set h.a list boolean", c.SetKeyValue("h.a", []bool{false, true, false}))
	printT("set h.f list float", c.SetKeyValue("h.f", []float64{1.2, 2.3, 3.4}))
	printT("set h.b byte size 10T", c.SetKeyValue("h.b", "10T"))

	// get key value
	printT("get a.b.c def:example", c.GetString("a.b.c", "example"))
	printT("get a Difficult!", c.GetString("a", "example"))
	printT("get a list nil\t", c.GetList("a"))
	printT("get h.a list boolean", c.GetBooleanList("h.a"))
	printT("get h.f list float", c.GetFloatList("h.f"))
	printT("get h float not exist", c.GetFloat("h"))
	printT("set h.b byte size 10t", c.GetByteSize("h.b"))
	printT("get b.d def:example", c.GetString("b.d", "example"))
	printT("get b.d []object 1->4", c.GetList("b.d"))
	printT("get b.d []string nil", c.GetStringList("b.d"))
	printT("get b.d []int 1->4", c.GetIntList("b.d"))

	printT("set b.d [\"1\",\"2\",\"3\"]", c.SetKeyValue("b.d", []string{"1", "2", "3"}))
	printT("get b.d []string 1->3", c.GetStringList("b.d"))

	bs, _ := c.Dump()
	printT("last dump", "\n"+string(bs))
}

func printT(key string, v interface{}) {
	fmt.Printf("%s\t= %v\n", key, v)
}
