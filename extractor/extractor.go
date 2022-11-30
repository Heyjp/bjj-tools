package extractor

import (
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

// Function takes a parameter which points to the location of a video
// file and extracts the metadata file to a location
func ExtractMetadataFromVideo(v string, l string) string {
	var name = "-metadata.txt"

	// Extracts the name of a file from a location string
	// i.e text/video.mp4
	re := regexp.MustCompile(`[^\/]*$`)
	res := re.FindStringSubmatch(v)

	// Creates a substring of the file name excluding the file type
	index := strings.Index(res[0], ".")
	substr := res[0][0:index]

	// appends "-metadata.txt" to the substring
	metadataString := substr + name
	// Location where the metadata will go
	metadataLocation := l + "/" + metadataString

	if _, err := os.Stat(metadataLocation); err == nil {
		return metadataLocation
	}
	// runs ffmpeg taking in a file location and outputting a metadata
	// file
	cmd := exec.Command("ffmpeg", "-i", v, "-f", "ffmetadata", metadataLocation)
	e := cmd.Run()

	if e != nil {
		log.Fatal(e)
	}

	return metadataLocation
}

// Takes a video, metadata and output location and passes them to
// ffmpeg and outputs a new mp4 file
func Combine(video string, metadata string, output string) {
	if _, err := os.Stat(output); err == nil {
		os.Remove(output)
	}
	cmd := exec.Command("ffmpeg", "-i", video, "-i", metadata, "-map_metadata", "1", "-codec", "copy", output)
	e := cmd.Run()

	if e != nil {
		log.Fatal(e)
	}
}
