// Fauxbox, A 2D sandbox game
package main

import (
	"go.owls.io/fauxbox/engine"
	"go.owls.io/fauxbox/game"
)

func main() {
	game.Attach()                 // Attach game files
	engine.Fauxbox(game.MAINMENU) // Run Fauxbox
}
