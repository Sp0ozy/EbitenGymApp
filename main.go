package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func (a *app) Layout(outWidth, outHeight int) (w, h int) { return screenWidth, screenHeight }

func (a *app) Update() error {
	switch a.condition {
	case 0:
		a.updateHome()
		a.updateHud()
	case 1:
		a.updateSchedule()
		a.updateHud()
	case 2:
		a.updateTournament()
		a.updateHud()
	case 3:
		a.updateProgress()
		a.updateHud()
	case 4:
		a.updateWorkout()
		a.updateHud()
	case 5:
		a.updateProfile()
		a.updateHud()
	case 6:
		a.updateSignIn()
	case 7:
		a.updateSignUp()
	case 8:
		a.updateExercises()
		a.updateHud()
	case 9:
		a.updateInfo()
		a.updateHud()
	case 10:
		a.updateEx()
		a.updateHud()
	}
	return nil
}

func (a *app) Draw(screen *ebiten.Image) {
	switch a.condition {
	case 0:
		a.drawHome(screen)
		a.drawHud(screen)
	case 1:
		a.drawSchedule(screen)
		a.drawHud(screen)
	case 2:
		a.drawTournament(screen)
		a.drawHud(screen)
	case 3:
		a.drawProgress(screen)
		a.drawHud(screen)
	case 4:
		a.drawWorkout(screen)
		a.drawHud(screen)
	case 5:
		a.drawProfile(screen)
		a.drawHud(screen)
	case 6:
		a.drawSignIn(screen)
	case 7:
		a.drawSignUp(screen)
	case 8:
		a.drawExercises(screen)
		a.drawHud(screen)
	case 9:
		a.drawInfo(screen)
		a.drawHud(screen)
	case 10:
		a.drawEx(screen)
		a.drawHud(screen)
	}
}

func main() {
	ebiten.SetWindowTitle(winTitle)
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowResizable(true)
	a := NewApp()
	if err := ebiten.RunGame(a); err != nil {
		log.Fatal(err)
	}
}
