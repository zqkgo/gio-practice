package main

import (
	"image/color"
	"os"

	"gioui.org/app"
	"gioui.org/io/key"
	"gioui.org/io/system"
	"gioui.org/op"
	"gioui.org/op/paint"
)

var isFocused bool

func main() {
	go func() {
		w := app.NewWindow()
		ops := &op.Ops{}
		k := &isFocused
		for {
			e := <-w.Events()
			switch e := e.(type) {
			case system.DestroyEvent:
				os.Exit(0)
			case system.FrameEvent:
				// 声明handler
				key.InputOp{Tag: k}.Add(ops)
				key.FocusOp{Tag: k}.Add(ops)
				for _, ev := range e.Queue.Events(k) {
					if x, ok := ev.(key.FocusEvent); ok {
						if x.Focus {
							isFocused = true
						} else {
							isFocused = false
						}
					}
				}
				red := color.NRGBA{R: 0xE6, G: 0x05, B: 0x2A, A: 0xFF}
				green := color.NRGBA{R: 0x05, G: 0xE6, B: 0x05, A: 0xFF}
				// 设置刷子颜色
				if isFocused {
					paint.ColorOp{Color: red}.Add(ops)
				} else {
					paint.ColorOp{Color: green}.Add(ops)
				}
				paint.PaintOp{}.Add(ops)
				e.Frame(ops)
			}
		}
	}()
	app.Main()
}
