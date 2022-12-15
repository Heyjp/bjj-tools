package main

import (
	"os"

	f "github.com/heyjp/bjj-tools/search"
)

func main() {
	product := os.Args[1]
	var folder string = ""
	if len(os.Args) > 2 {
		folder = os.Args[2]
	}
	f.Search(product, folder)
}
