package chess

import (
	"github.com/gotk3/gotk3/gtk"

	"github.com/hultan/softteam/framework"
)

type Dialog struct {
	Dialog *gtk.Dialog
}

func NewDialog() *Dialog {
	return new(Dialog)
}

func (d *Dialog) OpenDialog(parent gtk.IWindow, fw *framework.Framework) {
	// Create a new softBuilder
	builder, err := fw.Gtk.CreateBuilder("main.glade")
	if err != nil {
		panic(err)
	}

	// Get the dialog window from glade
	dialog := builder.GetObject("settings_dialog").(*gtk.Dialog)

	dialog.SetTitle("settings dialog")
	dialog.SetTransientFor(parent)
	dialog.SetModal(true)

	dialog.Connect("response", func(dialog *gtk.Dialog, responseId gtk.ResponseType) {
		if responseId == gtk.RESPONSE_CANCEL || responseId == gtk.RESPONSE_DELETE_EVENT {
			dialog.Hide()
		} else if responseId == gtk.RESPONSE_ACCEPT || responseId == gtk.RESPONSE_OK || responseId == gtk.RESPONSE_YES {
			// Save
			dialog.Hide()
		}
	})

	dialog.Present()
}
