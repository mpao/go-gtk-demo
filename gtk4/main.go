package main

import (
	_ "embed"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
	"log"
	"os"
	"strconv"
)

//go:embed main.ui
var uiXML string

func main() {
	app := gtk.NewApplication("io.github.mpao", gio.ApplicationFlagsNone)

	app.ConnectStartup(func() {
		log.Println("application start up")
	})
	app.ConnectActivate(func() {
		activate(app)
	})
	app.ConnectShutdown(func() {
		log.Println("application shutdown")
	})

	if code := app.Run(os.Args); code > 0 {
		os.Exit(code)
	}
}

func activate(app *gtk.Application) {

	builder := gtk.NewBuilderFromString(uiXML, len(uiXML))

	// MainWindow and Button are object IDs from the UI file
	window := builder.GetObject("MainWindow").Cast().(*gtk.Window)
	button := builder.GetObject("Button").Cast().(*gtk.Button)

	counter := 0
	button.Connect("clicked", func() {
		button.SetLabel(strconv.Itoa(counter))
		counter++
	})

	app.AddWindow(window)
	window.Show()
}
