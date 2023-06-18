package crawler

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/anaskhan96/soup"
)

var sites = map[string]string{
	"all": "https://bjjfanatics.com/collections/instructional-videos",
	"new": "https://bjjfanatics.com/collections/new-releases",
}

func Crawl(siteKey string) {
	f, err := os.OpenFile("fanatics-products.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	fileInfo, err := f.Stat()

	if err != nil {
		log.Println(err)
		return
	}

	var site string
	if value, containsKey := sites[siteKey]; containsKey {
		site = value
	} else {
		site = sites["all"]
	}

	var q = "?page="

	r, errI := soup.Get(site)

	if errI != nil {
		log.Fatal(errI)
	}

	doc := soup.HTMLParse(r)
	div := doc.Find("div", "class", "pagination")
	spans := div.FindAll("a")

	// Target the last page element in the pagination object
	maxPages, err := strconv.Atoi(spans[len(spans)-2].Text())

	if site == sites["new"] {
		maxPages = 10
	}

	if err != nil {
		log.Println(err)
		return
	}

	re := regexp.MustCompile(`[\w|-]*$`)

	// Queries each page, parses it and extracts timestamp info
	// Then appends that data
	for i := 1; i <= maxPages; i++ {
		var products []string

		url := fmt.Sprintf("%s%s%d", site, q, i)
		log.Println(url)

		resp, err := soup.Get(url)

		if err != nil {
			log.Println(err)
			return
		}

		doc := soup.HTMLParse(resp)
		anchors := doc.FindAll("a", "class", "product-card")

		for _, a := range anchors {
			s := re.FindString(a.Attrs()["href"])

			if strings.Contains(s, "bundle") {
				continue
			}

			if fileInfo.Size() > 0 && productExists(f, s) == true {
				continue
			}

			products = append(products, s)
		}

		var c string

		for _, item := range products {
			c += fmt.Sprintf("%s chapters/%s\n", item, item)
		}

		if _, errS := f.WriteString(c); errS != nil {
			log.Println(err)
			return
		}
	}

}

// Reads f file line by line to see if there is a matching product
// already
func productExists(f *os.File, p string) bool {

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		b := scanner.Text()
		res := strings.Contains(b, p)
		if res == true {
			f.Seek(0, io.SeekStart)
			return true
		}

	}

	f.Seek(0, io.SeekStart)
	return false
}
