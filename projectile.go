package main

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type ProjectileType int

const (
	pBanana ProjectileType = iota
	pArrow
	pArrowS
	pBlue
	pRed
)

type Projectile struct {
	pType      ProjectileType
	rect       image.Rectangle
	damage     int
	speed      int
	cooldown   int
	isAnimated bool
}

func makeProjectile(pType ProjectileType, coords image.Point) Projectile {
	p := Projectile{
		pType: pType,
		rect:  image.Rectangle{Max: coords},
	}
	switch pType {
	case pBanana:
		p.configProjectile(9, 9, 3, 1, 6)
		p.isAnimated = true
	case pArrowS:
		p.configProjectile(3, 2, 1, 3, 3)
	case pArrow:
		p.configProjectile(5, 5, 1, 3, 3)
	case pBlue:
		p.configProjectile(3, 5, 1, 6, 1)
	case pRed:
		p.configProjectile(5, 5, 5, 1, 1)
	}

	return p

}

func (p *Projectile) configProjectile(sizeW, sizeH, dmg, speed, cd int) {
	p.rect.Min = image.Point{sizeW, sizeH}
	p.damage = dmg
	p.speed = speed
	p.cooldown = cd
}

func (p *Projectile) Draw(screen *ebiten.Image) {
	pImage, err := GlobalImageCache.Load(p.pType.String()) // shits wrong mann
	if err != nil {
		panic(fmt.Sprintf("error while trying to load an image %s", err))
	}
	pOpts := &ebiten.DrawImageOptions{}
	pOpts.GeoM.Translate(float64(p.rect.Min.X), float64(p.rect.Min.Y))
	screen.DrawImage(pImage, pOpts)

}
func (p *Projectile) Update() {
	coords := Coords(p.rect)

}
func (p *Projectile) Layer() *Layer {
	layer := LayerBullet
	return &layer
}
func (p *Projectile) Image() *ebiten.Image {
	return nil
}
