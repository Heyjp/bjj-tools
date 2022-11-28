package metadata

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var timebase string = "1/1000"

type Chapter struct {
	title string
	start int
}

// Takes two strings for a metadata file location and a chapters
// file location
func MergeChaptersWithMetadata(m string, c string) {
	if len(os.Args) < 2 {
		fmt.Println("You must enter in a file location")
		return
	}

	if len(os.Args) != 3 {
		log.Fatal("Please enter a metadata file and a chapters file")
	}

	TestMetaData(m)
	ClearVideoMetaData(m)
	chapters := CreateChapters(c)
	WriteChaptersToMetaData(m, chapters)
}

// Checks whether the file being passed in is a metadata file
func TestMetaData(location string) {
	file, err := os.Open(location)

	if err != nil {
		log.Panicf("Error opening test file: %v", err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	input, err := reader.ReadString('\n')

	fileType, err := regexp.Compile(";FFMETADATA1")

	if err != nil {
		log.Fatal("Bad File")
	}

	if fileType.MatchString(input) == false {
		log.Fatal("File is not a metadata file")
	}
}

// Searches through a file to find whether it already has chapters or
// not if chapters are found they are removed
func ClearVideoMetaData(location string) {
	file, err := os.Open(location)

	if err != nil {
		log.Panicf("Error opening file: %v", err)
	}
	defer file.Close()

	re, err := regexp.Compile("CHAPTER")

	if err != nil {
		log.Fatal("Bad Regex")
	}

	bytesRead := 0
	scanner := bufio.NewScanner(file)
out:
	for scanner.Scan() {

		b := scanner.Text()
		results := re.FindAllString(b, -1)
		for _, result := range results {
			if len(result) > 0 {
				break out
			}
		}

		// Add 1 for end of line marker unix
		bytesRead += len(b) + 1
	}

	fi, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	size := fi.Size()
	if int64(bytesRead) < size {
		size = int64(bytesRead)
	}
	truncErr := os.Truncate(location, size)
	if truncErr != nil {
		log.Fatal(err)
	}
}

// Returns a slice of Chapters objects for processing into the
// metadata file
func CreateChapters(location string) []Chapter {
	var chapters []Chapter
	file, err := os.Open(location)
	if err != nil {
		log.Fatal("failed to create %q", err)
	}

	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		re := regexp.MustCompile(`(\d{2}):(\d{2}):(\d{2}) (.*)`)
		res := re.FindStringSubmatch(fileScanner.Text())

		hour, _ := strconv.Atoi(res[1])
		min, _ := strconv.Atoi(res[2])
		sec, _ := strconv.Atoi(res[3])
		title := res[4]

		minutes := (hour * 60) + min
		seconds := sec + (minutes * 60)
		timestamp := (seconds * 1000)

		chap := Chapter{title, timestamp}
		chapters = append(chapters, chap)
	}
	return chapters
}

// Combines the chapters object with the metadata file
func WriteChaptersToMetaData(metaDataFile string, chapters []Chapter) {

	file, err := os.OpenFile(metaDataFile, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var c string

	for i, chapter := range chapters {
		c += "[CHAPTER]\n"
		c += fmt.Sprintf("TIMEBASE=%q\n", timebase)
		c += fmt.Sprintf("START=%d\n", chapter.start)
		end := chapter.start + 1
		if i < len(chapters)-1 {
			end = chapters[i+1].start - 1
		}
		c += fmt.Sprintf("END=%d\n", end)
		c += fmt.Sprintf("title=%s\n", chapter.title)
		c += "\n"
	}
	if _, err := file.WriteString(c); err != nil {
		log.Println(err)
	}

}
