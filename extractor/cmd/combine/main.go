package main

import (
	"os"

	e "github.com/heyjp/bjj-tools/extractor"
)

func main() {
	var v, m, o string
	v = os.Args[1]
	m = os.Args[2]
	o = os.Args[3]

	e.Combine(v, m, o)
}
