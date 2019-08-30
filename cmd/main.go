package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/bastengao/gncdu"
)

func main() {
	dir := "."
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}
	dir, err := filepath.Abs(dir)
	if err != nil {
		fmt.Println(err)
		return
	}

	gncdu.ShowUI(func() ([]*gncdu.FileData, error) {
		files, err := gncdu.ScanDirConcurrent(dir)

		if err != nil {
			return nil, err
		}

		return files, nil
	})
}
