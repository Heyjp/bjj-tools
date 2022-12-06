package chapters

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Timestamp struct {
	title   string
	hours   string
	minutes string
	seconds string
}

type Errorstamp struct {
	docLine string
	row     int
	chapter int
}

// Opens a text file at the location passed as an argument
// the file is then scanned for a string resembling a time / timestamp
func PrepareTimestamps(location string) ([]Timestamp, []Errorstamp) {
	file, err := os.Open(location)
	if err != nil {
		log.Fatal("failed to create %q", err)
	}

	defer file.Close()

	chapterRe := regexp.MustCompile("[0-9]+")
	cString := chapterRe.FindAllString(location, -1)
	chapter, _ := strconv.Atoi(cString[0])

	re := regexp.MustCompile(`(.*?)\s{1,4}(\d{0,2}?):?(\d{0,2}?):?(\d{0,2})$`)
	re2 := regexp.MustCompile(`^(\d{1,2}):?(\d{0,2}?):?(\d{0,2})\s{1,4}(.*)`)

	var stamps []Timestamp
	var errorStamps []Errorstamp

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	var row = 0

	for fileScanner.Scan() {
		text := fileScanner.Text()
		res := re.FindStringSubmatch(text)
		if len(res) == 0 {
			// If the previous regex does not match use fallback
			res = re2.FindStringSubmatch(text)
			if len(res) == 0 {
				log.Println("neither match")
				errorStamps = append(errorStamps, Errorstamp{text, row, chapter})
				row += 1
				continue
			}

			// Format is different so extract the time values for correction
			resSlice := res[1:4]

			// iterate through the values and swap the first and second values
			// if the second is missing. Regex subgrouping treats minutes as
			// hours so values need swapping
			for i, item := range resSlice {
				if item == "" {
					resSlice[i] = resSlice[i-1]
					resSlice[i-1] = ""
				}
			}
			newSlice := []string{res[4], res[4]}
			res = append(newSlice, resSlice...)
		}
		res = populateEmptyStrings(res)
		res = fixBadTimestamps(res)

		var timestamp = Timestamp{res[1], res[2], res[3], res[4]}

		// Checks whether the current timestamp is later than the
		// previous timestamp
		if len(stamps) > 1 {
			stampC, err := compareTimestamps(timestamp, stamps[len(stamps)-1])
			timestamp = stampC

			if err != nil {
				errorStamps = append(errorStamps, Errorstamp{err.Error(), row, chapter})
			}
		}
		stamps = append(stamps, timestamp)
		row += 1
	}

	return stamps, errorStamps
}

// removeEmptyStrings - Use to populate empty string values inside of an
// array
func populateEmptyStrings(s []string) []string {
	var r []string
	for _, str := range s {
		r = timesToStrings(r, str)
	}
	return r
}

// Checks whether the current timestamp is later than the previous
// timestamp
// Takes a slice of strings and transforms the would be time values
func timesToStrings(r []string, s string) []string {
	if s == "" {
		r = append(r, "00")
	}

	if len(s) == 1 {
		r = append(r, "0"+s)
	}
	if len(s) > 1 {
		r = append(r, s)
	}

	return r
}

// Recreates the chapter file with the timestamp first and the title
// second
func CreateChaptersFile(t []Timestamp, file string, forYt bool) {

	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var c string

	if forYt == true {
		c += "Chapters:\n"
	}

	for _, s := range t {
		var l string
		l += s.hours
		l += ":"
		l += s.minutes
		l += ":"
		l += s.seconds
		l += " " + s.title + "\n"
		c += l
	}

	if _, err := f.WriteString(c); err != nil {
		log.Println(err)
	}

}

// Adds some general rules that timestamps have to follow to make sure
// bad listings are attempted to be corrected

func fixBadTimestamps(t []string) []string {
	h, _ := strconv.Atoi(t[2])
	m, _ := strconv.Atoi(t[3])
	s, _ := strconv.Atoi(t[4])

	hours := strconv.Itoa(h)
	minutes := strconv.Itoa(m)
	seconds := strconv.Itoa(s)

	if h > 3 {
		hours = strconv.Itoa(0)
		minutes = strconv.Itoa(h)
		seconds = strconv.Itoa(m)
	}

	return populateEmptyStrings([]string{t[0], t[1], hours, minutes, seconds})
}

// Compare the current Timestamp object with the previous Timestamp
// object. If current timestamp is less than the previous do something
func compareTimestamps(c, p Timestamp) (Timestamp, error) {

	cT := prepareTimestamp(c)
	pT := prepareTimestamp(p)

	if cT > pT {
		return c, nil
	}

	// Increment the current timestamp by the interval and then convert it
	// back to a time friendly value
	h, m, s := incrementTimestamp(p)

	// Create an array of the new time values and convert them back to
	// strings to update the current timestamp to
	var v = []string{strconv.Itoa(h), strconv.Itoa(m), strconv.Itoa(s)}
	v = populateEmptyStrings(v)

	// Return the updated timestamp object
	c.hours = v[0]
	c.minutes = v[1]
	c.seconds = v[2]
	c.title += "- [Needs Edit]"
	return c, errors.New(fmt.Sprintf("%s:%s:%s %s", c.hours, c.minutes, c.seconds, c.title))
}

// Creates a single timestamp value out of a timestamp object {hours,
// minutes, seconds} used for comparing two different time values
// against each other
func prepareTimestamp(t Timestamp) int {
	h, _ := strconv.Atoi(t.hours)
	m, _ := strconv.Atoi(t.minutes)
	s, _ := strconv.Atoi(t.seconds)

	minutes := (h * 60) + m
	seconds := s + (minutes * 60)
	timestamp := (seconds * 1000)

	return timestamp
}

// When we get a bad timestamp value we take a previous good timestamp
// and update the value by 2:30
func incrementTimestamp(t Timestamp) (int, int, int) {
	h, _ := strconv.Atoi(t.hours)
	m, _ := strconv.Atoi(t.minutes)
	s, _ := strconv.Atoi(t.seconds)

	// increment times by 2:30 for safety
	m += 2
	s += 30

	if m > 60 {
		h += 1
		m = m % 60
	}

	if s > 60 {
		m += 1
		s = s % 60
	}

	return h, m, s
}

func CreateErrorsFile(e []Errorstamp, file string) {

	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	var s string

	s += fmt.Sprintf("[Volume %d]\n", e[0].chapter)
	for _, t := range e {
		s += fmt.Sprintf("row: %d, title: %s\n", t.row, t.docLine)
	}

	if _, err := f.WriteString(s); err != nil {
		log.Println(err)
	}
}
