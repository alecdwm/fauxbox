// Fauxbox, A 2D sandbox game
package main

import (
	"fmt"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/dradtke/go-allegro/allegro"
	"github.com/dradtke/go-allegro/allegro/audio"
	"github.com/dradtke/go-allegro/allegro/dialog"
	"github.com/dradtke/go-allegro/allegro/font"
	"github.com/dradtke/go-allegro/allegro/font/ttf"
	"github.com/dradtke/go-allegro/allegro/image"
	"github.com/dradtke/go-allegro/allegro/primitives"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/kardianos/osext"
)

type Bullet struct {
	active    bool
	livedtime float64
	lifetime  float64

	pos mgl64.Vec2
	vel mgl64.Vec2
}

type Bullets [40]Bullet

func (bullets *Bullets) fire(pos, vel mgl64.Vec2, lifetime float64) {
	for i, bullet := range *bullets {
		if !bullet.active {
			bullets[i] = Bullet{
				active:    true,
				livedtime: 0,
				lifetime:  lifetime,

				pos: pos,
				vel: vel,
			}
			logrus.WithFields(logrus.Fields{
				"i":         i,
				"active":    bullets[i].active,
				"livedtime": bullets[i].livedtime,
				"lifetime":  bullets[i].lifetime,
				"pos":       bullets[i].pos,
				"vel":       bullets[i].vel,
			}).Info("set active")
			break
		}
	}
}

func (bullets *Bullets) update(dt float64) {
	for i := range *bullets {
		if !bullets[i].active {
			continue
		}

		bullets[i].livedtime += dt
		if bullets[i].livedtime > bullets[i].lifetime {
			bullets[i].active = false
			continue
		}

		if bullets[i].pos.X() < 0 {
			bullets[i].pos = mgl64.Vec2{float64(width), bullets[i].pos.Y()}
		}
		if bullets[i].pos.Y() < 0 {
			bullets[i].pos = mgl64.Vec2{bullets[i].pos.X(), float64(height)}
		}
		if bullets[i].pos.X() > float64(width) {
			bullets[i].pos = mgl64.Vec2{0.0, bullets[i].pos.Y()}
		}
		if bullets[i].pos.Y() > float64(height) {
			bullets[i].pos = mgl64.Vec2{bullets[i].pos.X(), 0.0}
		}

		bullets[i].pos = bullets[i].pos.Add(bullets[i].vel.Mul(dt))
	}
}

func (bullets *Bullets) draw(dt float64) {
	for _, bullet := range *bullets {
		if !bullet.active {
			continue
		}

		primitives.DrawFilledCircle(
			primitives.Point{float32(bullet.pos.X()), float32(bullet.pos.Y())},
			3.0,
			allegro.MapRGB(255, 0, 0),
		)
	}
}

type MouseCannon struct {
	armed            bool
	NextBulletPos    mgl64.Vec2
	NextBulletAimPos mgl64.Vec2
}

func (mc *MouseCannon) isArmed() bool {
	return mc.armed
}

func (mc *MouseCannon) arm(mousePos mgl64.Vec2) {
	mc.NextBulletPos = mousePos
	mc.NextBulletAimPos = mousePos
	mc.armed = true
}

func (mc *MouseCannon) unarm() {
	mc.armed = false
}

func (mc *MouseCannon) aim(mousePos mgl64.Vec2) {
	if mc.armed {
		mc.NextBulletAimPos = mousePos
	}
}

func (mc *MouseCannon) fire(bulletPool *Bullets, mousePos mgl64.Vec2) {
	if mc.armed {
		bulletPool.fire(mc.NextBulletPos, mc.NextBulletAimPos.Sub(mc.NextBulletPos), 2)
		mc.armed = false
	}
}

func (mc *MouseCannon) draw(dt float64) {
	if mc.armed {
		primitives.DrawLine(
			primitives.Point{float32(mc.NextBulletPos.X()), float32(mc.NextBulletPos.Y())},
			primitives.Point{float32(mc.NextBulletAimPos.X()), float32(mc.NextBulletAimPos.Y())},
			allegro.MapRGB(255, 255, 255),
			1.0,
		)
	}
}

var (
	running bool = true

	defaultFont *font.Font
	fauxboximg  *allegro.Bitmap // this stuff will soon be split into objects not globals

	x     float64
	y     float64
	LEFT  bool
	RIGHT bool
	UP    bool
	DOWN  bool
	fps   float64

	width  int
	height int

	worldBullets Bullets
	mouseCannon  MouseCannon
)

func load() {
	execPath, err := osext.ExecutableFolder()
	if err != nil {
		logrus.WithError(err).Error("Getting executable's path")
	}
	resPath := execPath + "/resources"

	// FONTS
	defaultFont, err = font.LoadFont(resPath+"/Neuropol.ttf", 18, 0)
	if err != nil {
		logrus.WithError(err).Error("Loading font")
	}

	// IMAGES
	fauxboximg, err = allegro.LoadBitmap(resPath + "/fauxbox.tga")
	if err != nil {
		logrus.WithError(err).Error("Loading bitmap")
	}
}

func processEvent(event interface{}) {
	switch e := event.(type) {
	case allegro.DisplayCloseEvent:
		endGame()

	case allegro.KeyDownEvent:
		switch e.KeyCode() {
		case allegro.KEY_Q:
			endGame()

		case allegro.KEY_W, allegro.KEY_UP:
			UP = true
		case allegro.KEY_A, allegro.KEY_LEFT:
			LEFT = true
		case allegro.KEY_S, allegro.KEY_DOWN:
			DOWN = true
		case allegro.KEY_D, allegro.KEY_RIGHT:
			RIGHT = true
		}

	case allegro.KeyUpEvent:
		switch e.KeyCode() {
		case allegro.KEY_W, allegro.KEY_UP:
			UP = false
		case allegro.KEY_A, allegro.KEY_LEFT:
			LEFT = false
		case allegro.KEY_S, allegro.KEY_DOWN:
			DOWN = false
		case allegro.KEY_D, allegro.KEY_RIGHT:
			RIGHT = false
		}

	case allegro.MouseButtonDownEvent:
		if e.Button() == 1 {
			mouseCannon.arm(mgl64.Vec2{float64(e.X()), float64(e.Y())})
		} else if e.Button() == 2 {
			mouseCannon.unarm()
		}

	case allegro.MouseAxesEvent:
		if mouseCannon.isArmed() {
			mouseCannon.aim(mgl64.Vec2{float64(e.X()), float64(e.Y())})
		}

	case allegro.MouseButtonUpEvent:
		if e.Button() == 1 {
			mouseCannon.fire(&worldBullets, mgl64.Vec2{float64(e.X()), float64(e.Y())})
		}

	default:
		// Handle other events here.
	}
}

func update(dt float64) {
	fps = dt
	if UP && !DOWN {
		y -= 50 * dt
	}
	if LEFT && !RIGHT {
		x -= 50 * dt
	}
	if RIGHT && !LEFT {
		x += 50 * dt
	}
	if DOWN && !UP {
		y += 50 * dt
	}

	worldBullets.update(dt)
}

func draw(dt float64) {
	allegro.ClearToColor(allegro.MapRGB(0, 0, 0))

	fauxboximg.Draw(200, 200, allegro.FLIP_NONE)
	primitives.DrawCircle(primitives.Point{float32(x), float32(y)}, 20.0, allegro.MapRGB(255, 255, 255), 2.0)

	worldBullets.draw(dt)
	mouseCannon.draw(dt)

	font.DrawText(defaultFont, allegro.MapRGB(255, 255, 255), 20, 20, font.ALIGN_LEFT, "Hello World!")
	font.DrawText(defaultFont, allegro.MapRGB(255, 255, 255), 20, 60, font.ALIGN_LEFT, fmt.Sprintf("%.1f", 1.0/fps))

	allegro.FlipDisplay()
}

func endGame() {
	running = false
}

func main() {
	allegro.Run(game) // Run allegro
}

func game() {
	// Initialize allegro features
	allegro.InstallKeyboard() // keyboard input
	allegro.InstallMouse()    // mouse input
	allegro.InstallJoystick() // joystick input

	// Initialize allegro addons
	audio.Install()      // lets us run audio
	dialog.Install()     // lets us show native dialogs
	font.Install()       // lets us load fonts from popular file formats
	image.Install()      // lets us load bitmap images from popular file formats
	primitives.Install() // lets us use some 2d graphics primitives
	ttf.Install()        // lets us use true type fonts

	// LOAD RESOURCES
	load()

	// CREATE CONTEXT
	//   i.e. an allegro 'display' at 800px x 600px
	allegro.SetNewDisplayFlags(allegro.WINDOWED)
	display, err := allegro.CreateDisplay(800, 600)
	if err != nil {
		logrus.WithError(err).Error("Creating allegro display")
	}
	defer display.Destroy()

	width = display.Width()
	height = display.Height()

	// CONFIGURE CONTEXT
	display.SetWindowTitle("Fauxbox") // set a title

	// CREATE EVENT QUEUE
	eventQueue, err := allegro.CreateEventQueue()
	if err != nil {
		logrus.WithError(err).Error("Creating event queue")
	}
	defer eventQueue.Destroy()

	// REGISTER EVENT SOURCES
	checkEventSource := func(es *allegro.EventSource, err error) *allegro.EventSource {
		if err != nil {
			logrus.WithError(err).Error("Registering event source")
		}
		return es
	}
	eventQueue.RegisterEventSource(display.EventSource())
	eventQueue.RegisterEventSource(checkEventSource(allegro.KeyboardEventSource()))
	eventQueue.RegisterEventSource(checkEventSource(allegro.MouseEventSource()))
	eventQueue.RegisterEventSource(allegro.JoystickEventSource())

	// RENDER BLACK SCREEN
	allegro.ClearToColor(allegro.MapRGB(0, 0, 0))
	allegro.FlipDisplay()

	// PREPARE MAIN LOOP
	var event allegro.Event
	var thenUpdate,
		thenDraw,
		nowUpdate,
		nowDraw time.Time = time.Now(), time.Now(), time.Now(), time.Now()

	// RUN MAIN LOOP
	for running {
		// PROCESS EVENTS
		for !eventQueue.IsEmpty() {
			if e, err := eventQueue.GetNextEvent(&event); err == nil {
				processEvent(e)
			} else {
				logrus.WithError(err).Error("Event queue not empty, but error in retrieving event!")
			}
		}

		// UPDATE LOGIC
		nowUpdate = time.Now()
		update(nowUpdate.Sub(thenUpdate).Seconds())
		thenUpdate = nowUpdate

		// RENDER FRAME
		nowDraw = time.Now()
		draw(nowDraw.Sub(thenDraw).Seconds())
		thenDraw = nowDraw
	}
}
