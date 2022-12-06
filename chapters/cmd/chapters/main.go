package main

import (
	"os"

	c "github.com/heyjp/bjj-tools/chapters"
	d "github.com/heyjp/bjj-tools/dircheck"
)

func main() {
	timestamps, errorStamps := c.PrepareTimestamps(os.Args[1])
	c.CreateChaptersFile(timestamps, "chapters", false)
	if len(errorStamps) > 0 {
		location := "chapters/errors/"
		d.CheckOrCreateDirectory(location)
		c.CreateErrorsFile(errorStamps, location)
	}
}
