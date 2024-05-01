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

func (t *textbox) update(cnd, posX, posY, upperMargin, lowerMargin int) int {
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

// PosY+LowerMargin+b.PosY+
func (b *button) update(cnd, posX, posY, upperMargin, lowerMargin int) int {
	x, y := ebiten.CursorPosition()
	if x > b.posX+posX && x < b.posX+posX+b.sizeX && y > posY+abs(upperMargin-b.posY) && y < b.posY+posY+b.sizeY-lowerMargin {
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

// PosY+LowerMargin+b.PosY+
func (b *tickbutton) update(cnd, posX, posY, upperMargin, lowerMargin int) int {
	x, y := ebiten.CursorPosition()
	if x > b.posX+posX && x < b.posX+posX+b.sizeX && y > posY+abs(upperMargin-b.posY) && y < b.posY+posY+b.sizeY-lowerMargin {
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

func (c *container) update(cnd, posX, posY, upperMargin, lowerMargin int) int {
	for i := range c.elements {
		value := c.elements[i].update(cnd, c.posX+posX, c.posY+posY, upperMargin, lowerMargin)
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

	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(l.posX), float64(l.posY))
	screen.DrawImage(l.box, op)

}

func (l *list) update(cnd int, posX, posY int) int {
	_, y := ebiten.Wheel()
	l.offset = y * 40
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
		value := l.elements[i].update(cnd, l.posX+posX, l.posY+posY, upperMargin, lowerMargin)
		if value != cnd {
			return value
		}
	}
	return cnd
}
