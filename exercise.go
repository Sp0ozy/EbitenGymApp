package main

import "github.com/hajimehoshi/ebiten/v2"

// import (
// 	"github.com/hajimehoshi/ebiten/v2"
// 	"github.com/hajimehoshi/ebiten/v2/inpututil"
// )

func (a *app) drawExercises(screen *ebiten.Image) {
	screen.DrawImage(a.exercises.background, nil)
	a.exercises.search.draw(screen)
	if a.exercises.search.text == "c" {
		btn := NewButton(257, 16, 88, 33, 3, "assets/Get0.png", "assets/Get1.png", "assets/Get2.png")
		btn1 := NewButton(257, 16, 88, 33, 3, "assets/Get0.png", "assets/Get1.png", "assets/Get2.png")
		btn2 := NewButton(257, 16, 88, 33, 3, "assets/Get0.png", "assets/Get1.png", "assets/Get2.png")

		text := NewTextBox(10, 25, 180, 14, "Chest dip", newFont("fonts/Roboto-Light.ttf", 16, 96))
		text1 := NewTextBox(10, 25, 180, 14, "Chest fly machine", newFont("fonts/Roboto-Light.ttf", 16, 96))
		text2 := NewTextBox(10, 25, 180, 14, "Cable Triceps", newFont("fonts/Roboto-Light.ttf", 16, 96))

		cnt := NewContainer(0, 0, 360, 65, "assets/conatinerbg.png", []element{&btn, &text})
		cnt1 := NewContainer(0, 85, 360, 65, "assets/conatinerbg.png", []element{&btn1, &text1})
		cnt2 := NewContainer(0, 170, 360, 65, "assets/conatinerbg.png", []element{&btn2, &text2})

		tmp := NewList(17, 188, 360, 612, 20, 0, []element{&cnt, &cnt1, &cnt2})
		tmp.draw(screen)
	} else if a.exercises.search.text == "ch" {
		btn := NewButton(257, 16, 88, 33, 3, "assets/Get0.png", "assets/Get1.png", "assets/Get2.png")
		btn1 := NewButton(257, 16, 88, 33, 3, "assets/Get0.png", "assets/Get1.png", "assets/Get2.png")

		text := NewTextBox(10, 25, 180, 14, "Chest dip", newFont("fonts/Roboto-Light.ttf", 16, 96))
		text1 := NewTextBox(10, 25, 180, 14, "Chest fly machine", newFont("fonts/Roboto-Light.ttf", 16, 96))

		cnt := NewContainer(0, 0, 360, 65, "assets/conatinerbg.png", []element{&btn, &text})
		cnt1 := NewContainer(0, 85, 360, 65, "assets/conatinerbg.png", []element{&btn1, &text1})

		tmp := NewList(17, 188, 360, 612, 20, 0, []element{&cnt, &cnt1})
		tmp.draw(screen)
	} else {
		a.exercises.exercises.draw(screen)
	}
}

func (a *app) updateExercises() {
	a.condition = a.exercises.search.update(a.condition, 0, 0, 0, 0, 0)

	a.condition = a.exercises.exercises.update(a.condition, 0, 0)

}
