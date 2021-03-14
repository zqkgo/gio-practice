package main

import (
	"image"
	"image/color"
	"os"

	"gioui.org/app"
	"gioui.org/f32"
	"gioui.org/io/system"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
)

var end float32

func main() {
	go func() {
		w := app.NewWindow()
		ops := &op.Ops{}
		for {
			e := <-w.Events()
			switch e := e.(type) {
			case system.DestroyEvent:
				os.Exit(0)
			case system.FrameEvent:
				ops.Reset()
				move(ops)
				e.Frame(ops)
			}
		}
	}()
	app.Main()
}

func move(ops *op.Ops) {
	if end < 500 {
		end++
		op.InvalidateOp{}.Add(ops)
	}
	moveRight(ops)
	moveDown(ops)
}

func moveRight(ops *op.Ops) {
	defer op.Save(ops).Load()
	op.Offset(f32.Pt(end, 0)).Add(ops)
	clip.Rect{Min: image.Point{0, 0}, Max: image.Point{100, 100}}.Add(ops)
	cl := color.NRGBA{R: 0x01, G: 0x33, B: 0x16, A: 0xFF}
	paint.ColorOp{Color: cl}.Add(ops)
	paint.PaintOp{}.Add(ops)
}

func moveDown(ops *op.Ops) {
	defer op.Save(ops).Load()
	op.Offset(f32.Pt(0, end)).Add(ops)
	clip.Rect{Min: image.Point{0, 0}, Max: image.Point{100, 100}}.Add(ops)
	cl := color.NRGBA{R: 0xAE, G: 0x04, B: 0x84, A: 0xFF}
	paint.ColorOp{Color: cl}.Add(ops)
	paint.PaintOp{}.Add(ops)
}
