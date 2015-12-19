package game

import (
	"github.com/veandco/go-sdl2/sdl"
	"go.owls.io/fauxbox/engine"
)

////////////////////////////////////////////////////////////////////////////////
// SYSTEM //
///////////

type MainMenu struct{}

func init() {
	engine.Register(&MainMenu{})
}

////////////////////////////////////////////////////////////////////////////////
// STATES //
///////////

var MainMenuStates map[engine.State]bool = map[engine.State]bool{MAINMENU: true}

func (mm *MainMenu) States() map[engine.State]bool {
	return MainMenuStates
}

////////////////////////////////////////////////////////////////////////////////
// CALLBACKS //
//////////////

func (mm *MainMenu) ProcessEvent(event sdl.Event) {
	switch e := event.(type) {
	case *sdl.KeyDownEvent:
		if e.Keysym.Sym == sdl.K_RETURN || e.Keysym.Sym == sdl.K_RETURN2 ||
			e.Keysym.Sym == sdl.K_SPACE {
			engine.States.Switch(INGAME)
		}
	}
}
