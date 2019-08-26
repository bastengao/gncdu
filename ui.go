package gncdu

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func ShowUI(scanDir func() []*FileData) {
	app := tview.NewApplication()

	done := make(chan bool)
	go func() {
		files := scanDir()
		close(done)
		app.QueueUpdateDraw(func() {
			showResultTable(app, files, nil)
		})
	}()

	scanningView := newScanningView(app, done)
	if err := app.SetRoot(scanningView, true).SetFocus(scanningView).Run(); err != nil {
		panic(err)
	}
}

func newScanningView(app *tview.Application, done chan bool) tview.Primitive {
	info := tview.NewTextView().
		SetText("[ctrl+c] close")

	modal := tview.NewModal().
		SetText("Scanning       \n\nTime 0s")

	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(modal, 0, 1, true).
		AddItem(info, 1, 1, false)

	go func() {
		tick := time.Tick(time.Millisecond * 500)
		dots := []byte{'.', '.', '.', '.', '.', '.'}
		spaces := []byte{' ', ' ', ' ', ' ', ' ', ' '}
		count := 0
		for {
			select {
			case <-tick:
				count++
				p := count % 7
				s := string(dots[0:p])
				b := string(spaces[0:(6 - p)])
				app.QueueUpdateDraw(func() {
					modal.SetText(fmt.Sprintf("Scanning %s%s\n\nTime %ds", s, b, count/2))
				})
			case <-done:
				return
			}
		}
	}()

	return layout
}

func showResultList(app *tview.Application, files []*FileData, parent *FileData) {
	list := tview.NewList().
		ShowSecondaryText(false)

	var title string
	if parent != nil {
		title = parent.Path()
		list = list.AddItem("...", "", ' ', func() {
			if parent.parent.root() {
				showResultList(app, parent.parent.Children, nil)
			} else {
				showResultList(app, parent.parent.Children, parent.parent)
			}
		})
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Size() > files[j].Size()
	})

	format := formatter(files)

	for _, file := range files {
		name := fmt.Sprintf(format, file.info.Name(), ToHumanSize(file.Size()), file.Count())
		list = list.AddItem(name, "", ' ', func(f *FileData) func() {
			return func() {
				if len(f.Children) <= 0 {
					return
				}
				showResultList(app, f.Children, f)
			}
		}(file))
	}
	app.SetRoot(newLayout(title, list), true)
}

func showResultTable(app *tview.Application, files []*FileData, parent *FileData) {
	sort.Slice(files, func(i, j int) bool {
		return files[i].Size() > files[j].Size()
	})

	offset := 1
	var title string
	if parent != nil {
		offset = 2
		title = parent.Path()
	}

	table := tview.NewTable().
		SetFixed(1, 1).
		SetSelectable(true, false).
		SetSelectedFunc(func(row, column int) {
			if row == 0 {
				return
			}

			if row == offset-1 {
				if parent.parent.root() {
					showResultTable(app, parent.parent.Children, nil)
				} else {
					showResultTable(app, parent.parent.Children, parent.parent)
				}
				return
			}

			file := files[row-offset]
			if !file.info.IsDir() {
				return
			}
			showResultTable(app, file.Children, file)
		})

	color := tcell.ColorYellow
	table.SetCell(0, 0, tview.NewTableCell("Name").SetTextColor(color).SetSelectable(false))
	table.SetCell(0, 1, tview.NewTableCell("Size").SetTextColor(color).SetSelectable(false))
	table.SetCell(0, 2, tview.NewTableCell("Items").SetTextColor(color).SetSelectable(false))

	if parent != nil {
		table.SetCellSimple(1, 0, "...")
	}

	for i, file := range files {
		table.SetCellSimple(i+offset, 0, file.info.Name())
		table.SetCell(i+offset, 1,
			tview.NewTableCell(ToHumanSize(file.Size())).
				SetAlign(tview.AlignRight))
		table.SetCell(i+offset, 2,
			tview.NewTableCell(strconv.Itoa((file.Count()))).
				SetAlign(tview.AlignRight))
	}

	layout := newLayout(title, table)
	app.SetRoot(layout, true).SetFocus(layout)
}

func newLayout(title string, content tview.Primitive) tview.Primitive {
	t := tview.NewTextView().
		SetText(title)
	info := tview.NewTextView().
		SetText("[ctrl+c] close")

	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(t, 1, 1, false).
		AddItem(content, 0, 1, true).
		AddItem(info, 1, 1, false)

	return layout
}
