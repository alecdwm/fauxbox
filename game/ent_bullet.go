package game

// import "github.com/go-gl/mathgl/mgl64"

// ////////////////////////////////////////////////////////////////////////////////
// // ENT //
// ////////

// type Bullet struct {
// 	Entity

// 	active    bool
// 	livedtime float64
// 	lifetime  float64

// 	pos mgl64.Vec2
// 	vel mgl64.Vec2
// }

// ////////////////////////////////////////////////////////////////////////////////
// // GETTER //
// ///////////

// type BulletEnt interface {
// 	Get() *Bullet
// }

// func (b *Bullet) Get() *Bullet {
// 	return b
// }

// var BulletIDs []uint

// func (s *State) NewBullet() *Bullet {
// 	newBullet := &Bullet{
// 		Entity{ID: s.GetNextEntID()},

// 		false,
// 		0,
// 		0,

// 		mgl64.Vec2{},
// 		mgl64.Vec2{},
// 	}

// 	BulletIDs = append(BulletIDs, newBullet.ID)

// 	s.Ents[newBullet.ID] = newBullet

// 	return newBullet
// }

// func (s *State) GetBullets() (bulletEnts []*Bullet) {
// 	for _, id := range BulletIDs {
// 		bulletEnts = append(bulletEnts, s.Ents[id].(Bullet))
// 	}
// 	return bulletEnts
// }

// ////////////////////////////////////////////////////////////////////////////////
// // FUNCTIONS //
// //////////////

// func (b *Bullet) Fire(pos, vel mgl64.Vec2, lifetime float64) {
// 	if b.active {
// 		return
// 	}

// 	b.active = true
// 	b.livedtime = 0
// 	b.lifetime = lifetime
// 	b.pos = pos
// 	b.vel = vel
// }
