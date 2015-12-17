package game

import (
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/dradtke/go-allegro/allegro"
	"github.com/dradtke/go-allegro/allegro/dialog"
	"github.com/dradtke/go-allegro/allegro/font"
	"go.owls.io/fauxbox/engine"
)

////////////////////////////////////////////////////////////////////////////////
// SYSTEM //////////////////////////////////////////////////////////////////////
////////////

type UI struct {
	defaultFont *font.Font
}

func init() {
	engine.Register(&UI{})
}

////////////////////////////////////////////////////////////////////////////////
// CALLBACKS ///////////////////////////////////////////////////////////////////
///////////////

func (ui *UI) Load(resPath string, textLog *dialog.TextLog) {
	var err error

	// FONTS
	textLog.Appendln("UI: Loading fonts...")
	if ui.defaultFont, err = font.LoadFont(resPath+"/Neuropol.ttf", 18, 0); err != nil {
		logrus.WithError(err).Error("Loading font")
	}
}

func (ui *UI) Draw(dt float64) {
	font.DrawText(ui.defaultFont, allegro.MapRGB(255, 255, 255), 20, 20, font.ALIGN_LEFT, "Hello World!")
	font.DrawText(ui.defaultFont, allegro.MapRGB(255, 255, 255), 20, 60, font.ALIGN_LEFT, fmt.Sprintf("%.1f", engine.FPS))
}
