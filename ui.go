package gncdu

import (
	"fmt"
	"strconv"

	"github.com/gdamore/tcell"
	"github.com/gdamore/tcell/views"
)

type App struct {
	app   *views.Application
	view  views.View
	panel views.Widget

	views.WidgetWatchers
}

func (a *App) Draw() {
	if a.panel != nil {
		a.panel.Draw()
	}
}

func (a *App) Resize() {
	if a.panel != nil {
		a.panel.Resize()
	}
}

func (a *App) HandleEvent(ev tcell.Event) bool {
	switch ev := ev.(type) {
	case *tcell.EventKey:
		switch ev.Key() {
		// Intercept a few control keys up front, for global handling.
		case tcell.KeyCtrlC:
			a.app.Quit()
			return true
		case tcell.KeyCtrlL:
			a.app.Refresh()
			return true
		}
	}
	return false
}

func (a *App) SetView(v views.View) {
	if a.panel != nil {
		a.panel.SetView(v)
	}
}

func (a *App) Size() (int, int) {
	if a.panel != nil {
		return a.panel.Size()
	}
	return 0, 0
}

func (a *App) Run() {
	fmt.Println("run")
	a.app.SetRootWidget(a)
	a.panel = mainView()
	a.app.Run()
}

func NewApp() *App {
	app := &App{}
	app.app = &views.Application{}
	return app
}

func mainView() views.Widget {
	view := &views.Panel{}
	titleBar := &views.SimpleStyledTextBar{}
	titleBar.SetCenter("Hello world")
	view.SetTitle(titleBar)

	statusBar := &views.SimpleStyledTextBar{}
	statusBar.SetCenter("[ctrl+c] close")
	view.SetStatus(statusBar)

	text := views.NewTextArea()
	lines := []string{"Helo", "world"}
	for i := 0; i < 20; i++ {
		lines = append(lines, strconv.Itoa(i))
	}
	text.SetLines(lines)
	text.EnableCursor(true)
	view.SetContent(text)
	return view
}
