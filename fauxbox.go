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
	"github.com/dradtke/go-allegro/allegro/image"
	"github.com/dradtke/go-allegro/allegro/primitives"
	"github.com/kardianos/osext"
)

var (
	running     bool = true
	defaultFont *font.Font
	alecdwm     *allegro.Bitmap // this stuff will soon be split into objects not globals
	x           float64
	y           float64
	LEFT        bool
	RIGHT       bool
	UP          bool
	DOWN        bool
	fps         float64
)

func load() {
	execPath, err := osext.ExecutableFolder()
	if err != nil {
		logrus.WithError(err).Error("Getting executable's path")
	}
	resPath := execPath + "/resources"

	// FONTS
	defaultFont, err = font.Builtin()
	if err != nil {
		logrus.WithError(err).Error("Loading font")
	}

	// IMAGES
	alecdwm, err = allegro.LoadBitmap(resPath + "/alecdwm.jpg")
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
}

func draw(dt float64) {
	allegro.ClearToColor(allegro.MapRGB(0, 0, 0))

	alecdwm.Draw(200, 200, allegro.FLIP_NONE)
	primitives.DrawCircle(primitives.Point{float32(x), float32(y)}, 20.0, allegro.MapRGB(255, 255, 255), 2.0)
	font.DrawText(defaultFont, allegro.MapRGB(255, 255, 255), 20, 20, font.ALIGN_LEFT, "Hello World!")
	font.DrawText(defaultFont, allegro.MapRGB(255, 255, 255), 20, 60, font.ALIGN_LEFT, fmt.Sprintf("%.6f", fps))

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
