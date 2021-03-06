package game

import (
	"github.com/veandco/go-sdl2/sdl"
	"go.owls.io/fauxbox/engine"
)

////////////////////////////////////////////////////////////////////////////////
// SYSTEM //
///////////

type GlobalKeybinds struct{}

func init() {
	engine.Register(&GlobalKeybinds{})
}

////////////////////////////////////////////////////////////////////////////////
// CALLBACKS //
//////////////

func (gk GlobalKeybinds) ProcessEvent(event sdl.Event) {
	switch e := event.(type) {
	case *sdl.KeyDownEvent:
		switch e.Keysym.Sym {
		case sdl.K_q:
			engine.EndGame()
		}
	}
}
