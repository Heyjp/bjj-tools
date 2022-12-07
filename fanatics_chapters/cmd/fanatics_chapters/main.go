package main

import (
	"fmt"
	"log"
	"os"

	c "github.com/heyjp/bjj-tools/fanatics_chapters"
	p "github.com/heyjp/bjj-tools/fanatics_crawler"
	f "github.com/heyjp/bjj-tools/fantatics_search"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: main <crawl|chapters|search|merge>")
	}
	switch os.Args[1] {
	case "crawl":
		if len(os.Args) < 3 {
			log.Fatal("Usage: main crawl <all|new>")
		}
		siteKey := os.Args[2]
		p.Crawl(siteKey)
	case "chapters":
		c.LoopThroughProducts()
	case "search":
		if len(os.Args) < 3 {
			log.Fatal("Usage: main search <fanatics-product-link>")
		}
		product := os.Args[3]
		folder := os.Getwd()
		f.Search(product, folder)
	case "merge":
		fmt.Println("Enter in a folder location containing videos")
	default:
		fmt.Println("main <search|crawl|chapters|merge>")
	}

	// p.Craw l("all")
	// c.LoopThroughProducts()
}
