package fanatics_chapters

import (
	"bufio"
	"log"
	"os"
	"strings"

	c "github.com/heyjp/bjj-tools/chapters"
	d "github.com/heyjp/bjj-tools/dircheck"
	p "github.com/heyjp/bjj-tools/fanatics_crawler"
	s "github.com/heyjp/bjj-tools/fanatics_search"
)

func LoopThroughProducts() {
	p.Crawl()

	file, err := os.Open("fanatics-products.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		log.Println(line)
		s := strings.Split(line, " ")
		CreateChapters(s[0], s[1], false)
	}
}

func CreateChapters(product, location string, yt bool) {
	// arg - product name on bjj fanatics
	// creates the chapters folder in the current working directory
	// Searches bjj fantatics and returns chapters files
	// converts chapters into useable chapters
	d.CheckOrCreateDirectory(location)
	s.Search(product, location)

	files, err := os.ReadDir(location)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		t := c.PrepareTimestamps(location + "/" + file.Name())
		c.CreateChaptersFile(t, location+"/"+file.Name(), yt)
	}

}