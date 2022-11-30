package fanatics_crawler

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/anaskhan96/soup"
)

func Crawl() {
	if _, err := os.Stat("./fanatics-products.txt"); err == nil {
		return
	}

	var site = "https://bjjfanatics.com/collections/instructional-videos"
	var q = "?page="

	respI, errI := soup.Get(site)

	if errI != nil {
		os.Exit(1)
	}

	docI := soup.HTMLParse(respI)

	div := docI.Find("div", "class", "pagination")
	spans := div.FindAll("a")

	maxPages, err := strconv.Atoi(spans[len(spans)-2].Text())

	if err != nil {
		log.Println(err)
	}

	re := regexp.MustCompile(`[\w|-]*$`)
	var products []string
	// Start the query at page 2 since we have the first page from the
	// inital site query
	for i := 1; i <= maxPages; i++ {
		url := fmt.Sprintf("%s%s%d", site, q, i)
		log.Println(url)

		resp, err := soup.Get(url)

		if err != nil {
			log.Println(err)
		}

		doc := soup.HTMLParse(resp)
		anchors := doc.FindAll("a", "class", "product-card")

		for _, a := range anchors {
			s := re.FindString(a.Attrs()["href"])
			products = append(products, s)
		}

	}

	f, err := os.OpenFile("fanatics-products.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var c string

	for _, item := range products {
		if strings.Contains(item, "bundle") {
			continue
		}
		c += fmt.Sprintf("%s chapters/%s\n", item, item)
	}

	if _, errS := f.WriteString(c); errS != nil {
		log.Println(err)
	}

}
