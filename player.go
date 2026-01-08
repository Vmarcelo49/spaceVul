package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type Ship struct {
	sprites []Entity // 2 fire sprites, 1 body sprite, ?? bullet sprites
	rect    image.Rectangle
}

func MakeShip() *Ship {
	ship := &Ship{}

	shipSprite, err := NewSprite("SpaceShip")
	fire := FireAnimation()
	bArrow, err := NewSprite("bulletArrow")
	bArrowSmall, err := NewSprite("bulletArrowSmall")
	bBlueLine, err := NewSprite("bulletBlueLine")
	bRed, err := NewSprite("bulletRed")
	bBanana := BananaAnimation()
	if err != nil {
		panic(err)
	}

	ship.rect = image.Rectangle{Max: image.Point{X: (ScreenWidth / 2) - (shipSprite.rect.Dx() / 2), Y: ScreenHeight / 2}}

	ship.sprites = []Entity{fire, shipSprite, bArrow, bArrowSmall, bBlueLine, bRed, bBanana}
	return ship
}

func (s *Ship) Update() {
	for _, entt := range s.sprites {
		entt.Update()
	}
}

func (s *Ship) DrawShip(screen *ebiten.Image) {
	fireCoord := []Vector2{{11, 27}, {20, 27}}

	//bulletInitCoord := []Vector2{{6, 1}, Vector2{25, 1}}

	ship := Vector2{float64(s.rect.Dx()), float64(s.rect.Dy())}

	opts := &ebiten.DrawImageOptions{}

	fireSpr := s.sprites[0].Image()
	opts.GeoM.Translate(ship[0]+fireCoord[0][0]-float64(fireSpr.Bounds().Dx()/2), ship[1]+fireCoord[0][1])
	screen.DrawImage(fireSpr, opts)
	opts.GeoM.Reset()

	opts.GeoM.Translate(ship[0]+fireCoord[1][0]-float64(fireSpr.Bounds().Dx()/2), ship[1]+fireCoord[1][1])
	screen.DrawImage(fireSpr, opts)
	opts.GeoM.Reset()

	opts.GeoM.Translate(ship[0], ship[1])
	screen.DrawImage(s.sprites[1].Image(), opts)

}
