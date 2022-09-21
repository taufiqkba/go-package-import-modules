package main

import (
	"example.go/go-package-import-modules/helper"
	"example.go/go-package-import-modules/other"
	"fmt"
)

func main() {
	fmt.Println(helper.SayHello("Budi"))
	fmt.Println(other.SayHelloOther("taufiq"))
}
