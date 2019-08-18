package gncdu

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Print(files []*FileData) {
	sort.Slice(files, func(i, j int) bool {
		return files[i].Size() > files[j].Size()
	})

	format := formatter(files)
	for _, f := range files {
		fmt.Sprintf(format, f.Path(), ToHumanSize(f.Size()), f.Count())
	}
}

func formatter(files []*FileData) string {
	maxLenght := 0
	for _, f := range files {
		l := strings.Count(f.info.Name(), "")
		if l > maxLenght {
			maxLenght = l
		}
	}
	return "%-" + strconv.Itoa(maxLenght) + "s" + " %10s  items %-5d\n"
}
