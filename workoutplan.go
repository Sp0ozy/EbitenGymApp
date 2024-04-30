package main

import "github.com/hajimehoshi/ebiten/v2"

// import (
// 	"github.com/hajimehoshi/ebiten/v2"
// 	"github.com/hajimehoshi/ebiten/v2/inpututil"
// )

func (a *app) drawWorkout(screen *ebiten.Image) {
	screen.DrawImage(a.workoutplan.background, nil)
	// drawList(screen, [])
	a.workoutplan.exercises.draw(screen)
}

func (a *app) updateWorkout() {
	a.condition = a.workoutplan.exercises.update(a.condition, 0, 0)
}

// func drawList(screen *ebiten.Image, btn button) {
// 	op := &ebiten.DrawImageOptions{}
// 	op.GeoM.Translate(float64(btn.posX), float64(btn.posY))
// 	screen.DrawImage(btn.image[btn.state], op)
// }
// func (b *button) updateList(cnd int) int {
// 	x, y := ebiten.CursorPosition()
// 	if x > b.posX && x < b.posX+b.sizeX && y > b.posY && y < b.posY+b.sizeY {
// 		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
// 			b.state = clicked
// 			return b.direction
// 		}
// 		b.state = hover
// 		return cnd
// 	}
// 	b.state = def
// 	return cnd
// }
