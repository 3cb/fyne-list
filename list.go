package list

import (
	"image/color"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

// List widget is a scrollable list of fyne.CanvasObjects with a header at the top.
type List struct {
	widget.BaseWidget

	Header  fyne.CanvasObject
	box     *widget.Box
	content fyne.CanvasObject
}

// NewList creates a new grouped list widget with a header and the specified list of child objects
func NewList(header fyne.CanvasObject, children ...fyne.CanvasObject) *List {
	box := widget.NewVBox(children...)
	list := &List{widget.BaseWidget{}, header, box, box}
	return list
}

// NewListWithScroller creates a new grouped list widget with a title and
// the specified list of child objects.
// This list will scroll when the available space is less than needed to display the
// items it contains.
func NewListWithScroller(header fyne.CanvasObject, children ...fyne.CanvasObject) *List {
	box := widget.NewVBox(children...)
	list := &List{widget.BaseWidget{}, header, box, widget.NewScrollContainer(box)}
	return list
}

// Prepend inserts a new CanvasObject at the top of the group
// Method returns the index of the added object
func (l *List) Prepend(object fyne.CanvasObject) int {
	l.box.Prepend(object)

	canvas.Refresh(l)
	return 0
}

// Append adds a new CanvasObject to the end of the group
// Method returns the index of the added object
func (l *List) Append(object fyne.CanvasObject) int {
	l.box.Append(object)

	canvas.Refresh(l)
	return len(l.box.Children) - 1
}

// Pop removes last object from the List
func (l *List) Pop() {
	l.box.Children = l.box.Children[:len(l.box.Children)-1]
}

// GetRow uses index to return object
func (l *List) GetRow(i int) fyne.CanvasObject {
	return l.box.Children[i]
}

// Remove deletes row from List
func (l *List) Remove(i int) {
	l.box.Children = append(l.box.Children[:i], l.box.Children[i+1:]...)
}

// MinSize returns the size that this widget should not shrink below
func (l *List) MinSize() fyne.Size {
	l.ExtendBaseWidget(l)
	return l.BaseWidget.MinSize()
}

// CreateRenderer is a private method to Fyne which links this widget to its renderer
func (l *List) CreateRenderer() fyne.WidgetRenderer {
	l.ExtendBaseWidget(l)
	header := l.Header
	objects := []fyne.CanvasObject{header, l.content}
	return &listRenderer{header: header, objects: objects, list: l}
}

type listRenderer struct {
	header fyne.CanvasObject

	objects []fyne.CanvasObject
	list    *List
}

func (l *listRenderer) MinSize() fyne.Size {
	headerMin := l.header.MinSize()
	listMin := l.list.content.MinSize()

	return fyne.NewSize(fyne.Max(headerMin.Width, listMin.Width), headerMin.Height+listMin.Height+theme.Padding())
}

func (l *listRenderer) Layout(size fyne.Size) {
	headerHeight := l.header.MinSize().Height

	l.header.Move(fyne.NewPos(0, 0))
	l.header.Resize(fyne.NewSize(size.Width, headerHeight))

	l.list.content.Move(fyne.NewPos(0, headerHeight+theme.Padding()))
	l.list.content.Resize(fyne.NewSize(size.Width, size.Height-headerHeight-theme.Padding()))
}

func (l *listRenderer) BackgroundColor() color.Color {
	return theme.BackgroundColor()
}

func (l *listRenderer) Objects() []fyne.CanvasObject {
	return l.objects
}

func (l *listRenderer) Refresh() {
	l.Layout(l.list.Size())

	l.header.Refresh()
	l.list.content.Refresh()
}

func (l *listRenderer) Destroy() {
}
