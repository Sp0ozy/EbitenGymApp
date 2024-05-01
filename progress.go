package main

import "github.com/hajimehoshi/ebiten/v2"

func (a *app) drawProgress(screen *ebiten.Image) {
	screen.DrawImage(a.apppr.background, nil)
	a.apppr.moreinfo.draw(screen)
}
func (a *app) updateProgress() {
	a.condition = a.apppr.moreinfo.update(a.condition, 0, 0, 0, 0, 0)

}
