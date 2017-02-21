// GNU GPL v3 License

// Copyright (c) 2017 github.com:go-trellis

package main

import (
	"fmt"

	"github.com/go-trellis/config"
)

func main() {
	c, e := config.NewConfig("example.yml")
	if e != nil {
		fmt.Println(e)
		return
	}
	// get key value
	printT("get a Easy!\t", c.GetString("a"))
	printT("get a.b.c def:example", c.GetString("a.b.c", "example"))
	printT("get b.c.e\t", c.GetString("b.c.e"))
	printT("get h is 1.01\t", c.GetFloat("h"))
	printT("get b.c.f def:100,but 2", c.GetInt("b.c.f", 100))
	printT("get b.c.e not exist", c.GetInt("b.c.e"))
	printT("get b.c.g ON,return T", c.GetBoolean("b.c.g"))
	printT("get b.c.x def:true", c.GetBoolean("b.c.x", true))
	printT("get b config\t", c.GetConfig("b"))
	printT("get b.d list 1->2", c.GetList("b.d"))

	// set key value
	printT("set a.b.c Correct", c.SetKeyValue("a.b.c", "Correct"))
	printT("set b.c.e Correct", c.SetKeyValue("b.c.e", "Correct"))
	printT("set b.c.d d\t", c.SetKeyValue("b.c.d", "d"))
	printT("set b.c.g false ", c.SetKeyValue("b.c.g", false))
	printT("set b.d list 1->4", c.SetKeyValue("b.d", []int{1, 2, 3, 4}))

	// get key value
	printT("get a string\t", c.GetString("a", "example"))
	printT("get a interface\t", c.GetInterface("a", "example"))
	printT("get a.b.c set Correct", c.GetString("a.b.c", "example"))
	printT("get b.c.e set Correct", c.GetString("b.c.e", "example"))
	printT("get b.c.g set false", c.GetBoolean("b.c.g"))
	printT("get b.c.d set d\t", c.GetString("b.c.d", "example"))

	// set key value
	printT("set a Difficult!", c.SetKeyValue("a", "Difficult!"))

	// get key value
	printT("get a.b.c def:example", c.GetString("a.b.c", "example"))
	printT("get a Difficult!", c.GetString("a", "example"))
	printT("get a list nil\t", c.GetList("a"))
	printT("get b.d string\t", c.GetString("b.d", "example"))
	printT("get b.d list 1->4", c.GetList("b.d"))

	bs, _ := c.Dump()
	printT("last dump", "\n"+string(bs))
}

func printT(key string, v interface{}) {
	fmt.Printf("%s\t= %v\n", key, v)
}
