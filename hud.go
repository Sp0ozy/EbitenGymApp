package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func (a *app) drawHud(screen *ebiten.Image) {
	screen.DrawImage(a.apphud.hudUP[a.condition], nil)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 757)
	screen.DrawImage(a.apphud.lowerhud, op)
	a.apphud.profilePic.draw(screen)
	a.apphud.home.draw(screen)
	a.apphud.schedule.draw(screen)
	a.apphud.tournaments.draw(screen)
	a.apphud.progress.draw(screen)

}
func (a *app) updateHud() {
	a.condition = a.apphud.profilePic.update(a.condition, 0, 0, 0, 0, 0)
	a.condition = a.apphud.home.update(a.condition, 0, 0, 0, 0, 0)
	a.condition = a.apphud.schedule.update(a.condition, 0, 0, 0, 0, 0)
	a.condition = a.apphud.tournaments.update(a.condition, 0, 0, 0, 0, 0)
	a.condition = a.apphud.progress.update(a.condition, 0, 0, 0, 0, 0)
}
