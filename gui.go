package main

import (
	"fmt"
	"log"

	"github.com/jroimartin/gocui"
)

var (
	viewArr = []string{"v1", "v2", "v3"}
	active  = 0
)

func setCurrentViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	if _, err := g.SetCurrentView(name); err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)
}

func nextView(g *gocui.Gui, v *gocui.View) error {

	out, err := g.View("v1")
	if err != nil {
		return err
	}
	fmt.Fprintln(out) //==========

	return nil
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("v1", 0, 2*maxY/3, maxX-3, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Instructions"
		v.Editable = true
		v.Wrap = true

		if _, err = setCurrentViewOnTop(g, "v1"); err != nil {
			return err
		}
	}

	if v, err := g.SetView("v2", maxX/4+1, 0, maxX-3, 2*maxY/3-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "EnvironmentInfo"
		v.Wrap = true
		v.Autoscroll = true
	}

	if v, err := g.SetView("v3", 0, 0, maxX/4-1, 2*maxY/3-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "UserInfo"
		v.Editable = true
	}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.Highlight = true
	g.Cursor = true
	g.SelFgColor = gocui.ColorGreen

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, nextView); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
