package game

// import (
// 	"github.com/Sirupsen/logrus"
// 	"github.com/veandco/go-sdl2/sdl"
// 	"github.com/veandco/go-sdl2/sdl_image"
// 	"go.owls.io/fauxbox/engine"
// )

// ////////////////////////////////////////////////////////////////////////////////
// // SYSTEM //
// ///////////

// type FauxboxIMG struct {
// 	logoTex *sdl.Texture
// 	logoObj sdl.Rect
// }

// func init() {
// 	engine.Register(&FauxboxIMG{})
// }

// ////////////////////////////////////////////////////////////////////////////////
// // CALLBACKS //
// //////////////

// func (fimg *FauxboxIMG) Load(resPath string) {
// 	// BITMAPS
// 	surface, err := img.Load(resPath + "/fauxbox.tga")
// 	if err != nil {
// 		logrus.WithError(err).Error("Loading bitmap")
// 	}

// 	if fimg.logoTex, err = engine.Renderer.CreateTextureFromSurface(surface); err != nil {
// 		logrus.WithError(err).Error("Texture from surface")
// 	}
// 	surface.Free()

// 	fimg.logoObj = sdl.Rect{int32(World.X(200)), int32(World.Y(200)), 256, 256}
// }

// func (fimg *FauxboxIMG) Update(dt float64) {
// 	fimg.logoObj.X = int32(World.X(200))
// 	fimg.logoObj.Y = int32(World.Y(200))
// }

// func (fimg *FauxboxIMG) Draw(dt float64) {
// 	engine.Renderer.Copy(fimg.logoTex, nil, &fimg.logoObj)
// }
