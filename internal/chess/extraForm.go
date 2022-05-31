package chess

import (
	"github.com/gotk3/gotk3/gtk"

	"github.com/hultan/softteam/framework"
)

type ExtraForm struct {
	Window *gtk.Window
}

func NewExtraForm() *ExtraForm {
	return new(ExtraForm)
}

func (e *ExtraForm) OpenForm(fw *framework.Framework) {
	// Create a new gtk helper
	builder, err := fw.Gtk.CreateBuilder("main.glade")
	if err != nil {
		panic(err)
	}
	// Get the extra window from glade
	extraWindow := builder.GetObject("extra_window").(*gtk.Window)

	// Set up the extra window
	extraWindow.SetTitle("extra form")
	extraWindow.HideOnDelete()
	extraWindow.SetModal(true)
	extraWindow.SetKeepAbove(true)
	extraWindow.SetPosition(gtk.WIN_POS_CENTER_ALWAYS)

	// Hook up the destroy event
	extraWindow.Connect("destroy", extraWindow.Destroy)

	// Close button
	button := builder.GetObject("extra_window_close_button").(*gtk.Button)
	button.Connect("clicked", extraWindow.Destroy)

	e.Window = extraWindow

	// Show the window
	extraWindow.ShowAll()
}
