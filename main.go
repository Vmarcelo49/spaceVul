package main

import (
	"fmt"
	"image"

	"github.com/ebitengine/debugui"
	"github.com/hajimehoshi/ebiten/v2"
)

var ScreenWidth = 256
var ScreenHeight = 240

type Game struct {
	debugui       debugui.DebugUI
	ship          *Ship
	otherEntities []Entity
}

func (g *Game) Update() error {
	for _, entt := range g.otherEntities {
		entt.Update()
	}
	g.Input()
	g.ship.Update()
	if _, err := g.debugui.Update(func(ctx *debugui.Context) error {
		ctx.Window("info", image.Rect(0, 0, ScreenWidth/2, ScreenHeight/2), func(layout debugui.ContainerLayout) {
			ctx.Text(fmt.Sprint("X:", g.ship.rect.Dx()))
			ctx.Text(fmt.Sprint("Y:", g.ship.rect.Dy()))
		})
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	DrawWITHLayers(screen, g.otherEntities)
	g.ship.DrawShip(screen)

	g.debugui.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenHeight
}

func main() {
	game := &Game{ship: MakeShip()}
	ebiten.SetWindowSize(256*4, 240*4)
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
