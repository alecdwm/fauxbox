package game

import (
	"github.com/Sirupsen/logrus"
	"github.com/dradtke/go-allegro/allegro"
	"github.com/dradtke/go-allegro/allegro/dialog"
	"go.owls.io/fauxbox/engine"
)

////////////////////////////////////////////////////////////////////////////////
// SYSTEM //////////////////////////////////////////////////////////////////////
////////////

type FauxboxIMG struct {
	logo *allegro.Bitmap
}

func init() {
	engine.Register(&FauxboxIMG{})
}

////////////////////////////////////////////////////////////////////////////////
// CALLBACKS ///////////////////////////////////////////////////////////////////
///////////////

func (fimg *FauxboxIMG) Load(resPath string, textLog *dialog.TextLog) {
	var err error

	// BITMAPS
	textLog.Appendln("FauxBoxIMG: Loading bitmaps...")
	if fimg.logo, err = allegro.LoadBitmap(resPath + "/fauxbox.tga"); err != nil {
		logrus.WithError(err).Error("Loading bitmap")
	}
}

func (fimg *FauxboxIMG) Draw(dt float64) {
	fimg.logo.Draw(World.X(200), World.Y(200), allegro.FLIP_NONE)
}
