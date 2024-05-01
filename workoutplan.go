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
