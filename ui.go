package gncdu

import (
	"sync/atomic"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func ShowUI(scanDir func() []*FileData) {
	app := tview.NewApplication()

	var scanningPage Page
	scanningPage = NewScanningPage(app)

	var isDone atomic.Value
	done := make(chan bool)
	go func() {
		files := scanDir()
		close(done)
		isDone.Store(true)
		app.QueueUpdateDraw(func() {
			var resultPage Page
			resultPage = NewResultPage(app, files, nil)
			navigator.Push(resultPage)
		})
	}()

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if isDone.Load() != nil && event.Rune() == '?' {
			help := NewHelpPage(app)
			navigator.Push(help)
		}
		return event
	})

	navigator.Push(scanningPage)
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func newInfoView() tview.Primitive {
	return tview.NewTextView().
		SetText("[ctrl+c] close  [?] help")
}

func newLayout(title string, content tview.Primitive) tview.Primitive {
	t := tview.NewTextView().
		SetText(title)
	info := newInfoView()

	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(t, 1, 1, false).
		AddItem(content, 0, 1, true).
		AddItem(info, 1, 1, false)

	return layout
}
