package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/bastengao/gncdu/scan"
	"github.com/bastengao/gncdu/ui"
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

	ui.ShowUI(func() ([]*scan.FileData, error) {
		files, err := scan.ScanDirConcurrent(dir)

		if err != nil {
			return nil, err
		}

		return files, nil
	})
}
