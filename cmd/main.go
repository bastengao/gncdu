package main

import (
	"fmt"
	"github.com/bastengao/gncdu"
	"os"
	"path/filepath"
)

func main() {
	dir := "."
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}
	dir, err := filepath.Abs(dir)
	if err != nil {
		fmt.Println(err)
	}

	gncdu.ShowUI(func() []*gncdu.FileData {
		files, err := gncdu.ScanDirConcurrent(dir)

		if err != nil {
			fmt.Println(err)
		}

		return files
	})
}
