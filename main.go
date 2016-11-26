package main

import (
	"log"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	// Initialize GTK without parsing any command line arguments.
	gtk.Init(nil)

	// Create a builder to load Glade GTK3 file
	builder, err := gtk.BuilderNew()
	if err != nil {
		//t.Error("Unable to create builder")
	}

	str := `
<interface>
  <!-- interface-requires gtk+ 3.0 -->
  <object class="GtkWindow" id="window">
    <property name="can_focus">False</property>
    <signal name="destroy" handler="on_window_destroy" swapped="no"/>
    <child>
      <object class="GtkLabel" id="label">
        <property name="width_request">250</property>
        <property name="height_request">100</property>
        <property name="visible">True</property>
        <property name="can_focus">False</property>
        <property name="label" translatable="yes">Hello, World!</property>
      </object>
    </child>
  </object>
</interface>
`
	// Read the Glade GTK3 file
	err = builder.AddFromString(str)
	if err != nil {
		log.Fatal("Unable to add from string")
	}

	// widget, err := builder.GetObject("ok_button")
	// if err != nil {
	// 	//t.Error("Unable to get widget from string")
	// }

	// button, ok := widget.(*gtk.Button)
	// if !ok {
	// 	//t.Error("Unable to cast to gtk.Button")
	// }

	// l, err := button.GetLabel()
	// if err != nil {
	// 	//t.Error("Unable to get button label")
	// }

	// dialgWidget, err := builder.GetObject("dialog1")
	// if err != nil {

	// }

	// dialg, ok := dialgWidget.(*gtk.Dialog)
	// if !ok {

	// }

	// if l != "gtk-ok" {
	// 	//t.Errorf("Label has the wrong value: %q", l)
	// }

	// done := make(chan bool)

	// builder.ConnectSignals(map[string]interface{}{
	// 	"ok_button_clicked": func() {
	// 		done <- true
	// 	},
	// })

	//go button.Emit("clicked")

	// Add the label to the window.
	//win.Add(dialg)

	// Get the Window object from the Glade file.  Object named "window"
	winWdg, err := builder.GetObject("window")
	if err != nil {

	}

	// Create the GTK window object
	win, ok := winWdg.(*gtk.Window)
	if !ok {
		log.Println("Error creating the window object")
	}

	// Set the shutdown process when the window is closed
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// Set the Window Name
	// Make sure this name and the CSS name match
	win.SetName("MyWindow")

	// Set the title of the window
	win.SetTitle("Echo v2 - RoweTech Inc")

	// Set the default window size.
	win.SetDefaultSize(800, 600)

	// Create a Style Provider to load CSS data
	styleProvider, err := gtk.CssProviderNew()
	if err != nil {
		log.Println("Style Provider NEW error")
	}

	// Temp CSS
	css := `
    #MyWindow {
        background-color: #F00;
    }
    `

	// Load the CSS style
	styleProvider.LoadFromData(css)

	// Get the default screen
	screen, err := gdk.ScreenGetDefault()
	if err != nil {
		log.Println("Error Getting GDK Screen")
	}

	// Add the Style to the screen
	gtk.AddProviderForScreen(screen, styleProvider, gtk.STYLE_PROVIDER_PRIORITY_USER)

	// Recursively show all widgets contained in this window.
	win.ShowAll()

	// Begin executing the GTK main loop.  This blocks until
	// gtk.MainQuit() is run.
	gtk.Main()

}
