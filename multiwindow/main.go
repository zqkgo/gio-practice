package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/unit"
)

func main() {
	closed := make(chan struct{}, 100)
	num := 50
	for i := 0; i < num; i++ {
		go newWindow(closed, i)
	}
	go func() {
		time.Sleep(5 * time.Second)
		for i := 0; i < num; i++ {
			closed <- struct{}{}
			time.Sleep(100 * time.Millisecond)
		}
		os.Exit(0)
	}()
	app.Main()
}

func newWindow(closed chan struct{}, idx int) {
	w := app.NewWindow(
		app.Size(unit.Dp(400), unit.Dp(300)),
		app.Title(fmt.Sprintf("Window %d", idx)),
	)
	for {
		select {
		case <-closed:
			log.Println("send destroy event")
			w.Close()
		case e := <-w.Events():
			if _, ok := e.(system.DestroyEvent); ok {
				log.Println("destroy event received, quit event loop")
				return
			}
		}
	}
}
