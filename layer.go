package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Layer int

const (
	MaxLayer = 4
)

const (
	LayerBG Layer = iota
	LayerShipFire
	LayerShipBody
	LayerBullet
)

func (s *Sprite) Layer() *Layer {
	return &s.layer
}

func (s *AnimatedSprite) Layer() *Layer {
	return &s.layer
}

func DrawWITHLayers(screen *ebiten.Image, entities []Entity) {
	layers := [MaxLayer][]Entity{}
	populateLayers(&layers, entities)
	for i := range layers {
		for _, renderable := range layers[i] {
			renderable.Draw(screen)
		}
	}
}

func populateLayers(layers *[MaxLayer][]Entity, entities []Entity) {
	for _, ent := range entities {
		if ent == nil {
			continue
		}

		layerPtr := ent.Layer()
		if layerPtr == nil {
			continue
		}

		idx := int(*layerPtr)
		layers[idx] = append(layers[idx], ent)
	}
}
