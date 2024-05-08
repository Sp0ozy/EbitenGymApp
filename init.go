package main

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

func NewApp() *app {
	return &app{
		apphud:      NewHud(),
		apphome:     NewHome(),
		appschedule: NewSchedule(),
		apptr:       NewTournament(),
		apppr:       NewProgress(),
		workoutplan: NewWorkout(),
		profile:     NewProfile(),
		appIn:       NewSignIn(),
		appUp:       NewSignUp(),
		exercises:   NewallExercises(),
		benchPress:  benchInit(),
		addEx:       NewAddExercise(),
		condition:   6,
	}
}
func benchInit() exInfo {
	return exInfo{
		background: newImage("assets/benchbg.png"),
		delete:     NewButton(202, 691, 175, 50, 4, "assets/delete0.png", "assets/delete1.png", "assets/delete2.png"),
		settings:   NewButton(17, 691, 175, 50, 10, "assets/change0.png", "assets/change1.png", "assets/change2.png"),
	}
}
func NewHud() hud {
	return hud{
		hudUP:       []*ebiten.Image{newImage("assets/0.png"), newImage("assets/1.png"), newImage("assets/2.png"), newImage("assets/3.png"), newImage("assets/4.png"), newImage("assets/5.png"), newImage("assets/5.png"), newImage("assets/5.png"), newImage("assets/8.png"), newImage("assets/9.png"), newImage("assets/9.png")},
		lowerhud:    newImage("assets/lowerhud.png"),
		home:        NewButton(31, 775, 51, 48, 0, "assets/homedef.png", "assets/homehover.png", "assets/homeclicked.png"),
		schedule:    NewButton(131, 778, 45, 45, 1, "assets/scheduledef.png", "assets/schedulehover.png", "assets/scheduleclicked.png"),
		tournaments: NewButton(225, 778, 41, 44, 2, "assets/tournamentdef.png", "assets/tournamenthover.png", "assets/tournamentclicked.png"),
		progress:    NewButton(315, 774, 47, 51, 3, "assets/progressdef.png", "assets/progresshover.png", "assets/progressclicked.png"),
		profilePic:  NewButton(25, 43, 58, 58, 5, "assets/profdef.png", "assets/profhover.png", "assets/profclicked.png"),
	}
}
func NewHome() home {
	return home{
		background: newImage("assets/home.png"),
	}

}
func NewAddExercise() settings {
	return settings{
		background: newImage("assets/settingsbg.png"),
		save:       NewButton(213, 680, 140, 45, 4, "assets/save0.png", "assets/save1.png", "assets/save2.png"),
		exit:       NewButton(40, 680, 140, 45, 4, "assets/exit0.png", "assets/exit1.png", "assets/exit2.png"),
		weight:     textinput{57, 526, 280, 20, false, true, "Weight", "80", newFont("fonts/Roboto-Light.ttf", 16, 96)},
		sets:       textinput{57, 605, 280, 20, false, true, "Sets", "4", newFont("fonts/Roboto-Light.ttf", 16, 96)},
		reps:       textinput{57, 447, 280, 20, false, true, "Repetitions", "8", newFont("fonts/Roboto-Light.ttf", 16, 96)},
	}
}
func NewSignUp() signUp {
	return signUp{
		background: newImage("assets/signupbg.png"),
		login:      textinput{57, 288, 280, 20, false, true, "Login", "", newFont("fonts/Roboto-Light.ttf", 16, 96)},
		password:   textinput{57, 525, 280, 20, false, false, "Password", "", newFont("fonts/Roboto-Light.ttf", 16, 96)},
		email:      textinput{57, 367, 280, 20, false, true, "Email", "", newFont("fonts/Roboto-Light.ttf", 16, 96)},
		name:       textinput{57, 446, 280, 20, false, true, "Name", "", newFont("fonts/Roboto-Light.ttf", 16, 96)},
		password2:  textinput{57, 604, 280, 20, false, false, "Repeat password", "", newFont("fonts/Roboto-Light.ttf", 16, 96)},
		signin:     NewButton(40, 680, 147, 45, 6, "assets/signIndef2.png", "assets/signInhover2.png", "assets/signInclicked2.png"),
		signup:     NewButton(213, 680, 140, 45, 0, "assets/signUpdef2.png", "assets/signUphover2.png", "assets/signUpclicked2.png"),
	}
}
func NewSignIn() signIn {
	return signIn{
		background: newImage("assets/signinbg.png"),
		login:      textinput{57, 441, 280, 20, false, true, "Login or email", "", newFont("fonts/Roboto-Light.ttf", 16, 96)},
		password:   textinput{57, 521, 280, 20, false, false, "Password", "", newFont("fonts/Roboto-Light.ttf", 16, 96)},
		signin:     NewButton(213, 600, 140, 45, 0, "assets/signIndef.png", "assets/signInhover.png", "assets/signInclicked.png"),
		signup:     NewButton(40, 600, 140, 45, 7, "assets/signUpdef.png", "assets/signUphover.png", "assets/signUpclicked.png"),
	}
}
func NewallExercises() allexercises {
	btn := NewButton(257, 16, 88, 33, 3, "assets/Get0.png", "assets/Get1.png", "assets/Get2.png")
	btn1 := NewButton(257, 16, 88, 33, 3, "assets/Get0.png", "assets/Get1.png", "assets/Get2.png")
	btn2 := NewButton(257, 16, 88, 33, 3, "assets/Get0.png", "assets/Get1.png", "assets/Get2.png")
	btn3 := NewButton(257, 16, 88, 33, 3, "assets/Get0.png", "assets/Get1.png", "assets/Get2.png")
	btn4 := NewButton(257, 16, 88, 33, 3, "assets/Get0.png", "assets/Get1.png", "assets/Get2.png")
	btn5 := NewButton(257, 16, 88, 33, 3, "assets/Get0.png", "assets/Get1.png", "assets/Get2.png")
	btn6 := NewButton(257, 16, 88, 33, 3, "assets/Get0.png", "assets/Get1.png", "assets/Get2.png")

	text := NewTextBox(10, 25, 180, 14, "Chest dip", newFont("fonts/Roboto-Light.ttf", 16, 96))
	text1 := NewTextBox(10, 25, 180, 14, "Bench press", newFont("fonts/Roboto-Light.ttf", 16, 96))
	text2 := NewTextBox(10, 25, 180, 14, "Incline bench  press", newFont("fonts/Roboto-Light.ttf", 14, 96))
	text3 := NewTextBox(10, 25, 180, 14, "Chest fly machine", newFont("fonts/Roboto-Light.ttf", 16, 96))
	text4 := NewTextBox(10, 25, 180, 14, "Cable Triceps", newFont("fonts/Roboto-Light.ttf", 16, 96))
	text5 := NewTextBox(10, 25, 180, 12, "Biceps Curl", newFont("fonts/Roboto-Light.ttf", 16, 96))
	text6 := NewTextBox(10, 25, 180, 14, "Triceps Pushdown", newFont("fonts/Roboto-Light.ttf", 16, 96))
	// text7 := NewTextBox(10, 25, 180, 14, "Overhead press", newFont("fonts/Roboto-Light.ttf", 16, 96))

	cnt := NewContainer(0, 0, 360, 65, "assets/conatinerbg.png", []element{&btn, &text})
	cnt1 := NewContainer(0, 85, 360, 65, "assets/conatinerbg.png", []element{&btn1, &text1})
	cnt2 := NewContainer(0, 170, 360, 65, "assets/conatinerbg.png", []element{&btn2, &text2})
	cnt3 := NewContainer(0, 255, 360, 65, "assets/conatinerbg.png", []element{&btn3, &text3})
	cnt4 := NewContainer(0, 340, 360, 65, "assets/conatinerbg.png", []element{&btn4, &text4})
	cnt5 := NewContainer(0, 425, 360, 65, "assets/conatinerbg.png", []element{&btn5, &text5})
	cnt6 := NewContainer(0, 510, 360, 65, "assets/conatinerbg.png", []element{&btn6, &text6})
	// cnt7 := NewContainer(0, 595, 360, 65, "assets/conatinerbg.png", []element{&btn3, &btn14, &text7})
	return allexercises{
		background: newImage("assets/list.png"),
		exercises:  NewList(17, 188, 360, 612, 20, 0, []element{&cnt, &cnt1, &cnt2, &cnt3, &cnt4, &cnt5, &cnt6}),
		search:     textinput{97, 135, 250, 16, false, true, "Search for the exercise", "", newFont("fonts/Roboto-Light.ttf", 16, 96)},

		// exercises:  NewList(17, 145, 360, 65, 20, 0, nil),
	}
}

func NewProfile() prof {
	return prof{
		background: newImage("assets/profbg.png"),
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
		moreinfo:   NewButton(17, 521, 360, 59, 8, "assets/moredef.png", "assets/morehover.png", "assets/moreclicked.png"),
	}
}
func NewWorkout() workout {
	btn := NewButton(197, 16, 88, 33, 9, "assets/settingsdef.png", "assets/settingshover.png", "assets/settingsclicked.png")
	btn1 := NewButton(197, 16, 88, 33, 9, "assets/settingsdef.png", "assets/settingshover.png", "assets/settingsclicked.png")
	btn2 := NewButton(197, 16, 88, 33, 9, "assets/settingsdef.png", "assets/settingshover.png", "assets/settingsclicked.png")
	btn3 := NewButton(197, 16, 88, 33, 9, "assets/settingsdef.png", "assets/settingshover.png", "assets/settingsclicked.png")
	btn4 := NewButton(197, 16, 88, 33, 9, "assets/settingsdef.png", "assets/settingshover.png", "assets/settingsclicked.png")
	btn5 := NewButton(197, 16, 88, 33, 9, "assets/settingsdef.png", "assets/settingshover.png", "assets/settingsclicked.png")
	btn6 := NewButton(197, 16, 88, 33, 9, "assets/settingsdef.png", "assets/settingshover.png", "assets/settingsclicked.png")

	btn7 := NewTickButton(304, 13, 88, 33, 1, "assets/tickbuttondef.png", "assets/tickbuttonhover.png", "assets/2tickbuttondef.png", "assets/2tickbuttonhover.png", false)
	btn8 := NewTickButton(304, 13, 88, 33, 1, "assets/tickbuttondef.png", "assets/tickbuttonhover.png", "assets/2tickbuttondef.png", "assets/2tickbuttonhover.png", false)
	btn9 := NewTickButton(304, 13, 88, 33, 1, "assets/tickbuttondef.png", "assets/tickbuttonhover.png", "assets/2tickbuttondef.png", "assets/2tickbuttonhover.png", false)
	btn10 := NewTickButton(304, 13, 88, 33, 1, "assets/tickbuttondef.png", "assets/tickbuttonhover.png", "assets/2tickbuttondef.png", "assets/2tickbuttonhover.png", false)
	btn11 := NewTickButton(304, 13, 88, 33, 1, "assets/tickbuttondef.png", "assets/tickbuttonhover.png", "assets/2tickbuttondef.png", "assets/2tickbuttonhover.png", false)
	btn12 := NewTickButton(304, 13, 88, 33, 1, "assets/tickbuttondef.png", "assets/tickbuttonhover.png", "assets/2tickbuttondef.png", "assets/2tickbuttonhover.png", false)
	btn13 := NewTickButton(304, 13, 88, 33, 1, "assets/tickbuttondef.png", "assets/tickbuttonhover.png", "assets/2tickbuttondef.png", "assets/2tickbuttonhover.png", false)
	// btn14 := NewTickButton(304, 13, 88, 33, 1, "assets/tickbuttondef.png", "assets/tickbuttonhover.png", "assets/2tickbuttondef.png", "assets/2tickbuttonhover.png", false)

	text := NewTextBox(10, 25, 180, 14, "Chest dip", newFont("fonts/Roboto-Light.ttf", 16, 96))
	text1 := NewTextBox(10, 25, 180, 14, "Bench press", newFont("fonts/Roboto-Light.ttf", 16, 96))
	text2 := NewTextBox(10, 25, 180, 14, "Incline bench  press", newFont("fonts/Roboto-Light.ttf", 14, 96))
	text3 := NewTextBox(10, 25, 180, 14, "Chest fly machine", newFont("fonts/Roboto-Light.ttf", 16, 96))
	text4 := NewTextBox(10, 25, 180, 14, "Cable Triceps", newFont("fonts/Roboto-Light.ttf", 16, 96))
	text5 := NewTextBox(10, 25, 180, 12, "Biceps Curl", newFont("fonts/Roboto-Light.ttf", 16, 96))
	text6 := NewTextBox(10, 25, 180, 14, "Triceps Pushdown", newFont("fonts/Roboto-Light.ttf", 16, 96))
	// text7 := NewTextBox(10, 25, 180, 14, "Overhead press", newFont("fonts/Roboto-Light.ttf", 16, 96))

	cnt := NewContainer(0, 0, 360, 65, "assets/conatinerbg.png", []element{&btn, &btn7, &text})
	cnt1 := NewContainer(0, 85, 360, 65, "assets/conatinerbg.png", []element{&btn1, &btn8, &text1})
	cnt2 := NewContainer(0, 170, 360, 65, "assets/conatinerbg.png", []element{&btn2, &btn9, &text2})
	cnt3 := NewContainer(0, 255, 360, 65, "assets/conatinerbg.png", []element{&btn3, &btn10, &text3})
	cnt4 := NewContainer(0, 340, 360, 65, "assets/conatinerbg.png", []element{&btn4, &btn11, &text4})
	cnt5 := NewContainer(0, 425, 360, 65, "assets/conatinerbg.png", []element{&btn5, &btn12, &text5})
	cnt6 := NewContainer(0, 510, 360, 65, "assets/conatinerbg.png", []element{&btn6, &btn13, &text6})
	// cnt7 := NewContainer(0, 595, 360, 65, "assets/conatinerbg.png", []element{&btn3, &btn14, &text7})
	return workout{
		background: newImage("assets/background.png"),
		exercises:  NewList(17, 145, 360, 612, 20, 0, []element{&cnt, &cnt1, &cnt2, &cnt3, &cnt4, &cnt5, &cnt6}),
		plus:       NewButton(302, 660, 75, 75, 8, "assets/plusIcondef.png", "assets/plusIconhover.png", "assets/plusIconclicked.png"),
		// exercises:  NewList(17, 145, 360, 65, 20, 0, nil),
	}
}
func NewContainer(x, y, w, h int, root string, el []element) container {
	return container{
		posX:       x,
		posY:       y,
		sizeX:      w,
		sizeY:      h,
		elements:   el,
		background: newImage(root),
		box:        ebiten.NewImage(w, h),
	}
}
func NewList(x, y, w, h, space int, scroll float64, el []element) list {
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
func NewTextBox(x, y, w, h int, content string, f font.Face) textbox {
	return textbox{
		posX:  x,
		posY:  y,
		sizeX: w,
		sizeY: h,
		text:  content,
		font:  f,
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
func NewTickButton(x, y, w, h, dir int, p1, p2, p3, p4 string, t bool) tickbutton {
	return tickbutton{
		state: 0,
		image: [4]*ebiten.Image{newImage(p1), newImage(p2), newImage(p3), newImage(p4)},
		posX:  x,
		posY:  y,
		sizeX: w,
		sizeY: h,
		tick:  false,
	}

}
func newImage(name string) *ebiten.Image {
	rect, _, err := ebitenutil.NewImageFromFile(name)
	if err != nil {
		log.Fatal(err)
	}

	return rect
}

func newFont(name string, size, dpi float64) font.Face {
	fontBytes, err := os.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}

	tt, err := opentype.Parse(fontBytes)
	if err != nil {
		log.Fatal(err)
	}
	font, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    size,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatal(err)
	}
	return font
}
