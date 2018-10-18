package main

import (
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/jroimartin/gocui"
)

var mainGui gocui.Gui

func main() {
	mainGui, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer mainGui.Close()

	mainGui.Cursor = true
	mainGui.SetManagerFunc(layout)

	if err := mainGui.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	if err := mainGui.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

// InputEditor - the input line for the ui
var InputEditor gocui.Editor = gocui.EditorFunc(simpleEditor)

func simpleEditor(v *gocui.View, key gocui.Key, ch rune, mod gocui.Modifier) {
	switch {
	case ch != 0 && mod == 0:
		v.EditWrite(ch)
	case key == gocui.KeySpace:
		v.EditWrite(' ')
	case key == gocui.KeyBackspace || key == gocui.KeyBackspace2:
		v.EditDelete(true)
	case key == gocui.KeyEnter:
		if line := v.ViewBuffer(); len(line) > 0 {
			handleLine(line, v)
			v.Clear()
			v.SetCursor(1, 0)
			v.SetOrigin(1, 0)
		}
	}
}

// layout - creates the UI layout
func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	// the log output view
	if v, err := g.SetView("logs", 0, 0, maxX-1, maxY-4); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		fmt.Fprintln(v, "Log view")
	}

	// the input box view
	if v, err := g.SetView("input", 0, maxY-3, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		_, err := g.SetCurrentView("input")

		if err != nil {
			return err
		}

		v.Editor = InputEditor
		v.FgColor = gocui.Attribute(15 + 1)
		v.BgColor = gocui.ColorDefault
		v.Autoscroll = false
		v.Editable = true
		v.Wrap = false
		v.Frame = true
	}

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func handleLine(line string, v *gocui.View) error {
	vws, err := mainGui.View("input")
	if err != nil {
		log.Panicln(err)
	}

	fmt.Println(vws)

	return nil
}
