package main

import (
	"os"

	c "github.com/heyjp/bjj-tools/chapters"
)

func main() {
	timestamps := c.PrepareTimestamps(os.Args[1])
	c.CreateChaptersFile(timestamps, "chapters", false)
}
