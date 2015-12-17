package game

import (
	"math"

	"github.com/dradtke/go-allegro/allegro"
	"github.com/dradtke/go-allegro/allegro/primitives"
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

func (p *Player) ProcessEvent(event interface{}) {
	switch e := event.(type) {
	case allegro.KeyDownEvent:
		switch e.KeyCode() {
		case allegro.KEY_W, allegro.KEY_UP:
			p.UP = true
		case allegro.KEY_A, allegro.KEY_LEFT:
			p.LEFT = true
		case allegro.KEY_S, allegro.KEY_DOWN:
			p.DOWN = true
		case allegro.KEY_D, allegro.KEY_RIGHT:
			p.RIGHT = true
		}

	case allegro.KeyUpEvent:
		switch e.KeyCode() {
		case allegro.KEY_W, allegro.KEY_UP:
			p.UP = false
		case allegro.KEY_A, allegro.KEY_LEFT:
			p.LEFT = false
		case allegro.KEY_S, allegro.KEY_DOWN:
			p.DOWN = false
		case allegro.KEY_D, allegro.KEY_RIGHT:
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
}

func (p *Player) Draw(dt float64) {
	primitives.DrawCircle(primitives.Point{float32(p.x), float32(p.y)}, 20.0, allegro.MapRGB(255, 255, 255), 2.0)
}
