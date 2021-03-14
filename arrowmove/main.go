package main

import (
	"image"
	"image/color"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/io/key"
	"gioui.org/io/system"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

var x, y float32

func main() {
	go func() {
		w := app.NewWindow()
		ops := new(op.Ops)
		for e := range w.Events() {
			switch e := e.(type) {
			case system.FrameEvent:
				ops.Reset()
				drawSqure(ops)
				e.Frame(ops)
			case key.Event:
				switch e.Name {
				case key.NameDownArrow:
					y += 10
				case key.NameUpArrow:
					y -= 10
				case key.NameLeftArrow:
					x -= 10
				case key.NameRightArrow:
					x += 10
				}
				w.Invalidate()
			}
		}
	}()
	app.Main()
}

func drawSqure(ops *op.Ops) {
	defer op.Save(ops).Load()
	op.Offset(f32.Pt(x, y)).Add(ops)
	clip.Rect{Min: image.Pt(0, 0), Max: image.Pt(100, 100)}.Add(ops)
	red := color.NRGBA{R: 0xFB, G: 0x1B, B: 0x13, A: 0xFF}
	paint.ColorOp{Color: red}.Add(ops)
	paint.PaintOp{}.Add(ops)
}
