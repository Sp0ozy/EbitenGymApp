package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// BUTTON
func (b *button) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.posX), float64(b.posY))
	screen.DrawImage(b.image[b.state], op)
}
func (b *button) update(cnd, posX, posY int) int {
	x, y := ebiten.CursorPosition()
	if x > b.posX+posX && x < b.posX+posX+b.sizeX && y > b.posY+posY && y < b.posY+posY+b.sizeY {
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

// CONTAINER
func (c *container) draw(screen *ebiten.Image) {
	c.box = ebiten.NewImage(c.sizeX, c.sizeY)
	for _, v := range c.elements {
		v.draw(c.box)
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.posX), float64(c.posY))
	screen.DrawImage(c.box, op)

}

func (c *container) update(cnd int, posX, posY int) int {
	for _, v := range c.elements {
		if value := v.update(cnd, c.posX+posX, c.posY+posY); value != cnd {
			return value
		}
	}
	return cnd
}
