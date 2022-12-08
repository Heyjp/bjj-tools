package video_merge

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	d "github.com/heyjp/bjj-tools/dircheck"
	e "github.com/heyjp/bjj-tools/extractor"
	m "github.com/heyjp/bjj-tools/metadata"
)

func Merge() {
	f := d.GetDirectoryFiles()

	// Find the directory with the chapters files "chapters-1.txt
	// chapters-2.txt
	var chapterDirectory string
	res, _ := os.ReadDir(".")

	for _, file := range res {
		if file.IsDir() {
			if file.Name() == "chapters" {
				chapterDirectory = file.Name()
				break
			}

			dir, _ := os.ReadDir(file.Name())

			var matches []string
			for _, subFile := range dir {
				if matched, _ := regexp.MatchString(`chapters*.txt`, subFile.Name()); matched {
					matches = append(matches, subFile.Name())
				}
			}

			if len(f) == len(matches) {
				chapterDirectory = file.Name()
				break
			}
		}
	}

	if chapterDirectory == "" {
		log.Fatal("Chapter Directory not found")
	}

	// d.SortStrings(f)
	d.CheckOrCreateDirectory("metadata")

	l, _ := os.Getwd()
	l += "/metadata"

	// Prepare directory for Outputted files
	o := "output"
	d.CheckOrCreateDirectory(o)

	n := "chapters-"
	for i, file := range f {
		metaFile := e.ExtractMetadataFromVideo(file, l)
		// Combine metadata with chapters
		c := fmt.Sprintf("%s/%s%d.txt", chapterDirectory, n, i+1)
		m.MergeChaptersWithMetadata(metaFile, c)

		s := strings.SplitN(file, ".", 2)

		// Remerge metadata with the video file
		outputLocation := fmt.Sprintf("%s/%s.mp4", o, s[0])
		e.Combine(file, metaFile, outputLocation)
	}

	fmt.Println("Complete")
}
