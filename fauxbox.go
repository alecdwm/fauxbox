// Fauxbox, A 2D sandbox game
package main

import (
	"github.com/dradtke/go-allegro/allegro"
	"go.owls.io/fauxbox/engine"
	"go.owls.io/fauxbox/game"
)

func main() {
	game.Attach() // Attach game files

	allegro.Run(engine.Fauxbox) // Run Fauxbox
}
