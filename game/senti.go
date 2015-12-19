package game

import (
	"fmt"
	"math"

	"github.com/go-gl/mathgl/mgl64"
	"github.com/veandco/go-sdl2/sdl"
	"go.owls.io/fauxbox/engine"
)

////////////////////////////////////////////////////////////////////////////////
// ENT //
////////

type Senti struct {
	pos       mgl64.Vec2
	target    mgl64.Vec2
	direction Direction

	color  sdl.Color
	speed  float64
	player bool

	posRect    *sdl.Rect
	targetRect *sdl.Rect
}

func (s *Senti) ID() uint {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// CALLBACKS //
//////////////

func (s *Senti) StateEntry() {
	s.posRect = &sdl.Rect{0, 0, 20, 20}
	s.targetRect = &sdl.Rect{0, 0, 6, 6}
}

func (s *Senti) ProcessEvent(event sdl.Event) {
	switch e := event.(type) {
	case *sdl.KeyDownEvent:
		switch e.Keysym.Sym {
		case sdl.K_w, sdl.K_UP:
			s.MoveTarget(0, -100)

		case sdl.K_a, sdl.K_LEFT:
			s.MoveTarget(-100, 0)

		case sdl.K_s, sdl.K_DOWN:
			s.MoveTarget(0, 100)

		case sdl.K_d, sdl.K_RIGHT:
			s.MoveTarget(100, 0)
		}
	}
}

func (s *Senti) Update(dt float64) {
	// Use dir for facing the sprite
	dir := math.Acos(s.pos.Normalize().Dot(s.target.Normalize()))
	fmt.Println(dir)

	s.pos = s.pos.Add(s.target.Sub(s.pos).Mul( /*s.speed * */ dt))

	fmt.Println(s.pos.X(), s.pos.Y())

	s.posRect.X = int32(GameWorld.Pos(s.pos).X()) - s.posRect.W/2
	s.posRect.Y = int32(GameWorld.Pos(s.pos).Y()) - s.posRect.H/2

	s.targetRect.X = int32(GameWorld.Pos(s.target).X()) - s.targetRect.W/2
	s.targetRect.Y = int32(GameWorld.Pos(s.target).Y()) - s.targetRect.H/2

	GameWorld.Camera.SetTarget(s.pos)
}

func (s *Senti) Draw(dt float64) {
	engine.Renderer.SetDrawColor(255, 255, 255, 255)
	engine.Renderer.DrawRect(s.posRect)

	engine.Renderer.SetDrawColor(255, 0, 0, 255)
	engine.Renderer.DrawRect(s.targetRect)
}

////////////////////////////////////////////////////////////////////////////////
// FUNCTIONS //
//////////////

func (s *Senti) MoveTarget(targetX, targetY float64) {
	s.target = s.target.Add(mgl64.Vec2{targetX, targetY})
}

func (s *Senti) SetTarget(targetX, targetY float64) {
	s.target = mgl64.Vec2{targetX, targetY}
}

func (s *Senti) SetSpeed(speed float64) {
	s.speed = speed
}

func (s *Senti) SetColor(color sdl.Color) {
	s.color = color
}
