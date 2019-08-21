package gncdu

import (
	"os"
)

type FileData struct {
	parent   *FileData
	dir      string
	info     os.FileInfo
	size     int64
	Children []*FileData
	count    int
}

func newRootFileData(dir string) *FileData {
	return &FileData{dir: dir, size: 0, count: 0}
}

func newFileData(parant *FileData, file os.FileInfo) *FileData {
	var size int64 = -1
	count := -1
	if !file.IsDir() {
		size = file.Size()
		count = 0
	}
	return &FileData{parent: parant, dir: parant.Path(), info: file, size: size, count: count}
}

func (d FileData) root() bool {
	return d.info == nil
}

func (d FileData) Path() string {
	if d.root() {
		return d.dir
	}

	return d.dir + "/" + d.info.Name()
}

func (d FileData) String() string {
	return d.Path()
}

func (d *FileData) Count() int {
	if d.count != -1 {
		return d.count
	}
	c := len(d.Children)
	for _, f := range d.Children {
		c += f.Count()
	}
	d.count = c
	return c
}

func (d *FileData) Size() int64 {
	if d.size != -1 {
		return d.size
	}

	s := d.info.Size()
	for _, f := range d.Children {
		s += f.Size()
	}
	d.size = s
	return s
}

func hasDir(files []os.FileInfo) bool {
	for _, file := range files {
		if file.IsDir() {
			return true
		}
	}
	return false
}
