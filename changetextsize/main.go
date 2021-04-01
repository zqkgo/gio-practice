package main

import (
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

func main() {
	w := app.NewWindow(
		app.Title("Change Text Size"),
		app.Size(unit.Dp(600), unit.Dp(300)),
	)
	go NewUI().Run(w)
	app.Main()
}

type UI struct {
	th   *material.Theme
	size *widget.Float
	txt  material.LabelStyle
}

func NewUI() *UI {
	ui := &UI{
		th:   material.NewTheme(gofont.Collection()),
		size: &widget.Float{Value: 14},
	}
	ui.txt = material.H6(ui.th, "Hello, world")
	return ui
}

func (ui *UI) Run(w *app.Window) {
	ops := new(op.Ops)
	for e := range w.Events() {
		switch e := e.(type) {
		case system.DestroyEvent:
			os.Exit(0)
		case system.FrameEvent:
			gtx := layout.NewContext(ops, e)
			ui.Layout(gtx)
			e.Frame(gtx.Ops)
		}
	}
}

func (ui *UI) Layout(gtx layout.Context) layout.Dimensions {
	if ui.size.Changed() {
		ui.txt.TextSize = unit.Dp(ui.size.Value)
	}
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return ui.txt.Layout(gtx)
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			// 滑动条位于底部
			return layout.Stack{Alignment: layout.S}.Layout(gtx,
				layout.Stacked(func(gtx layout.Context) layout.Dimensions {
					// 填满窗口
					return layout.Flex{}.Layout(gtx, layout.Flexed(1, material.Slider(ui.th, ui.size, 0, 100).Layout))
				}),
			)
		}),
	)
}
