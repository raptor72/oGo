package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

const str = "Hello, OTUS!"

func main() {
	rev := stringutil.Reverse(str)
	fmt.Println(rev)
}
