package list

import (
	"image/color"
	"sort"

	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

// Header is a fyne.Widget that allows multiple strings in the header of
// a List object
type Header struct {
	widget.BaseWidget

	// labels will be displayed in order left to right
	labels []string
	color  color.Color
}

// NewHeader creates a new instance of Header widget
func NewHeader(color color.Color, labels ...string) *Header {
	h := &Header{widget.BaseWidget{}, labels, color}
	h.ExtendBaseWidget(h)
	return h
}

// CreateRenderer is a private method to Fyne which links this widget to its renderer
func (h *Header) CreateRenderer() fyne.WidgetRenderer {
	h.ExtendBaseWidget(h)
	objects := []fyne.CanvasObject{}
	style := fyne.TextStyle{
		Bold: true,
	}
	for _, label := range h.labels {
		obj := canvas.NewText(label, h.color)
		obj.Alignment = fyne.TextAlignTrailing
		obj.TextStyle = style
		objects = append(objects, obj)
	}
	// add 5 space margin on right side
	margin := canvas.NewText("     ", h.color)
	objects = append(objects, margin)

	return &headerRenderer{objects, margin, h}
}

type headerRenderer struct {
	objects []fyne.CanvasObject
	margin  *canvas.Text
	header  *Header
}

func (h *headerRenderer) MinSize() fyne.Size {
	minWidths := []int{}
	marginWidth := h.margin.MinSize().Width
	minWidths = append(minWidths, marginWidth)
	for _, o := range h.objects {
		minWidths = append(minWidths, o.MinSize().Width)
	}
	sort.Ints(minWidths)
	return fyne.NewSize(3*(minWidths[len(minWidths)-1])+marginWidth, h.objects[1].MinSize().Height)
}

func (h *headerRenderer) Layout(size fyne.Size) {
	numObj := len(h.objects)
	marWidth := h.margin.MinSize().Width

	// don't loop over margin string
	for i := 0; i < numObj-1; i++ {
		if i == 0 {
			h.objects[i].(*canvas.Text).Move(fyne.NewPos(0, 0))
		} else {
			h.objects[i].(*canvas.Text).Move(fyne.NewPos(((size.Width-marWidth)/(numObj-1))*(i*1), 0))
		}
		h.objects[i].(*canvas.Text).Resize(fyne.NewSize((size.Width-marWidth)/(numObj-1), size.Height))
	}
	h.objects[numObj-1].(*canvas.Text).Move(fyne.NewPos((size.Width - marWidth), 0))
	h.objects[numObj-1].(*canvas.Text).Resize(fyne.NewSize(marWidth, size.Height))
}

func (h *headerRenderer) BackgroundColor() color.Color {
	return theme.BackgroundColor()
}

func (h *headerRenderer) Objects() []fyne.CanvasObject {
	return h.objects
}

func (h *headerRenderer) Refresh() {
	h.Layout(h.header.Size())
	for _, o := range h.objects {
		o.Refresh()
	}
}

func (h *headerRenderer) Destroy() {}
