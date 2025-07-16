// cmd/hello/main.go
package main

import (
	"flag"
	"os"

	"github.com/yuyan42517/hello/pkg/hello"
)

var name string

func init() {
	flag.StringVar(&name, "name", "world", "name to say hello")
}

func main() {
	flag.Parse()
	msg := hello.Hello(name)
	_, err := os.Stdout.WriteString(msg)
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
}
