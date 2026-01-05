package main

import (
	"github.com/ebitengine/debugui"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	debugui  debugui.DebugUI
	ship     *Ship
	entities []Entity
}

func (g *Game) Update() error {
	for _, entt := range g.entities {
		entt.Update()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, entt := range g.entities {
		for layer := range g.entities.la
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 256, 240
}

func main() {
	game := &Game{ship: MakeShip()}
	ebiten.SetWindowSize(256*4, 240*4)
	game.entities = game.ship.sprites
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
