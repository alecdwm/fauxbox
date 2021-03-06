package engine

import (
	"runtime"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"
)

var (
	// Public
	FPS      float64
	Width    int
	Height   int
	Window   *sdl.Window
	Renderer *sdl.Renderer
	Keyboard []uint8
	Mouse    MouseState

	// Private
	running bool = true
)

////////////////////////////////////////////////////////////////////////////////
// PUBLIC (BEGIN/RESIZE/END) //
//////////////////////////////
func Fauxbox(enterState State) {
	// SDL2 is not designed to work across multiple threads
	runtime.LockOSThread()

	// Initialize SDL2
	err := sdl.Init(sdl.INIT_VIDEO | sdl.INIT_AUDIO)
	if err != nil {
		logrus.WithError(err).Error("Initializing SDL2")
	}
	defer sdl.Quit()

	// Initialize SDL2 TTF Library
	if err = ttf.Init(); err != nil {
		logrus.WithError(err).Error("Initializing SDL2_TTF")
	}

	// CREATE CONTEXT
	//   i.e. an sdl window + renderer at 800px x 600px
	Window, Renderer, err = sdl.CreateWindowAndRenderer(
		800, 600, sdl.WINDOW_RESIZABLE|sdl.WINDOW_INPUT_FOCUS)
	if err != nil {
		logrus.WithError(err).Error("Creating window and renderer")
	}
	defer Renderer.Destroy()
	defer Window.Destroy()

	// CONFIGURE CONTEXT
	Window.SetTitle("Fauxbox")
	Width, Height = Window.GetSize()

	// LOAD RESOURCES
	load()

	// RENDER BLACK SCREEN
	Renderer.SetDrawColor(0, 0, 0, 255)
	Renderer.Clear()
	firstRender()
	Renderer.Present()

	// SETUP ENTRY STATE
	States.Current = enterState
	stateEntry()
	Keyboard = sdl.GetKeyboardState()

	// PREPARE MAIN LOOP
	var event sdl.Event
	var thenUpdate,
		thenDraw,
		nowUpdate,
		nowDraw time.Time = time.Now(), time.Now(), time.Now(), time.Now()

	// RUN MAIN LOOP
	for running {
		// PROCESS EVENTS
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.KeyDownEvent, *sdl.KeyUpEvent:
				sdl.PumpEvents() // Update engine.Keyboard

			case *sdl.MouseButtonEvent, *sdl.MouseMotionEvent, *sdl.MouseWheelEvent:
				Mouse.X,
					Mouse.Y,
					Mouse.State = sdl.GetMouseState() // Update engine.Mouse
			}
			processEvent(event)
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
	Width, Height = Window.GetSize()
}

func EndGame() {
	running = false
}
