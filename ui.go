package gncdu

import (
	"fmt"
	"sort"

	"github.com/rivo/tview"
)

func ShowUI(scanDir func() []*FileData) {
	app := tview.NewApplication()

	modal := tview.NewModal().
		SetText("Scanning ...")

	go func() {
		files := scanDir()
		app.QueueUpdateDraw(func() {
			showResult(app, files, []*FileData{})
		})
	}()

	if err := app.SetRoot(modal, true).SetFocus(modal).Run(); err != nil {
		panic(err)
	}
}

func showResult(app *tview.Application, files []*FileData, parent []*FileData) {
	list := tview.NewList().
		ShowSecondaryText(false)

	if len(parent) != 0 {
		list = list.AddItem("...", "", ' ', func() {
			showResult(app, parent, []*FileData{})
		})
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Size() > files[j].Size()
	})

	format := formatter(files)

	for _, file := range files {
		name := fmt.Sprintf(format, file.Path(), ToHumanSize(file.Size()), file.Count())
		list = list.AddItem(name, "", ' ', func(f *FileData) func() {
			return func() {
				if len(f.Children) <= 0 {
					return
				}
				showResult(app, f.Children, files)
			}
		}(file))
	}
	app.SetRoot(list, true)
}
