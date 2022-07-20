package main

import (
	"flag"
	"fmt"
)

var (
	from, to      string
	limit, offset int64
)

func init() {
	flag.StringVar(&from, "from", "testdata/input.txt", "file to read from")
	flag.StringVar(&to, "to", ".", "file to write to")
	flag.Int64Var(&limit, "limit", 0, "limit of bytes to copy")
	flag.Int64Var(&offset, "offset", 0, "offset in input file")
}

func main() {
	flag.Parse()
	err := Copy(*&from, *&to, *&limit, *&offset)
	fmt.Printf("Got err %v in main\n", err)
}
