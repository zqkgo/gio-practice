package main

import (
	"fmt"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

var btnNum = 100

type UI struct {
	th *material.Theme
	ls *layout.List
	cs []*widget.Clickable
}

func New() *UI {
	ui := &UI{
		th: material.NewTheme(gofont.Collection()),
		ls: &layout.List{Axis: layout.Vertical},
	}
	for i := 0; i < btnNum; i++ {
		ui.cs = append(ui.cs, &widget.Clickable{})
	}
	return ui
}

func (ui *UI) Layout(gtx layout.Context) layout.Dimensions {
	return layout.UniformInset(unit.Dp(10)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return ui.ls.Layout(gtx, btnNum, func(gtx layout.Context, i int) layout.Dimensions {
			return layout.UniformInset(unit.Dp(1)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return material.Button(ui.th, ui.cs[i], fmt.Sprintf("btn-%d", i)).Layout(gtx)
			})
		})
	})
}

func (ui *UI) Loop(w *app.Window) {
	ops := new(op.Ops)
	for e := range w.Events() {
		switch e := e.(type) {
		case system.FrameEvent:
			gtx := layout.NewContext(ops, e)
			ui.Layout(gtx)
			e.Frame(gtx.Ops)
		case system.DestroyEvent:
			os.Exit(0)
		}
	}
}

func main() {
	ui := New()
	w := app.NewWindow(
		app.Title("List Buttons"),
		app.Size(unit.Dp(350), unit.Dp(500)),
	)
	go ui.Loop(w)
	app.Main()
}
