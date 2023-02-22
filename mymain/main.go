package main

import (
	"fmt"
	"plugin"

	_ "golang.org/x/sys/execabs"
)

func main() {
	_, err := plugin.Open("myplugin.so")
	if err != nil {
		panic(err)
	}

	fmt.Println("success")
}
