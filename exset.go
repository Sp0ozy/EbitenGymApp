package main

import "github.com/hajimehoshi/ebiten/v2"

// import (
// 	"github.com/hajimehoshi/ebiten/v2"
// 	"github.com/hajimehoshi/ebiten/v2/inpututil"
// )

func (a *app) drawEx(screen *ebiten.Image) {
	screen.DrawImage(a.addEx.background, nil)
	a.addEx.exit.draw(screen)
	a.addEx.save.draw(screen)
	a.addEx.reps.draw(screen)
	a.addEx.weight.draw(screen)
	a.addEx.sets.draw(screen)
}

func (a *app) updateEx() {
	a.condition = a.addEx.exit.update(a.condition, 0, 0, 0, 0, 0)
	a.condition = a.addEx.save.update(a.condition, 0, 0, 0, 0, 0)
	a.condition = a.addEx.reps.update(a.condition, 0, 0, 0, 0, 0)
	a.condition = a.addEx.weight.update(a.condition, 0, 0, 0, 0, 0)
	a.condition = a.addEx.sets.update(a.condition, 0, 0, 0, 0, 0)

}
