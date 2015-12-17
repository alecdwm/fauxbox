package engine

import (
	"github.com/Sirupsen/logrus"
	"github.com/dradtke/go-allegro/allegro"
	"github.com/dradtke/go-allegro/allegro/dialog"
	"github.com/kardianos/osext"
)

////////////////////////////////////////////////////////////////////////////////
// LOADER //////////////////////////////////////////////////////////////////////
////////////

type Loader interface {
	Load(resourcePath string, textLog *dialog.TextLog)
}

func load() {
	// CREATE TEXTLOG
	textLog, err := dialog.OpenNativeTextLog("Loading Resources",
		dialog.TEXTLOG_NO_CLOSE|dialog.TEXTLOG_MONOSPACE)
	if err != nil {
		logrus.WithError(err).Error("Opening loading textlog")
	}
	defer textLog.Close()

	// DETERMINE RESOURCE PATH
	textLog.Appendln("Determining resource path...")
	execPath, err := osext.ExecutableFolder()
	if err != nil {
		logrus.WithError(err).Error("Getting executable's path")
	}
	resPath := execPath + "/resources"

	// CALL LOADERS
	for _, system := range systems {
		if loader, ok := system.(Loader); ok {
			loader.Load(resPath, textLog)
		}
	}
}

////////////////////////////////////////////////////////////////////////////////
// PROCESS EVENTS //////////////////////////////////////////////////////////////
////////////////////

type EventProcessor interface {
	ProcessEvent(event interface{})
}

func processEvent(event interface{}) {
	CallEventProcessors := func(event interface{}) {
		for _, system := range systems {
			if eventProcessor, ok := system.(EventProcessor); ok {
				eventProcessor.ProcessEvent(event)
			}
		}
	}

	switch event.(type) {
	case allegro.DisplayCloseEvent:
		CallEventProcessors(event)
		EndGame()

	case allegro.DisplayResizeEvent:
		Resized()
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
	allegro.ClearToColor(allegro.MapRGB(0, 0, 0))

	for _, system := range systems {
		if drawer, ok := system.(Drawer); ok {
			drawer.Draw(dt)
		}
	}

	allegro.FlipDisplay()
}
