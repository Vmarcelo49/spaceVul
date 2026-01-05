package main

type Layer int

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
