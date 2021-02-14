package main

import (
	"image/color"
	"os"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/io/system"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

func main() {
	go func() {
		w := app.NewWindow()
		var ops op.Ops
		for {
			e := <-w.Events()
			switch e := e.(type) {
			case system.DestroyEvent:
				os.Exit(0)
			case system.FrameEvent:
				// ops unchanged after calling 'first'
				green(&ops)
				red(&ops)
				blue(&ops)
				e.Frame(&ops)
			}
		}
	}()
	app.Main()
}

func green(ops *op.Ops) {
	const r = 10
	bounds := f32.Rect(0, 0, 800, 800)
	clip.RRect{Rect: bounds, SE: r, SW: r, NW: r, NE: r}.Add(ops)
	paint.ColorOp{Color: color.NRGBA{R: 0x4F, G: 0xA1, B: 0x26, A: 0xFF}}.Add(ops)
	paint.PaintOp{}.Add(ops)
}

func red(ops *op.Ops) {
	defer op.Save(ops).Load()
	const r = 10
	bounds := f32.Rect(0, 0, 400, 400)
	clip.RRect{Rect: bounds, SE: r, SW: r, NW: r, NE: r}.Add(ops)
	paint.ColorOp{Color: color.NRGBA{R: 0xF8, G: 0x0D, B: 0x0D, A: 0xFF}}.Add(ops)
	paint.PaintOp{}.Add(ops)
}

func blue(ops *op.Ops) {
	defer op.Save(ops).Load()
	const r = 10
	bounds := f32.Rect(0, 0, 200, 200)
	clip.RRect{Rect: bounds, SE: r, SW: r, NW: r, NE: r}.Add(ops)
	paint.ColorOp{Color: color.NRGBA{R: 0x30, G: 0x0D, B: 0xF8, A: 0xFF}}.Add(ops)
	paint.PaintOp{}.Add(ops)
}
