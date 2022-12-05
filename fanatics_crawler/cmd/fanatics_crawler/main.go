package main

import (
	"os"

	c "github.com/heyjp/bjj-tools/fanatics_crawler"
)

func main() {
	var siteKey string
	if len(os.Args) < 2 {
		siteKey = "all"
	} else {
		siteKey = os.Args[1]
	}

	c.Crawl(siteKey)
}
