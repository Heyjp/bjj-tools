package fanatics_chapters

import (
	"bufio"
	"log"
	"os"
	"strings"

	d "github.com/heyjp/bjj-tools/dircheck"
	s "github.com/heyjp/bjj-tools/search"
	c "github.com/heyjp/bjj-tools/timestamps"
)

// Reads a fanatics-products.txt file and extracting the product
// name and folder location and transforms the scraped timestamps
// into a useable format to be combined with a metadata file
func LoopThroughProducts() {
	file, err := os.Open("fanatics-products.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		s := strings.Split(line, " ")
		PrepareChapterFiles(s[0], s[1], false)
	}
}

// Loops through a folder of chapter-[n].txt files and puts them in
// a format ready to be combined with a metadata file
func PrepareChapterFiles(product, location string, yt bool) {
	_, errF := os.Stat(location)
	if errF == nil {
		return
	}

	d.CheckOrCreateDirectory(location)
	s.Search(product, location)

	files, err := os.ReadDir(location)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		t, e := c.PrepareTimestamps(location + "/" + file.Name())
		c.CreateChaptersFile(t, location+"/"+file.Name(), yt)

		if len(e) > 0 {
			c.CreateErrorsFile(e, location)
		}

	}

}
