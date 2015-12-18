package engine

import (
	"github.com/Sirupsen/logrus"
	"github.com/kardianos/osext"
	"github.com/veandco/go-sdl2/sdl"
)

////////////////////////////////////////////////////////////////////////////////
// LOADER //////////////////////////////////////////////////////////////////////
////////////

type Loader interface {
	Load(resourcePath string)
}

func load() {
	// DETERMINE RESOURCE PATH
	execPath, err := osext.ExecutableFolder()
	if err != nil {
		logrus.WithError(err).Error("Getting executable's path")
	}
	resPath := execPath + "/resources"

	// CALL LOADERS
	for _, system := range systems {
		if loader, ok := system.(Loader); ok {
			loader.Load(resPath)
		}
	}
}

////////////////////////////////////////////////////////////////////////////////
// PROCESS EVENTS //////////////////////////////////////////////////////////////
////////////////////

type EventProcessor interface {
	ProcessEvent(event sdl.Event)
}

func processEvent(event sdl.Event) {
	CallEventProcessors := func(event sdl.Event) {
		for _, system := range systems {
			if eventProcessor, ok := system.(EventProcessor); ok {
				eventProcessor.ProcessEvent(event)
			}
		}
	}

	switch e := event.(type) {
	case *sdl.QuitEvent:
		CallEventProcessors(event)
		EndGame()

	case *sdl.WindowEvent:
		if e.Event == sdl.WINDOWEVENT_RESIZED {
			Resized()
		}
		CallEventProcessors(event)

	default:
		CallEventProcessors(event)
	}
}

////////////////////////////////////////////////////////////////////////////////
// UPDATE //////////////////////////////////////////////////////////////////////
////////////

type Updater interface {
	Update(dt float64)
}

func update(dt float64) {
	FPS = 1.0 / dt

	for _, system := range systems {
		if updater, ok := system.(Updater); ok {
			updater.Update(dt)
		}
	}
}

////////////////////////////////////////////////////////////////////////////////
// DRAW ////////////////////////////////////////////////////////////////////////
//////////

type Drawer interface {
	Draw(dt float64)
}

func draw(dt float64) {
	Renderer.SetDrawColor(0, 0, 0, 255)
	Renderer.Clear()

	for _, system := range systems {
		if drawer, ok := system.(Drawer); ok {
			drawer.Draw(dt)
		}
	}

	Renderer.Present()
}
