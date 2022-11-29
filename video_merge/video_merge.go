package video_merge

import (
	"fmt"
	"os"

	d "github.com/heyjp/bjj-tools/dircheck"
	e "github.com/heyjp/bjj-tools/extractor"
	m "github.com/heyjp/bjj-tools/metadata"
)

func Merge() {
	f := d.GetDirectoryFiles()
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
		c := fmt.Sprintf("chapters/%s%d.txt", n, i+1)
		m.MergeChaptersWithMetadata(metaFile, c)

		// Remerge metadata with the video file
		outputLocation := fmt.Sprintf("%s/%s.mp4", o, file)
		e.Combine(file, metaFile, outputLocation)
	}

	fmt.Println("Complete")
}
