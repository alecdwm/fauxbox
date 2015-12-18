package game

import "go.owls.io/fauxbox/engine"

////////////////////////////////////////////////////////////////////////////////
// SYSTEM //////////////////////////////////////////////////////////////////////
////////////

type GameWorld struct {
	CameraX float64
	CameraY float64
	TargetX float64
	TargetY float64
}

var World GameWorld

func init() {
	engine.Register(&World)
}

////////////////////////////////////////////////////////////////////////////////
// CALLBACKS ///////////////////////////////////////////////////////////////////
///////////////

func (w *GameWorld) Update(dt float64) {
	w.CameraX += (w.TargetX - w.CameraX) * dt
	w.CameraY += (w.TargetY - w.CameraY) * dt
}

func (w *GameWorld) Draw(dt float64) {
	// Debugging camera location
	// primitives.DrawCircle(primitives.Point{float32(w.CameraX), float32(w.CameraY)}, 10.0, allegro.MapRGB(255, 0, 255), 2.0)
}

////////////////////////////////////////////////////////////////////////////////
// FUNCTIONS ///////////////////////////////////////////////////////////////////
///////////////

// Convert world-space coordinate into screen-space coordinate
func (w *GameWorld) X(worldX float64) (screenX float32) {
	screenX = float32(worldX-w.CameraX) + float32(engine.Width/2)
	return screenX
}

// Convert world-space coordinate into screen-space coordinate
func (w *GameWorld) Y(worldY float64) (screenY float32) {
	screenY = float32(worldY-w.CameraY) + float32(engine.Height/2)
	return screenY
}

// Convert screen-space coordinate into world-space coordinate
func (w *GameWorld) CamX(cameraX float32) (worldX float64) {
	worldX = float64(cameraX) + w.CameraX - float64(engine.Width/2)
	return worldX
}

// Convert screen-space coordinate into world-space coordinate
func (w *GameWorld) CamY(cameraY float32) (worldY float64) {
	worldY = float64(cameraY) + w.CameraY - float64(engine.Height/2)
	return worldY
}

// Convert screen-space coordinate into world-space coordinate
func (w *GameWorld) CamXInt(cameraX int) (worldX float64) {
	worldX = float64(cameraX) + w.CameraX - float64(engine.Width/2)
	return worldX
}

// Convert screen-space coordinate into world-space coordinate
func (w *GameWorld) CamYInt(cameraY int) (worldY float64) {
	worldY = float64(cameraY) + w.CameraY - float64(engine.Height/2)
	return worldY
}
