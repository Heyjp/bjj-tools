package main

import (
	d "github.com/heyjp/bjj-tools/dircheck"
)

func main() {
	f := d.GetDirectoryFiles()
	d.SortStrings(f)
}
