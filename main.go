package main

import (
	"github.com/ebitengine/debugui"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	debugui       debugui.DebugUI
	ship          *Ship
	otherEntities []Entity
}

func (g *Game) Update() error {
	for _, entt := range g.otherEntities {
		entt.Update()
	}
	g.ship.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	DrawWITHLayers(screen, g.otherEntities)
	DrawWITHLayers(screen, g.ship.sprites)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 256, 240
}

func main() {
	game := &Game{ship: MakeShip()}
	ebiten.SetWindowSize(256*4, 240*4)
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
