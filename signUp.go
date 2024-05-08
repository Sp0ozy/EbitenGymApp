package main

import "github.com/hajimehoshi/ebiten/v2"

func (a *app) drawSignUp(screen *ebiten.Image) {
	screen.DrawImage(a.appUp.background, nil)
	a.appUp.login.draw(screen)
	a.appUp.password.draw(screen)
	a.appUp.password2.draw(screen)
	a.appUp.email.draw(screen)
	a.appUp.name.draw(screen)
	a.appUp.signin.draw(screen)
	a.appUp.signup.draw(screen)
}
func (a *app) updateSignUp() {
	a.condition = a.appUp.login.update(a.condition, 0, 0, 0, 0, 0)
	a.condition = a.appUp.password.update(a.condition, 0, 0, 0, 0, 0)
	a.appUp.password2.update(a.condition, 0, 0, 0, 0, 0)
	a.appUp.email.update(a.condition, 0, 0, 0, 0, 0)
	a.appUp.name.update(a.condition, 0, 0, 0, 0, 0)
	a.condition = a.appUp.signin.update(a.condition, 0, 0, 0, 0, 0)
	a.condition = a.appUp.signup.update(a.condition, 0, 0, 0, 0, 0)
}
