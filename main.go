//go:build !gui

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gotk3/gotk3/gtk"
)

func main() {
	args := os.Args

	if len(args) == 2 {
		// CLI
		c, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
			log.Fatalf("can't convert %v", args[1])
		}

		fmt.Printf("max: %.2f\n", c*1.25)
		fmt.Printf("min: %.2f\n", c*0.92)

	} else if len(args) == 1 {
		// GUI
		gtk.Init(nil)

		// Create a new toplevel wndow, set its title, and connect it to the
		// "destroy" signal to exit the GTK main loop when it is destroyed.
		win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
		if err != nil {
			log.Fatal("Unable to create window:", err)
		}

		win.SetTitle("qCal")
		win.Connect("destroy", func() {
			gtk.MainQuit()
		})

		// Create a new grid widget to arrange child widgets
		grid, err := gtk.GridNew()
		if err != nil {
			log.Fatal("Unable to create grid:", err)
		}

		grid.SetOrientation(gtk.ORIENTATION_VERTICAL)
		grid.SetRowSpacing(10)
		grid.SetColumnSpacing(10)

		lab1, err := gtk.LabelNew("输入期望计算值")
		if err != nil {
			log.Fatal("Unable to create label:", err)
		}

		entry1, err := gtk.EntryNew()
		if err != nil {
			log.Fatal("Unable to create entry:", err)
		}

		grid.Attach(lab1, 0, 0, 1, 1)
		grid.Attach(entry1, 0, 1, 1, 1)

		// Add the grid to the window, and show all widgets.
		win.Add(grid)

		win.ShowAll()

		gtk.Main()
	}
}
