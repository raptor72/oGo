package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	env, err := ReadDir("testdata/env")
	if err != nil {
		fmt.Println(err)
	}

	if code := RunCmd(args[2:], env); code != 0 {
		fmt.Println("got err exit code: ", code)
	}
}
