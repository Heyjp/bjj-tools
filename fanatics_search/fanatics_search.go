package fanatics_search

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/anaskhan96/soup"
)

var site = "https://bjjfanatics.com/products/"

// Looks up an individual product on the fanatics site and extracts
// the timestamp data from the page and places it into a series of chapters files
func Search(product string, folder string) {

	if product == "" {
		log.Fatal("fanatics_search <product> <folder>")
	}

	if len(folder) < 1 {
		folder = "./chapters"
	}

	productRe := regexp.MustCompile(`[\w|-]*$`)
	matches := productRe.FindAllString(product, -1)

	if len(matches) == 0 {
		log.Fatal("Please enter a valid product: https://bjjfanatics.com/products/<product>")
	}

	resp, err := soup.Get(site + matches[0])

	if err != nil {
		log.Fatal(err)
	}

	doc := soup.HTMLParse(resp)

	section := doc.Find("section", "id", "contents-no-mobile")
	rows := section.FindAll("tbody")

	re, _ := regexp.Compile(`^[^\d]*([\d:\.]+)`)
	wordRe, _ := regexp.Compile(`[a-zA-Z]+`)

	for i, tbody := range rows {
		var c [][]string
		for _, child := range tbody.FindAll("tr") {
			var chapterStamp []string

			// Should return 2 elements one for the title and one for the
			//	 timestamp
			// "Title" "0:01:00"
			for _, tr := range child.FindAll("td") {
				textMatch := tr.Text()
				textMatch = strings.ReplaceAll(textMatch, "\n", " ")

				// Dealing with a timestamp
				timeMatch := re.FindStringSubmatch(textMatch)
				wordMatch := wordRe.MatchString(textMatch)

				// sanitize the timestamp changing period to colons and also
				// extracting the first value incase of 12:32:22 - 23:42:21
				// style timestamps
				if len(timeMatch) > 0 && wordMatch == false {
					textMatch = strings.ReplaceAll(timeMatch[0], ".", ":")
					//	textMatch = timeMatch[0]
				}
				chapterStamp = append(chapterStamp, textMatch)
			}
			c = append(c, chapterStamp)
		}

		fileString := fmt.Sprintf("%s/chapters-%d.txt", folder, i+1)
		file, err := os.Create(fileString)

		if err != nil {
			log.Fatal(err)
		}

		defer file.Close()

		var chapterString string
		for _, chapter := range c {
			chapterString += strings.Join(chapter, ",")
			chapterString += "\n"
		}
		_, errF := file.WriteString(chapterString)

		if errF != nil {
			log.Fatal(errF)
		}
	}

}
