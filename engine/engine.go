package engine

import (
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/dradtke/go-allegro/allegro"
	"github.com/dradtke/go-allegro/allegro/audio"
	"github.com/dradtke/go-allegro/allegro/dialog"
	"github.com/dradtke/go-allegro/allegro/font"
	"github.com/dradtke/go-allegro/allegro/font/ttf"
	"github.com/dradtke/go-allegro/allegro/image"
	"github.com/dradtke/go-allegro/allegro/primitives"
)

var (
	// Public
	FPS    float64
	Width  int
	Height int

	// Private
	display *allegro.Display
	running bool = true
)

////////////////////////////////////////////////////////////////////////////////
// PUBLIC (BEGIN/RESIZE/END) ///////////////////////////////////////////////////
///////////////////////////////
func Fauxbox() {
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
	var err error
	display, err = allegro.CreateDisplay(800, 600)
	if err != nil {
		logrus.WithError(err).Error("Creating allegro display")
	}
	defer display.Destroy()

	Width = display.Width()
	Height = display.Height()

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
				logrus.WithError(err).Error("Retrieving next event from not-empty queue")
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

func Resized() {
	Width = display.Width()
	Height = display.Height()
}

func EndGame() {
	running = false
}
