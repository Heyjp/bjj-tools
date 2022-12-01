package main

import (
	"os"

	c "github.com/heyjp/bjj-tools/fanatics_chapters"
	p "github.com/heyjp/bjj-tools/fanatics_crawler"
)

func main() {
	if os.Args[1] == "crawl" {
		p.Crawl()
		return
	}

	if os.Args[1] == "chapters" {
		c.LoopThroughProducts()
		return
	}

	p.Crawl()
	c.LoopThroughProducts()
}
