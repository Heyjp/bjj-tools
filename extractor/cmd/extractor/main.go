package main

import (
	"os"

	e "github.com/heyjp/bjj-tools/extractor"
)

func main() {
	var v = os.Args[1]
	// var m = os.Args[2]
	e.Extractor(v)
}
