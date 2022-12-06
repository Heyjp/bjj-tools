package main

import (
	"os"

	c "github.com/heyjp/bjj-tools/fanatics_chapters"
	p "github.com/heyjp/bjj-tools/fanatics_crawler"
)

func main() {
	var siteKey string
	if len(os.Args) < 2 {
		siteKey = "all"
	} else {
		siteKey = os.Args[2]
	}

	if os.Args[1] == "crawl" {
		p.Crawl(siteKey)
		return
	}

	if os.Args[1] == "chapters" {
		c.LoopThroughProducts()
		return
	}

	p.Crawl(siteKey)
	c.LoopThroughProducts()
}
