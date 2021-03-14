package main

import (
	"image/color"
	"math/rand"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
)

func main() {
	go func() {
		w := app.NewWindow(
			app.Size(unit.Dp(400), unit.Dp(300)),
		)
		t := time.NewTicker(500 * time.Millisecond)
		var ops op.Ops
		for {
			select {
			case <-t.C:
				// change color
				w.Invalidate()
			case e := <-w.Events():
				switch e := e.(type) {
				case system.DestroyEvent:
					os.Exit(0)
				case system.FrameEvent:
					// set color
					paint.ColorOp{Color: pickColor()}.Add(&ops)
					// paint window
					paint.PaintOp{}.Add(&ops)
					e.Frame(&ops)
				}
			}
		}
	}()
	app.Main()
}

func pickColor() color.NRGBA {
	colors := []color.NRGBA{
		{R: 0xFF, A: 0xFF},
		{R: 0x0A, G: 0xFF, B: 0x50, A: 0xFF},
		{R: 0x1F, G: 0x22, B: 0xFF, A: 0xFF},
		{R: 0xFF, G: 0x66, B: 0xCE, A: 0xFF},
		{R: 0x3D, G: 0xD4, B: 0xCE, A: 0xFF},
		{R: 0xFD, G: 0x0C, B: 0xE1, A: 0xFF},
	}
	return colors[rand.Intn(6)]
}
