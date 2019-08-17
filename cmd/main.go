package main

import (
	"fmt"
	"gncdu"
	"os"
)

func main() {
	dir := "."
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}

	gncdu.ShowUI(func() []*gncdu.FileData {
		files, err := gncdu.ScanDir(dir)
		if err != nil {
			fmt.Println(err)
		}

		return files
	})
}
