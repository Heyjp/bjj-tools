package main

import (
	"os"

	c "github.com/heyjp/bjj-tools/chapters"
)

func main() {
	timestamps, errorStamps := c.PrepareTimestamps(os.Args[1])
	c.CreateChaptersFile(timestamps, "chapters", false)
	if len(errorStamps) > 0 {
		c.CreateErrorsFile(errorStamps, "chapters/errors/errors.txt")
	}
}
