package main

import (
	"fmt"

	"github.com/mibesr/ghosts"
)

func main() {
	fmt.Println(ghosts.ConfigDir())
	ghosts.Draw()
}
