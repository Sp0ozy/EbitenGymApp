package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// TEXTBOX
func (t *textbox) offset(f float64) {
	t.posY += int(f)
}
func (t *textbox) pos() (int, int) {
	return t.posX, t.posY
}
func (t *textbox) size() (int, int) {
	return t.sizeX, t.sizeY
}
func (t *textbox) draw(screen *ebiten.Image) {
	text.Draw(screen, t.text, t.font, t.posX, t.posY+t.sizeY, color.RGBA{0xDB, 0xDB, 0xDB, 0xff})
}

func (t *textbox) update(cnd, posX, posY, upperMargin, lowerMargin, offset int) int {
	return cnd
}

// TextInput
func nchars(b string, n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += b
	}
	return s
}

func (t *textinput) offset(f float64) {
	t.posY += int(f)
}
func (t *textinput) pos() (int, int) {
	return t.posX, t.posY
}
func (t *textinput) size() (int, int) {
	return t.sizeX, t.sizeY
}
func (t *textinput) draw(screen *ebiten.Image) {
	if len(t.text) == 0 {
		text.Draw(screen, t.defalut, t.font, t.posX, t.posY+t.sizeY, color.RGBA{0xDB, 0xDB, 0xDB, 190})
	} else {
		if t.showText {
			text.Draw(screen, t.text, t.font, t.posX, t.posY+t.sizeY, color.RGBA{0xDB, 0xDB, 0xDB, 0xff})
		} else {
			text.Draw(screen, nchars("•", len(t.text)), t.font, t.posX, t.posY+t.sizeY, color.RGBA{0xDB, 0xDB, 0xDB, 0xff})

		}
	}
}

func (t *textinput) update(cnd, posX, posY, upperMargin, lowerMargin, offset int) int {
	x, y := ebiten.CursorPosition()
	if x > t.posX && x < t.posX+t.sizeX && y > posY && y < t.posY+t.sizeY {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
			t.isActive = true
		}
	}
	if !(x > t.posX+posX && x < t.posX+posX+t.sizeX && y > posY+offset+abs(upperMargin-t.posY) && y < t.posY+offset+posY+t.sizeY-lowerMargin) {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
			t.isActive = false
		}
	}

	if t.isActive {
		// keyboard.Open()
		var tmp string
		if inpututil.IsKeyJustPressed(ebiten.KeyA) {
			tmp = "a"
		} else if inpututil.IsKeyJustPressed(ebiten.KeyB) {
			tmp = "b"
		} else if inpututil.IsKeyJustPressed(ebiten.KeyC) {
			tmp = "c"
		} else if inpututil.IsKeyJustPressed(ebiten.KeyD) {
			tmp = "d"
		} else if inpututil.IsKeyJustPressed(ebiten.KeyE) {
			tmp = "e"
		} else if inpututil.IsKeyJustPressed(ebiten.KeyF) {
			tmp = "f"
		} else if inpututil.IsKeyJustPressed(ebiten.KeyG) {
			tmp = "g"
		} else if inpututil.IsKeyJustPressed(ebiten.KeyH) {
			tmp = "h"
		} else if inpututil.IsKeyJustPressed(ebiten.KeyI) {
			tmp = "i"
		} else if inpututil.IsKeyJustPressed(ebiten.KeyJ) {
			tmp = "j"
		} else if inpututil.IsKeyJustPressed(ebiten.KeyK) {
			tmp = "k"
		} else if inpututil.IsKeyJustPressed(ebiten.KeyL) {
			tmp = "l"
		} else if inpututil.IsKeyJustPressed(ebiten.KeyM) {
			tmp = "m"
		} else if inpututil.IsKeyJustPressed(ebiten.KeyN) {
			tmp = "n"
		} else if inpututil.IsKeyJustPressed(ebiten.KeyO) {
			tmp = "o"
		} else if inpututil.IsKeyJustPressed(ebiten.KeyP) {
			tmp = "p"
		} else if inpututil.IsKeyJustPressed(ebiten.KeyQ) {
			tmp = "q"
		} else if inpututil.IsKeyJustPressed(ebiten.KeyR) {
			tmp = "r"
		} else if inpututil.IsKeyJustPressed(ebiten.KeyS) {
			tmp = "s"
		} else if inpututil.IsKeyJustPressed(ebiten.KeyT) {
			tmp = "t"
		} else if inpututil.IsKeyJustPressed(ebiten.KeyU) {
			tmp = "u"
		} else if inpututil.IsKeyJustPressed(ebiten.KeyV) {
			tmp = "v"
		} else if inpututil.IsKeyJustPressed(ebiten.KeyW) {
			tmp = "w"
		} else if inpututil.IsKeyJustPressed(ebiten.KeyX) {
			tmp = "x"
		} else if inpututil.IsKeyJustPressed(ebiten.KeyY) {
			tmp = "y"
		} else if inpututil.IsKeyJustPressed(ebiten.KeyZ) {
			tmp = "z"
		} else if inpututil.IsKeyJustPressed(ebiten.Key0) {
			tmp = "0"
		} else if inpututil.IsKeyJustPressed(ebiten.Key1) {
			tmp = "1"
		} else if inpututil.IsKeyJustPressed(ebiten.Key2) {
			tmp = "2"
		} else if inpututil.IsKeyJustPressed(ebiten.Key3) {
			tmp = "3"
		} else if inpututil.IsKeyJustPressed(ebiten.Key4) {
			tmp = "4"
		} else if inpututil.IsKeyJustPressed(ebiten.Key5) {
			tmp = "5"
		} else if inpututil.IsKeyJustPressed(ebiten.Key6) {
			tmp = "6"
		} else if inpututil.IsKeyJustPressed(ebiten.Key7) {
			tmp = "7"
		} else if inpututil.IsKeyJustPressed(ebiten.Key8) {
			tmp = "8"
		} else if inpututil.IsKeyJustPressed(ebiten.Key9) {
			tmp = "9"
		} else if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			tmp = " " // Space key
		} else if inpututil.IsKeyJustPressed(ebiten.KeyPeriod) {
			tmp = "."
		} else if inpututil.IsKeyJustPressed(ebiten.KeyComma) {
			tmp = ","
		} else if inpututil.IsKeyJustPressed(ebiten.KeySemicolon) {
			tmp = ";"
		} else if inpututil.IsKeyJustPressed(ebiten.KeyApostrophe) {
			tmp = "'"
		} else if inpututil.IsKeyJustPressed(ebiten.KeyGraveAccent) {
			tmp = "`"
		} else if inpututil.IsKeyJustPressed(ebiten.KeyMinus) {
			tmp = "-"
		} else if inpututil.IsKeyJustPressed(ebiten.KeyEqual) {
			tmp = "="
		} else if inpututil.IsKeyJustPressed(ebiten.KeyBackslash) {
			tmp = "\\"
		} else if inpututil.IsKeyJustPressed(ebiten.KeySlash) {
			tmp = "/"
		} else if inpututil.IsKeyJustPressed(ebiten.KeyLeftBracket) {
			tmp = "["
		} else if inpututil.IsKeyJustPressed(ebiten.KeyRightBracket) {
			tmp = "]"
		} else if inpututil.IsKeyJustPressed(ebiten.KeyBackspace) {
			if len(t.text) > 0 {
				t.text = t.text[:len(t.text)-1]
			}
		}
		// fmt.Println(event.Rune)
		t.text += string(tmp)

	}
	return cnd

}

// BUTTON
func (b *button) offset(f float64) {
	b.posY += int(f)
}
func (b *button) pos() (int, int) {
	return b.posX, b.posY
}
func (b *button) size() (int, int) {
	return b.sizeX, b.sizeY
}
func (b *button) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.posX), float64(b.posY))
	screen.DrawImage(b.image[b.state], op)
}

func (b *button) update(cnd, posX, posY, upperMargin, lowerMargin, offset int) int {
	x, y := ebiten.CursorPosition()
	if x > b.posX+posX && x < b.posX+posX+b.sizeX && y > posY+offset+abs(upperMargin-b.posY) && y < b.posY+offset+posY+b.sizeY-lowerMargin {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
			b.state = clicked
			return b.direction
		}
		b.state = hover
		return cnd
	}
	b.state = def
	return cnd
}

// TICKBUTTON
func (b *tickbutton) offset(f float64) {
	b.posY += int(f)
}
func (b *tickbutton) pos() (int, int) {
	return b.posX, b.posY
}
func (b *tickbutton) size() (int, int) {
	return b.sizeX, b.sizeY
}
func (b *tickbutton) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.posX), float64(b.posY))
	if !b.tick {
		screen.DrawImage(b.image[b.state], op)
	} else {
		screen.DrawImage(b.image[b.state+2], op)
	}

}

func (b *tickbutton) update(cnd, posX, posY, upperMargin, lowerMargin, offset int) int {
	x, y := ebiten.CursorPosition()
	if x > b.posX+posX && x < b.posX+posX+b.sizeX && y > posY+offset+abs(upperMargin-b.posY) && y < b.posY+offset+posY+b.sizeY-lowerMargin {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton0) {
			b.tick = !b.tick
		}
		b.state = 1
		return cnd
	}
	b.state = 0
	return cnd
}

// CONTAINER
func (c *container) offset(f float64) {
	c.posY += int(f)
}
func (c *container) pos() (int, int) {
	return c.posX, c.posY
}
func (c *container) size() (int, int) {
	return c.sizeX, c.sizeY
}
func (c *container) draw(screen *ebiten.Image) {
	c.box = ebiten.NewImage(c.sizeX, c.sizeY)
	// c.box = c.background
	c.box.DrawImage(c.background, nil)
	for _, v := range c.elements {
		v.draw(c.box)
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.posX), float64(c.posY))
	screen.DrawImage(c.box, op)

}

func (c *container) update(cnd, posX, posY, upperMargin, lowerMargin, offset int) int {
	for i := range c.elements {
		value := c.elements[i].update(cnd, c.posX+posX, c.posY+posY, upperMargin, lowerMargin, offset)
		if value != cnd {
			return value
		}
	}
	return cnd
}

// LIST
func (l *list) draw(screen *ebiten.Image) {
	l.box = ebiten.NewImage(l.sizeX, l.sizeY)
	for _, v := range l.elements {
		v.offset(l.offset)
		_, pY := v.pos()
		_, sY := v.size()
		if (pY+sY > 0) && (pY < l.sizeY) {
			v.draw(l.box)
		}
		v.offset(-l.offset)
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(l.posX), float64(l.posY))
	screen.DrawImage(l.box, op)

}

func (l *list) update(cnd int, posX, posY int) int {
	_, y := ebiten.Wheel()
	l.offset += y * 40
	if int(l.offset) < -((len(l.elements)*65+len(l.elements)*l.spacer-l.spacer)-l.sizeY)-100 {
		l.offset = float64(-((len(l.elements)*65 + len(l.elements)*l.spacer - l.spacer) - l.sizeY) - 100)
	}
	if l.offset > 0 {
		l.offset = 0
	}
	for i := range l.elements {
		_, pY := l.elements[i].pos()
		_, sY := l.elements[i].size()
		var upperMargin, lowerMargin int
		if (pY < 0) && (pY+sY > 0) {
			upperMargin = -pY
		}
		if (pY < l.sizeY) && (pY+sY > l.sizeY) {
			lowerMargin = sY - (l.sizeY - pY)

		}
		value := l.elements[i].update(cnd, l.posX+posX, l.posY+posY, upperMargin, lowerMargin, int(l.offset))
		if value != cnd {
			return value
		}
	}
	return cnd
}
