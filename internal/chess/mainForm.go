package chess

import (
	"os"

	"github.com/gotk3/gotk3/gtk"

	"github.com/hultan/softteam/framework"
)

const applicationTitle = "chess"
const applicationVersion = "v 0.01"
const applicationCopyRight = "©SoftTeam AB, 2020"

type MainForm struct {
	Window      *gtk.ApplicationWindow
	builder     *framework.GtkBuilder
	AboutDialog *gtk.AboutDialog
}

// NewMainForm : Creates a new MainForm object
func NewMainForm() *MainForm {
	mainForm := new(MainForm)
	return mainForm
}

// OpenMainForm : Opens the MainForm window
func (m *MainForm) OpenMainForm(app *gtk.Application) {
	// Initialize gtk
	gtk.Init(&os.Args)

	// Create a new softBuilder
	fw := framework.NewFramework()
	builder, err := fw.Gtk.CreateBuilder("main.glade")
	if err != nil {
		panic(err)
	}
	m.builder = builder

	// Get the main window from the glade file
	m.Window = m.builder.GetObject("main_window").(*gtk.ApplicationWindow)

	// Set up main window
	m.Window.SetApplication(app)
	m.Window.SetTitle("chess main window")

	// Hook up the destroy event
	m.Window.Connect("destroy", m.Window.Close)

	// Quit button
	button := m.builder.GetObject("main_window_quit_button").(*gtk.ToolButton)
	button.Connect("clicked", m.Window.Close)

	// Status bar
	statusBar := m.builder.GetObject("main_window_status_bar").(*gtk.Statusbar)
	statusBar.Push(statusBar.GetContextId("chess"), "chess : version 0.1.0")

	// Open form button
	openFormButton := m.builder.GetObject("main_window_open_form_button").(*gtk.Button)
	openFormButton.Connect("clicked", func() {
		m.OpenForm(fw)
	})

	// Open dialog button
	openDialogButton := m.builder.GetObject("main_window_open_dialog_button").(*gtk.Button)
	openDialogButton.Connect("clicked", func() {
		m.OpenDialog(fw)
	})

	// Menu
	m.setupMenu(fw)

	// Show the main window
	m.Window.ShowAll()
}

func (m *MainForm) OpenForm(fw *framework.Framework) {
	extraForm := NewExtraForm()
	extraForm.OpenForm(fw)
}

func (m *MainForm) OpenDialog(fw *framework.Framework) {
	dialog := NewDialog()
	dialog.OpenDialog(m.Window, fw)
}

func (m *MainForm) setupMenu(fw *framework.Framework) {
	menuQuit := m.builder.GetObject("menu_file_quit").(*gtk.MenuItem)
	menuQuit.Connect("activate", m.Window.Close)

	menuHelpAbout := m.builder.GetObject("menu_help_about").(*gtk.MenuItem)
	menuHelpAbout.Connect("activate", func() {
		m.openAboutDialog(fw)
	})
}
