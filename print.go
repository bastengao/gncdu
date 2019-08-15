package gncdu

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Print(files []*FileData) {
	maxLenght := 0
	for _, f := range files {
		l := strings.Count(f.Path(), "")
		if l > maxLenght {
			maxLenght = l
		}
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Size() > files[j].Size()
	})

	format := "%-" + strconv.Itoa(maxLenght) + "s" + " %10s  items %-5d\n"
	for _, f := range files {
		fmt.Printf(format, f.Path(), ToHumanSize(f.Size()), f.Count())
	}
}
