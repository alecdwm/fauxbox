package game

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/veandco/go-sdl2/sdl"
	"go.owls.io/fauxbox/engine"
)

////////////////////////////////////////////////////////////////////////////////
// SYSTEM //
///////////

type World struct {
	Camera Camera

	Objects []Object
}

var GameWorld World

func init() {
	engine.Register(&GameWorld)
}

////////////////////////////////////////////////////////////////////////////////
// STATES //
///////////

var WorldStates map[engine.State]bool = map[engine.State]bool{INGAME: true}

func (w *World) States() map[engine.State]bool {
	return WorldStates
}

////////////////////////////////////////////////////////////////////////////////
// CALLBACKS //
//////////////

func (w *World) StateEntry() {
	w.Objects = append(w.Objects, &Senti{color: sdl.Color{255, 255, 255, 255}, speed: 150, player: true})
	engine.Register(w.Objects[len(w.Objects)-1])
	w.Objects[len(w.Objects)-1].(*Senti).StateEntry()
}

func (w *World) Update(dt float64) {

}

func (w *World) Render(dt float64) {

}

////////////////////////////////////////////////////////////////////////////////
// FUNCTIONS //
//////////////

// Convert world-space coordinate into screen-space coordinate
func (w *World) Pos(worldPos mgl64.Vec2) (renderPos mgl64.Vec2) {
	renderPos = worldPos.
		Sub(w.Camera.Pos).
		Add(mgl64.Vec2{float64(engine.Width / 2), float64(engine.Height / 2)})
	return renderPos
}

// Convert screen-space coordinates into world-space coordinates
func (w *World) Ray(renderPos mgl64.Vec2) (worldPos mgl64.Vec2) {
	worldPos = renderPos.
		Add(w.Camera.Pos).
		Sub(mgl64.Vec2{float64(engine.Width / 2), float64(engine.Height / 2)})
	return renderPos
}
