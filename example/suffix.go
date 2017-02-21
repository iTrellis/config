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

	fmt.Println(c)

	printT("a", c.GetString("a"))
	printT("a.b.c", c.GetString("a.b.c", "example"))
	printT("b.c.e", c.GetString("b.c.e"))
	printT("h", c.GetFloat("h"))
	printT("b.c.f", c.GetInt("b.c.f", 100))
	printT("b.c.g", c.GetBoolean("b.c.g"))
	printT("b.c.x", c.GetBoolean("b.c.x", true))
	printT("b", c.GetConfig("b"))
}

func printT(key string, v interface{}) {
	fmt.Printf("%s \t = %v\n", key, v)
}
