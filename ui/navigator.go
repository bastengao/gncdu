package ui

var navigator Navigator

type Navigator struct {
	previouse Page
	current   Page
}

func (n *Navigator) Push(page Page) {
	n.previouse = n.current
	n.current = page
	if n.previouse != nil {
		n.previouse.Dispose()
	}
	n.current.SetNavigator(n)
	n.current.SetPrevious(n.previouse)
	n.current.Show()
}

func (n *Navigator) Pop() {
	if n.previouse != nil {
		n.current.Dispose()
		n.current = n.previouse
		n.previouse = n.current.Previous()
		n.current.Show()
	}
}
