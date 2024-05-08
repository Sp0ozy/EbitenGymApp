package main

import "github.com/hajimehoshi/ebiten/v2"

func (a *app) drawInfo(screen *ebiten.Image) {
	screen.DrawImage(a.benchPress.background, nil)
	a.benchPress.delete.draw(screen)
	a.benchPress.settings.draw(screen)
}
func (a *app) updateInfo() {
	a.condition = a.benchPress.settings.update(a.condition, 0, 0, 0, 0, 0)
	a.condition = a.benchPress.delete.update(a.condition, 0, 0, 0, 0, 0)

}
