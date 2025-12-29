package main

import "image"

type Ship struct {
	sprites []Entity // 2 fire sprites, 1 body sprite, ?? bullet sprites
	rect    image.Rectangle
}

func MakeShip() *Ship {
	ship := &Ship{}

	fire := FireAnimation()
	shipSprite, err := NewSprite("SpaceShip")
	if err != nil {
		panic(err)
	}

	bArrow, err := NewSprite("bulletArrow")
	bArrowSmall, err := NewSprite("bulletArrowSmall")
	bBlueLine, err := NewSprite("bulletBlueLine")
	bRed, err := NewSprite("bulletRed")
	bBanana := BananaAnimation()

	ship.sprites = []Entity{fire,shipSprite,bArrow,bArrowSmall,bBlueLine,bRed,bBanana]}
	return ship
}
