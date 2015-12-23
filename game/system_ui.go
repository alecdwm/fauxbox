package game

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"
	"go.owls.io/fauxbox/engine"
)

////////////////////////////////////////////////////////////////////////////////
// SYSTEM //
///////////

type UI struct {
	font           *ttf.Font
	fpsCounterText *sdl.Surface
	fpsCounterTex  *sdl.Texture
	fpsCounterObj  *sdl.Rect
	helloWorldTex  *sdl.Texture
	helloWorldObj  *sdl.Rect
}

func init() {
	engine.Register(&UI{})
}

////////////////////////////////////////////////////////////////////////////////
// STATES //
///////////

var UIStates map[engine.State]bool = map[engine.State]bool{INGAME: true}

func (ui *UI) States() map[engine.State]bool {
	return UIStates
}

////////////////////////////////////////////////////////////////////////////////
// CALLBACKS //
//////////////

func (ui *UI) Load(resPath string) {
	var err error

	// FONTS
	if ui.font, err = ttf.OpenFont(resPath+"/Neuropol.ttf", 18); err != nil {
		logrus.WithError(err).Error("Loading font")
	}

	// OBJECTS
	// Hello World!
	helloWorldText, err := ui.font.RenderUTF8_Solid("Hello World!", sdl.Color{255, 255, 255, 255})
	if err != nil {
		logrus.WithError(err).Error("Rendering text")
	}
	if ui.helloWorldTex, err = engine.Renderer.CreateTextureFromSurface(helloWorldText); err != nil {
		logrus.WithError(err).Error("Texture from surface")
	}
	ui.helloWorldObj = &sdl.Rect{10, 0, helloWorldText.W, helloWorldText.H}

	// FPS Counter
	ui.fpsCounterObj = &sdl.Rect{10, 20, 0, 0}
}

func (ui *UI) Update(dt float64) {
	var err error
	if ui.fpsCounterText != nil {
		ui.fpsCounterText.Free()
	}
	if ui.fpsCounterText, err = ui.font.RenderUTF8_Solid(fmt.Sprintf("%.1f", engine.FPS), sdl.Color{255, 255, 255, 255}); err != nil {
		logrus.WithError(err).Error("Rendering text")
	}
	if ui.fpsCounterTex != nil {
		ui.fpsCounterTex.Destroy()
	}
	if ui.fpsCounterTex, err = engine.Renderer.CreateTextureFromSurface(ui.fpsCounterText); err != nil {
		logrus.WithError(err).Error("Texture from surface")
	}
	ui.fpsCounterObj.W = ui.fpsCounterText.W
	ui.fpsCounterObj.H = ui.fpsCounterText.H
}

func (ui *UI) Draw(dt float64) {
	engine.Renderer.SetDrawColor(255, 255, 255, 255)
	engine.Renderer.Copy(ui.helloWorldTex, nil, ui.helloWorldObj)
	engine.Renderer.Copy(ui.fpsCounterTex, nil, ui.fpsCounterObj)
}
