package main

import (
	"fmt"
	"gncdu"
	"os"
	"time"
)

func main() {
	start := time.Now()
	dir := "."
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}
	fmt.Println(dir)
	files, err := gncdu.ScanDir(dir)
	end := time.Now()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(end.Sub(start))

	gncdu.Print(files)

	// app := gncdu.NewApp()
	// app.Run()
}
