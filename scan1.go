package gncdu

import (
	"io/ioutil"
)

func ScanDirDirectly(dir string) ([]*FileData, error) {
	root := newRootFileData(dir)
	err := scanDir11(root)
	if err != nil {
		return nil, err
	}

	return root.Children, nil
}

func scanDir11(parent *FileData) error {
	if !parent.root() && (parent.size != -1 || !parent.info.IsDir()) {
		return nil
	}

	files, err := ioutil.ReadDir(parent.Path())
	if err != nil {
		return err
	}

	children := []*FileData{}
	for _, file := range files {
		f := newFileData(parent, file)
		scanDir11(f)

		children = append(children, f)
	}

	parent.Children = children
	return nil
}
