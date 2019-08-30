package gncdu

import (
	"io/ioutil"
	"sync"
)

func ScanDirConcurrent(dir string) ([]*FileData, error) {
	root := newRootFileData(dir)

	ch := make(chan *FileData)
	closeWait := &sync.WaitGroup{}

	var wait sync.WaitGroup
	wait.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for file := range ch {
				scanDir(file, ch, closeWait)
				closeWait.Done()
			}
			wait.Done()
		}()
	}

	err := scanDir(root, ch, closeWait)
	if err != nil {
		return nil, err
	}

	go func() {
		closeWait.Wait()
		close(ch)
	}()

	wait.Wait()

	return root.Children, nil
}

func scanDir(parent *FileData, ch chan *FileData, closeWait *sync.WaitGroup) error {
	if !parent.root() && (parent.size != -1 || !parent.info.IsDir()) {
		return nil
	}

	files, err := ioutil.ReadDir(parent.Path())
	if err != nil {
		return err
	}

	children := []*FileData{}
	closeWait.Add(len(files))
	for _, file := range files {
		f := newFileData(parent, file)
		go func() {
			ch <- f
		}()

		children = append(children, f)
	}

	parent.Children = children
	return nil
}
