package main

import "github.com/hajimehoshi/ebiten/v2"

func (a *AnimatedSprite) Draw(screen *ebiten.Image) {
	sprite := a.ActiveSprite()
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(sprite.rect.Min.X), float64(sprite.rect.Min.Y))
	screen.DrawImage(sprite.image, opts)
}

func (s *Sprite) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(s.rect.Min.X), float64(s.rect.Min.Y))
	screen.DrawImage(s.image, opts)
}
