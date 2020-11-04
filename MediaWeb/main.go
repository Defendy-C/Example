package main

import (
	"MediaWeb/controller"
	"os"
)

var static, _ = os.Getwd()

func main() {
	route := controller.NewRouter()

	s := controller.NewServer(route, "192.168.14.172:8000")
	panic(s.ListenAndServ())
}
