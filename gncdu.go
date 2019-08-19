package gncdu

import (
	"io/ioutil"
	"os"
	"sync"
)

type FileData struct {
	parent   *FileData
	dir      string
	info     os.FileInfo
	size     int64
	Children []*FileData
	count    int
}

func ScanDir(dir string) ([]*FileData, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	data := []*FileData{}
	ch := make(chan *FileData, len(files))

	var wait sync.WaitGroup
	wait.Add(5)
	for i := 0; i < 5; i++ {
		go func(ch chan *FileData) {
			for fileData := range ch {
				fileData.scan()
			}
			wait.Done()
		}(ch)
	}

	for _, file := range files {
		var size int64 = -1
		count := -1
		if !file.IsDir() {
			size = file.Size()
			count = 0
		}
		fileData := &FileData{dir: dir, info: file, size: size, count: count}
		if file.IsDir() {
			ch <- fileData
		}

		data = append(data, fileData)
	}

	parent := &FileData{size: 0, count: 0, Children: data}
	for _, file := range data {
		file.parent = parent
	}

	close(ch)
	wait.Wait()

	return data, nil
}

func (d *FileData) scan() error {
	if d.size != -1 {
		return nil
	}

	if !d.info.IsDir() {
		return nil
	}

	files, err := ioutil.ReadDir(d.Path())
	if err != nil {
		return err
	}

	if (len(files) == 0 || !hasDir(files)) {
		d.ScanDirectly(files)
	} else {
		d.ScanConcurrent(files)
	}
	return nil
}

func (d *FileData) ScanDirectly(files []os.FileInfo) {
	children := []*FileData{}
	for _, file := range files {
		var size int64 = -1
		count := -1
		if !file.IsDir() {
			size = file.Size()
			count = 0
		}
		fileData := &FileData{parent: d, dir: d.Path(), info: file, size: size, count: count}
	    fileData.scan()

		children = append(children, fileData)
	}

	d.Children = children
}

func (d *FileData) ScanConcurrent(files []os.FileInfo) {
	ch := make(chan *FileData, len(files))
	var wait sync.WaitGroup
	wait.Add(5)
	for i := 0; i < 5; i++ {
		go func(ch chan *FileData) {
			for fileData := range ch {
				fileData.scan()
			}
			wait.Done()
		}(ch)
	}

	children := []*FileData{}
	for _, file := range files {
		var size int64 = -1
		count := -1
		if !file.IsDir() {
			size = file.Size()
			count = 0
		}
		fileData := &FileData{parent: d, dir: d.Path(), info: file, size: size, count: count}
		if file.IsDir() {
			ch <- fileData
		}

		children = append(children, fileData)
	}

	close(ch)
	wait.Wait()

	d.Children = children
}

func (d FileData) root() bool {
	return d.info == nil
}

func (d FileData) Path() string {
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
