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
	a.workoutplan.plus.draw(screen)
}

func (a *app) updateWorkout() {
	if cnd := a.workoutplan.plus.update(a.condition, 0, 0, 0, 0, 0); a.condition != cnd {
		a.condition = cnd
		return

	}
	a.condition = a.workoutplan.exercises.update(a.condition, 0, 0)

}
