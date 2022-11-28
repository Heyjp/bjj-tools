package main

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strings"
)

// Function takes a parameter which points to the location of a video
// file and extracts the metadata file to a location
func ExtractMetaDataFromVideo(v string, l string) {
	var name = "-metadata.txt"

	// Extracts the name of a file from a location string
	// i.e text/video.mp4
	re := regexp.MustCompile(`[^//]*$`)
	res := re.FindStringSubmatch(v)

	// Creates a substring of the file name excluding the file type
	index := strings.Index(res[0], ".")
	substr := res[0][0:index]

	// appends "-metadata.txt" to the substring
	metadataString := substr + name
	// Location where the metadata will go
	metadataLocation := l + metadataString

	cmdString := fmt.Sprintf("ffmpeg -i %s -f ffmetadata %s", v, metadataLocation)
	cmd := exec.Command(cmdString)
	e := cmd.Run()

	if e != nil {
		log.Fatal(e)
	}
}

// Takes a video, metadata and output location and passes them to
// ffmpeg
func Combine(video string, metadata string, output string) {
	cmdString := fmt.Sprintf("ffmpeg -i %s -i %s -map_metadata 1 -codec copy %s", video, metadata, output)
	cmd := exec.Command(cmdString)
	e := cmd.Run()

	if e != nil {
		log.Fatal(e)
	}
}
