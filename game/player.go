package game

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
	"go.owls.io/fauxbox/engine"
)

////////////////////////////////////////////////////////////////////////////////
// SYSTEM //////////////////////////////////////////////////////////////////////
////////////

type Player struct {
	x     float64
	y     float64
	LEFT  bool
	RIGHT bool
	UP    bool
	DOWN  bool
	speed float64
}

func init() {
	engine.Register(&Player{speed: 150})
}

////////////////////////////////////////////////////////////////////////////////
// CALLBACKS ///////////////////////////////////////////////////////////////////
///////////////

func (p *Player) ProcessEvent(event sdl.Event) {
	switch e := event.(type) {
	case *sdl.KeyDownEvent:
		switch e.Keysym.Sym {
		case sdl.K_w, sdl.K_UP:
			p.UP = true
		case sdl.K_a, sdl.K_LEFT:
			p.LEFT = true
		case sdl.K_s, sdl.K_DOWN:
			p.DOWN = true
		case sdl.K_d, sdl.K_RIGHT:
			p.RIGHT = true
		}

	case *sdl.KeyUpEvent:
		switch e.Keysym.Sym {
		case sdl.K_w, sdl.K_UP:
			p.UP = false
		case sdl.K_a, sdl.K_LEFT:
			p.LEFT = false
		case sdl.K_s, sdl.K_DOWN:
			p.DOWN = false
		case sdl.K_d, sdl.K_RIGHT:
			p.RIGHT = false
		}
	}
}

func (p *Player) Update(dt float64) {
	if p.UP && !p.DOWN {
		if p.LEFT && !p.RIGHT {
			p.y -= p.speed * math.Cos(45) * dt
			p.x -= p.speed * math.Cos(45) * dt
		} else if p.RIGHT && !p.LEFT {
			p.y -= p.speed * math.Cos(45) * dt
			p.x += p.speed * math.Cos(45) * dt
		} else {
			p.y -= p.speed * dt
		}
	} else if p.DOWN && !p.UP {
		if p.LEFT && !p.RIGHT {
			p.y += p.speed * math.Cos(45) * dt
			p.x -= p.speed * math.Cos(45) * dt
		} else if p.RIGHT && !p.LEFT {
			p.y += p.speed * math.Cos(45) * dt
			p.x += p.speed * math.Cos(45) * dt
		} else {
			p.y += p.speed * dt
		}
	} else if p.LEFT && !p.RIGHT {
		p.x -= p.speed * dt
	} else if p.RIGHT && !p.LEFT {
		p.x += p.speed * dt
	}

	World.TargetX = p.x
	World.TargetY = p.y
}

func (p *Player) Draw(dt float64) {
	engine.Renderer.SetDrawColor(255, 255, 255, 255)
	engine.Renderer.DrawRect(&sdl.Rect{int32(World.X(p.x)) - 10, int32(World.Y(p.y)) - 10, 20, 20})
	engine.Renderer.DrawPoint(int(World.X(p.x)), int(World.Y(p.y)))
}
