package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func NewApp() *app {
	return &app{
		apphud:      NewHud(),
		apphome:     NewHome(),
		appschedule: NewSchedule(),
		apptr:       NewTournament(),
		apppr:       NewProgress(),
		workoutplan: NewWorkout(),
		condition:   0,
	}
}
func NewHud() hud {
	return hud{
		hudYP:       newImage("assets/hudyourprogress.png"),
		lowerhud:    newImage("assets/lowerhud.png"),
		profilePic:  NewButton(25, 43, 58, 58, 0, "assets/profpic.png", "assets/profilepic.png", "assets/profpic.png"),
		home:        NewButton(31, 775, 51, 48, 0, "assets/homedef.png", "assets/homehover.png", "assets/homeclicked.png"),
		schedule:    NewButton(131, 778, 45, 45, 1, "assets/scheduledef.png", "assets/schedulehover.png", "assets/scheduleclicked.png"),
		tournaments: NewButton(225, 778, 41, 44, 2, "assets/tournamentdef.png", "assets/tournamenthover.png", "assets/tournamentclicked.png"),
		progress:    NewButton(315, 774, 47, 51, 3, "assets/progressdef.png", "assets/progresshover.png", "assets/progressclicked.png"),
	}
}
func NewHome() home {
	return home{
		background: newImage("assets/home.png"),
	}
}
func NewSchedule() schedule {
	return schedule{
		background: newImage("assets/background.png"),
		monday:     NewButton(17, 130, 360, 65, 4, "assets/mondaydef.png", "assets/mondayhover.png", "assets/mondayclicked.png"),
		tuesday:    NewButton(17, 220, 360, 65, 4, "assets/tuesdaydef.png", "assets/tuesdayhover.png", "assets/tuesdayclicked.png"),
		wednesday:  NewButton(17, 310, 360, 65, 4, "assets/wednesdaydef.png", "assets/wednesdayhover.png", "assets/wednesdayclicked.png"),
		thursday:   NewButton(17, 400, 360, 65, 4, "assets/thursdaydef.png", "assets/thursdayhover.png", "assets/thursdayclicked.png"),
		friday:     NewButton(17, 490, 360, 65, 4, "assets/fridaydef.png", "assets/fridayhover.png", "assets/fridayclicked.png"),
		saturday:   NewButton(17, 580, 360, 65, 4, "assets/saturdaydef.png", "assets/saturdayhover.png", "assets/saturdayclicked.png"),
		sunday:     NewButton(17, 670, 360, 65, 4, "assets/sundaydef.png", "assets/sundayhover.png", "assets/sundayclicked.png"),
	}
}
func NewTournament() tournament {
	return tournament{
		background: newImage("assets/trbackground.png"),
		headIcon:   newImage("assets/treadmill.png"),
		semiIcon:   newImage("assets/manIcon.png"),
	}
}
func NewProgress() progress {
	return progress{
		background: newImage("assets/prbackground.png"),
		moreinfo:   NewButton(17, 521, 360, 59, 0, "assets/moredef.png", "assets/morehover.png", "assets/moreclicked.png"),
	}
}
func NewWorkout() workout {
	btn := NewButton(-10, -10, 51, 48, 0, "assets/homedef.png", "assets/homehover.png", "assets/homeclicked.png")
	btn2 := NewButton(70, 10, 51, 48, 0, "assets/homedef.png", "assets/homehover.png", "assets/homeclicked.png")
	return workout{
		background: newImage("assets/background.png"),
		exercises:  NewContainer(17, 145, 360, 65, []element{&btn, &btn2}),
		// exercises:  NewList(17, 145, 360, 65, 20, 0, nil),
	}
}
func NewContainer(x, y, w, h int, el []element) container {
	return container{
		posX:     x,
		posY:     y,
		sizeX:    w,
		sizeY:    h,
		elements: el,
		box:      ebiten.NewImage(w, h),
	}
}
func NewList(x, y, w, h, space int, scroll float64, el []container) list {
	return list{
		posX:     x,
		posY:     y,
		sizeX:    w,
		sizeY:    h,
		spacer:   space,
		offset:   scroll,
		elements: el,
	}
}
func NewButton(x, y, w, h, dir int, p1, p2, p3 string) button {
	return button{
		state:     0,
		image:     [3]*ebiten.Image{newImage(p1), newImage(p2), newImage(p3)},
		posX:      x,
		posY:      y,
		sizeX:     w,
		sizeY:     h,
		direction: dir,
	}

}
func newImage(name string) *ebiten.Image {
	rect, _, err := ebitenutil.NewImageFromFile(name)
	if err != nil {
		log.Fatal(err)
	}

	return rect
}
