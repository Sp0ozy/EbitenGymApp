package main

import "github.com/hajimehoshi/ebiten/v2"

func (a *app) drawHome(screen *ebiten.Image) {
	screen.DrawImage(a.apphome.background, nil)
}
func (a *app) updateHome() {

}
