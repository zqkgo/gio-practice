package main

import (
	"image/color"
	"math/rand"
	"os"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/key"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type UI struct {
	th *material.Theme

	clk *widget.Clickable
	// must be outside event loop
	idx    int
	colors []color.NRGBA
}

func NewUI() *UI {
	return &UI{
		th:  material.NewTheme(gofont.Collection()),
		clk: &widget.Clickable{},
		colors: []color.NRGBA{
			{R: 0xF1, G: 0x37, B: 0x5c, A: 0xFF},
			{R: 0xC2, G: 0x37, B: 0xF1, A: 0xFF},
			{R: 0x5E, G: 0xC9, B: 0xC4, A: 0xFF},
			{R: 0xA8, G: 0x80, B: 0x99, A: 0xFF},
			{R: 0x16, G: 0xAC, B: 0x63, A: 0xFF},
			{R: 0x70, G: 0xAC, B: 0x16, A: 0xFF},
			{R: 0x70, G: 0xAC, B: 0x16, A: 0xFF},
			{R: 0xAC, G: 0x52, B: 0x16, A: 0xFF},
			{R: 0xF0, G: 0x60, B: 0x00, A: 0xFF},
			{R: 0x00, G: 0x12, B: 0xd6, A: 0xFF},
		},
	}
}

func (ui *UI) Loop(w *app.Window) {
	ops := new(op.Ops)
	for e := range w.Events() {
		switch e := e.(type) {
		case system.DestroyEvent:
			os.Exit(0)
		case system.FrameEvent:
			gtx := layout.NewContext(ops, e)
			ui.Layout(gtx)
			e.Frame(gtx.Ops)
		case key.Event:
			if e.Name == key.NameEscape {
				w.Close()
			}
		}
	}
}

func (ui *UI) Layout(gtx layout.Context) layout.Dimensions {
	return layout.UniformInset(unit.Dp(20)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{}.Layout(gtx,
			layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
				btn := material.Button(ui.th, ui.clk, "Hello")
				for range ui.clk.Clicks() {
					// pick color
					for i := 0; i < 10; i++ {
						idx := rand.Intn(len(ui.colors))
						if idx != ui.idx {
							println(idx)
							ui.idx = idx
							break
						}
					}
				}
				btn.Background = ui.colors[ui.idx]
				return btn.Layout(gtx)
			}),
		)
	})
}

func main() {
	ui := NewUI()
	w := app.NewWindow(
		app.Size(unit.Dp(250), unit.Dp(120)),
		app.Title("Change Button Color"),
	)
	go ui.Loop(w)
	app.Main()
}
