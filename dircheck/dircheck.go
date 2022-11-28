package dircheck

import (
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func GetDirectoryFiles() []string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	files, err := os.ReadDir(path)
	re := regexp.MustCompile(`.*\.[mp4|mpeg|mkv]+$`)

	var v []string
	for _, file := range files {
		match := re.FindString(file.Name())
		if match != "" {
			v = append(v, match)
		}
	}

	return v
}

func SortStrings(data []string) []string {
	sort.Slice(data, func(i, j int) bool {
		// Extract first instance of a number inside of a string
		re := regexp.MustCompile(`^[^\d]*(\d+)`)
		a := data[i]
		b := data[j]

		aM := re.FindStringSubmatch(a)
		bM := re.FindStringSubmatch(b)

		if aM == nil || bM == nil {
			return true
		}

		aI, errA := strconv.ParseInt(aM[1], 10, 64)
		bI, errB := strconv.ParseInt(bM[1], 10, 64)

		if errA != nil || errB != nil {
			return true
		}

		return aI < bI
	})
	return data
}

func CheckOrCreateDirectory(location string) {
	_, errF := os.Stat(location)

	if os.IsNotExist(errF) {
		er := os.MkdirAll(location, 0755)
		if er != nil {
			log.Fatal(er)
		}
	}
}
