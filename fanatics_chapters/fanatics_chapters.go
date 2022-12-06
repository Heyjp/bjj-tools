package fanatics_chapters

import (
	"bufio"
	"log"
	"os"
	"strings"

	c "github.com/heyjp/bjj-tools/chapters"
	d "github.com/heyjp/bjj-tools/dircheck"
	s "github.com/heyjp/bjj-tools/fanatics_search"
)

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
		CreateChapters(s[0], s[1], false)
	}
}

func CreateChapters(product, location string, yt bool) {
	// arg - product name on bjj fanatics
	// creates the chapters folder in the current working directory
	// Searches bjj fantatics and returns chapters files
	// converts chapters into useable chapters

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
		if len(errorStamps) > 0 {
			c.CreateErrorsFile(e, location+"/errors/"+"errors.txt")
		}

	}

}
