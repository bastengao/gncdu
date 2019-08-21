package gncdu

import (
	"io/ioutil"
	"os"
	"sync"
)

func ScanDirDynamic(dir string) ([]*FileData, error) {
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

	root := newRootFileData(dir)
	for _, file := range files {
		fileData := newFileData(root, file)
		if file.IsDir() {
			ch <- fileData
		}

		data = append(data, fileData)
	}

	root.Children = data
	for _, file := range data {
		file.parent = root
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

	if len(files) == 0 || !hasDir(files) {
		d.scanDirectly(files)
	} else {
		d.scanConcurrent(files)
	}
	return nil
}

func (d *FileData) scanDirectly(files []os.FileInfo) {
	children := []*FileData{}
	for _, file := range files {
		fileData := newFileData(d, file)
		fileData.scan()

		children = append(children, fileData)
	}

	d.Children = children
}

func (d *FileData) scanConcurrent(files []os.FileInfo) {
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
