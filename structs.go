package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
)

const (
	winTitle     = "GymApp"
	screenWidth  = 393
	screenHeight = 852
	// cellSize     = 5
	// damageTicker = time.Second
	// dpi          = 200

)
const (
	def = iota
	hover
	clicked
)

type (
	app struct {
		apphud      hud
		apphome     home
		appschedule schedule
		apptr       tournament
		apppr       progress
		workoutplan workout
		condition   int
	}
	hud struct {
		hudYP, lowerhud                                   *ebiten.Image
		profilePic, home, schedule, tournaments, progress button
	}
	home struct {
		background *ebiten.Image
	}
	schedule struct {
		background                                                     *ebiten.Image
		monday, tuesday, wednesday, thursday, friday, saturday, sunday button
	}
	tournament struct {
		background         *ebiten.Image
		headIcon, semiIcon *ebiten.Image
	}
	progress struct {
		background *ebiten.Image
		moreinfo   button
	}
	workout struct {
		background *ebiten.Image
		exercises  list
		plus       button
	}
	container struct {
		box          *ebiten.Image
		elements     []element
		posX, posY   int
		sizeX, sizeY int
		background   *ebiten.Image
	}
	textbox struct {
		posX, posY   int
		sizeX, sizeY int
		text         string
		font         font.Face
	}

	/* Field button has 3 states( default, hover, clicked), 3 images to support those stages,
	position on X and Y axis, Horizontal and Vertical size, direction to which butoon changes the game if clicked. */

	button struct {
		state        int
		image        [3]*ebiten.Image
		posX, posY   int
		sizeX, sizeY int
		direction    int
	}
	tickbutton struct {
		state        int
		image        [4]*ebiten.Image
		posX, posY   int
		sizeX, sizeY int
		tick         bool
	}
	list struct {
		box          *ebiten.Image
		elements     []element
		posX, posY   int
		sizeX, sizeY int
		spacer       int
		offset       float64
	}
	element interface {
		offset(f float64)
		pos() (int, int)
		size() (int, int)
		draw(screen *ebiten.Image)
		update(cnd, posX, posY, upperMargin, lowerMargin, offset int) int
	}
)
