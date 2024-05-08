package main

import "github.com/hajimehoshi/ebiten/v2"

func (a *app) drawSignIn(screen *ebiten.Image) {
	screen.DrawImage(a.appIn.background, nil)
	a.appIn.login.draw(screen)
	a.appIn.password.draw(screen)
	a.appIn.signin.draw(screen)
	a.appIn.signup.draw(screen)
}
func (a *app) updateSignIn() {
	a.condition = a.appIn.login.update(a.condition, 0, 0, 0, 0, 0)
	a.condition = a.appIn.password.update(a.condition, 0, 0, 0, 0, 0)
	a.condition = a.appIn.signin.update(a.condition, 0, 0, 0, 0, 0)
	a.condition = a.appIn.signup.update(a.condition, 0, 0, 0, 0, 0)
}
