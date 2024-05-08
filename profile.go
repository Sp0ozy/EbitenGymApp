package main

import "github.com/hajimehoshi/ebiten/v2"

func (a *app) drawProfile(screen *ebiten.Image) {
	screen.DrawImage(a.profile.background, nil)
}
func (a *app) updateProfile() {

}
