package main

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

var speed int = 2

func (g *Game) Input() {
	shipPoint := &g.ship.rect.Max
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		shipPoint.X -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		shipPoint.X += speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		shipPoint.Y += speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		shipPoint.Y -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.shootProjectile()
	}
	g.checkShipOutofBounds()
}

func (g *Game) shootProjectile() {
	bulletInitCoord := []Vector2{{6, 1}, {25, 1}}
	bulletCoord := Vector2{float64(g.ship.rect.Max.X + int(bulletInitCoord[0][0])), float64(g.ship.rect.Max.Y + int(bulletInitCoord[0][1]))}
	bulletCoord2 := Vector2{float64(g.ship.rect.Max.X + int(bulletInitCoord[1][0])), float64(g.ship.rect.Max.Y + int(bulletInitCoord[1][1]))}

	g.spawnProjectile(g.ship.ProjectileType, bulletCoord)
	g.spawnProjectile(g.ship.ProjectileType, bulletCoord2)
}

func (g *Game) spawnProjectile(pType ProjectileType, coord Vector2) {
	bullet := makeProjectile(pType, image.Point{int(coord[0]), int(coord[1])})
	g.otherEntities = append(g.otherEntities, &bullet)
}

func (g *Game) checkShipOutofBounds() {
	shipPoint := &g.ship.rect.Max
	shipSize := g.ship.rect.Min

	maxPointRight := ScreenWidth
	maxPointLeft := shipSize.X
	maxPointDown := ScreenHeight - shipSize.Y/2

	if shipPoint.X > maxPointRight {
		shipPoint.X = maxPointRight
	}
	if shipPoint.X < maxPointLeft {
		shipPoint.X = maxPointLeft
	}
	if shipPoint.Y > maxPointDown {
		shipPoint.Y = maxPointDown
	}
	if shipPoint.Y < shipSize.Y {
		shipPoint.Y = shipSize.Y
	}
}
