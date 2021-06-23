package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(runtime.NumCPU(), runtime.GOOS, runtime.GOARCH)
}

//Test comment
