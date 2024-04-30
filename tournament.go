package main

import "github.com/hajimehoshi/ebiten/v2"

func (a *app) drawTournament(screen *ebiten.Image) {
	screen.DrawImage(a.apptr.background, nil)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(122, 144)
	screen.DrawImage(a.apptr.headIcon, op)
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(42, 414)
	screen.DrawImage(a.apptr.semiIcon, op)
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(301, 414)
	screen.DrawImage(a.apptr.semiIcon, op)
}
func (a *app) updateTournament() {
}
