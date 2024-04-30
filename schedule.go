package main

import "github.com/hajimehoshi/ebiten/v2"

func (a *app) drawSchedule(screen *ebiten.Image) {
	screen.DrawImage(a.appschedule.background, nil)
	a.appschedule.monday.draw(screen)
	a.appschedule.tuesday.draw(screen)
	a.appschedule.wednesday.draw(screen)
	a.appschedule.thursday.draw(screen)
	a.appschedule.friday.draw(screen)
	a.appschedule.saturday.draw(screen)
	a.appschedule.sunday.draw(screen)
}
func (a *app) updateSchedule() {
	a.condition = a.appschedule.monday.update(a.condition, 0, 0, 0, 0)
	a.condition = a.appschedule.tuesday.update(a.condition, 0, 0, 0, 0)
	a.condition = a.appschedule.wednesday.update(a.condition, 0, 0, 0, 0)
	a.condition = a.appschedule.thursday.update(a.condition, 0, 0, 0, 0)
	a.condition = a.appschedule.friday.update(a.condition, 0, 0, 0, 0)
	a.condition = a.appschedule.saturday.update(a.condition, 0, 0, 0, 0)
	a.condition = a.appschedule.sunday.update(a.condition, 0, 0, 0, 0)
}
