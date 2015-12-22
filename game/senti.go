package game

import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/veandco/go-sdl2/sdl"
	"go.owls.io/fauxbox/engine"
)

////////////////////////////////////////////////////////////////////////////////
// ENT //
////////

type Senti struct {
	pos       mgl64.Vec2
	vel       mgl64.Vec2
	target    mgl64.Vec2
	direction Direction

	color sdl.Color
	speed float64

	player    bool
	networked bool

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

func (s *Senti) Update(dt float64) {
	// Use dir for facing the sprite
	// fmt.Printf("%v\n", s.target.Sub(s.pos).Normalize())
	// dir := math.Acos(s.target.Sub(s.pos).Normalize().Dot(mgl64.Vec2{0, 1}))
	// fmt.Println(dir)

	if s.vel.X() != 0 || s.vel.Y() != 0 {
		s.pos = s.pos.Add(s.vel.Mul(dt))
	}

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

func (s *Senti) IsPlayer() bool {
	return s.player
}

func (s *Senti) IsNetworked() bool {
	return s.networked
}

func (s *Senti) SetPosition(newPos mgl64.Vec2) {
	s.pos = newPos
}

func (s *Senti) SetVelocity(newVel mgl64.Vec2) {
	s.vel = newVel
}

func (s *Senti) SetTarget(newTarget mgl64.Vec2) {
	s.target = newTarget
}

func (s *Senti) Speed() float64 {
	return s.speed
}

func (s *Senti) SetSpeed(speed float64) {
	s.speed = speed
}

func (s *Senti) SetColor(color sdl.Color) {
	s.color = color
}
