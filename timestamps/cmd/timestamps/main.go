package main

import (
	"os"

	c "github.com/heyjp/bjj-tools/timestamps"
)

func main() {
	t, errorStamps := c.PrepareTimestamps(os.Args[1])
	c.CreateChaptersFile(t, os.Args[1], false)
	if len(errorStamps) > 0 {
		c.CreateErrorsFile(errorStamps, os.Args[1])
	}
}
