package main

import (
	"fmt"

	do "gopkg.in/godo.v2"
)

func tasks(p *do.Project) {
	do.Env = `GOPATH=.vendor::$GOPATH`

	p.Task("default", do.S{"hello"}, nil)

	p.Task("hello", nil, func(c *do.Context) {
		name := c.Args.AsString("name", "n")
		if name == "" {
			c.Bash("echo Hello $USER!")
		} else {
			fmt.Println("Hello", name)
		}
	})
}

func main() {
	do.Godo(tasks)
}
