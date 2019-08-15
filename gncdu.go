package gncdu

import (
	"io/ioutil"
	"os"
	"sync"
)

type FileData struct {
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
	ch := make(chan *FileData, 1024)

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
		ch <- fileData

		data = append(data, fileData)
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

	children := []*FileData{}
	for _, file := range files {
		var size int64 = -1
		count := -1
		if !file.IsDir() {
			size = file.Size()
			count = 0
		}
		fileData := &FileData{dir: d.Path(), info: file, size: size, count: count}
		fileData.scan()

		children = append(children, fileData)
	}
	d.Children = children

	return nil
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
