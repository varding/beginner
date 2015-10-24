package main

import (
	"github.com/alecthomas/log4go"
	"route"
)

func main() {
	log4go.AddFilter("stdout", log4go.FINE, log4go.NewConsoleLogWriter())
	log4go.AddFilter("filelog", log4go.FINE, log4go.NewFileLogWriter("beginner.log", false))
	route.Run()
}
