package main

import (
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

func main() {
	go func() {
		w := app.NewWindow(
			app.Size(unit.Dp(360), unit.Dp(360)),
		)
		ops := new(op.Ops)
		th := material.NewTheme(gofont.Collection())
		var ed1, ed2 widget.Editor
		for e := range w.Events() {
			switch e := e.(type) {
			case system.FrameEvent:
				gtx := layout.NewContext(ops, e)
				layout.Flex{}.Layout(gtx,
					layout.Rigid(layout.Spacer{Width: unit.Dp(10)}.Layout),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return material.Editor(th, &ed1, "Hello").Layout(gtx)
					}),
					layout.Rigid(layout.Spacer{Width: unit.Dp(10)}.Layout),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						ed2.SetText(ed1.Text())
						return material.Editor(th, &ed2, "World").Layout(gtx)
					}),
				)
				e.Frame(gtx.Ops)
			case system.DestroyEvent:
				println("program exit")
				os.Exit(0)
			case key.Event:
				println("key name: ", e.Name)
			}
		}
	}()
	app.Main()
}
