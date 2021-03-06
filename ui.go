package ghosts

import (
	"fmt"
	g "github.com/AllenDang/giu"
)

func onClickMe() {
	fmt.Println("Hello world!")
}

func onImSoCute() {
	fmt.Println("Im sooooooo cute!!")
}

func loop() {
	g.SingleWindow("hello world", g.Layout{
		g.Label("Hello world from giu"),
		g.Line(
			g.Button("Click Me", onClickMe),
			g.Button("I'm so cute", onImSoCute)),
	})
}

func Draw() {
	wnd := g.NewMasterWindow("Hello world", 400, 200, true, nil)
	wnd.Main(loop)
}
