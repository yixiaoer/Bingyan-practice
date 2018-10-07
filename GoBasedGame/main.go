package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/jroimartin/gocui"

	"project/GoBasedGame/character"
)

var (
	viewArr = []string{"v1", "v2", "v3"}
	active  = 0
)

func SetViewOnTop(g *gocui.Gui, name string) (*gocui.View, error) {
	if _, err := g.SetCurrentView(name); err != nil {
		return nil, err
	}
	return g.SetViewOnTop(name)
}

func Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	var c character.Character
	c = character.Init()
	var direct string

	if v, err := g.SetView("v1", 0, 2*maxY/3, maxX-3, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Instructions"
		v.Editable = true
		v.Wrap = true

		fmt.Fscanln(v, &direct)

		if _, err = SetViewOnTop(g, "v1"); err != nil {
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
		fmt.Fprintln(v, "[Sound ||  I am so glad that you are here]")
		fmt.Fprintln(v, "[Sound ||  You're invited here for some special reasons]")
		fmt.Fprintln(v, "[Sound ||  Wanna know why?]")
		fmt.Fprintln(v, "[Sound ||  Just find it]")
		fmt.Fprintln(v, "[SILENCE]")
		fmt.Fprintln(v, "Go on your journey and have fun!")
		fmt.Fprintln(v, direct)
	}

	if v, err := g.SetView("v3", 0, 0, maxX/4-1, 2*maxY/3-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "UserInfo"
		v.Editable = true

		fmt.Fprintln(v, "x:", strconv.Itoa(c.Xp))
		fmt.Fprintln(v, "y:", strconv.Itoa(c.Yp))
		fmt.Fprintln(v, "point:", strconv.Itoa(c.Point))
		fmt.Fprintln(v, "==============")
	}
	return nil
}

func Quit(g *gocui.Gui, v *gocui.View) error {
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

	g.SetManagerFunc(Layout)

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, Quit); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}
