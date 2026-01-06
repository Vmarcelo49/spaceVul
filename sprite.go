package main

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	_ "image/png" // png decoder
)

type AnimatedSprite struct {
	sprites         []*Sprite
	spriteDurations []int
	FrameTimeLeft   int
	totalDuration   int
	frameIndex      int

	layer Layer
}

type Sprite struct {
	image *ebiten.Image
	rect  image.Rectangle

	layer Layer
}

func NewSprite(path string) (*Sprite, error) {
	img, imgimg, err := ebitenutil.NewImageFromFile(fmt.Sprintf("assets/%s.png", path)) //fire1.png, fire2.png, ...
	if err != nil {
		return nil, err
	}
	spr := &Sprite{
		image: img,
		rect:  imgimg.Bounds(),
	}

	return spr, nil
}

type Vector2 [2]float64

func LoadAnimatedSprite(path string, numOfImgs int) (*AnimatedSprite, error) {
	anim := &AnimatedSprite{}

	for i := range numOfImgs {
		img, imgimg, err := ebitenutil.NewImageFromFile(fmt.Sprintf("assets/%s%d.png", path, i+1)) //fire1.png, fire2.png, ...
		if err != nil {
			return nil, err
		}
		spr := &Sprite{
			image: img,
			rect:  imgimg.Bounds(),
		}
		anim.sprites = append(anim.sprites, spr)
	}
	return anim, nil

}

func FireAnimation() *AnimatedSprite {
	anim, err := LoadAnimatedSprite("fire", 3)
	if err != nil {
		panic(err)
	}
	duration := 6 // 100ms in 60fps
	anim.spriteDurations = make([]int, 3)
	for i := range 3 {
		anim.spriteDurations[i] = duration
	}
	anim.totalDuration = 3 * 6
	anim.layer = LayerShipFire
	return anim
}

func isPair(num int) bool {
	return num%2 == 0
}

func BananaAnimation() *AnimatedSprite {
	const animLen = 8
	anim, err := LoadAnimatedSprite("banana", animLen)
	if err != nil {
		panic(err)
	}
	duration1 := 6  // 100ms
	duration2 := 12 // 200ms
	anim.spriteDurations = make([]int, animLen)
	for i := range animLen {
		if isPair(i + 1) {
			anim.spriteDurations[i] = duration1
		} else {
			anim.spriteDurations[i] = duration2
		}
	}
	anim.totalDuration = (4 * 6) + (4 * 12)
	anim.layer = LayerBullet

	return anim
}

func (s *Sprite) Update() {}

func (a *AnimatedSprite) Update() {
	a.FrameTimeLeft--
	if a.FrameTimeLeft > 0 {
		return // we didnt finish this sprite, so dont advance
	}
	a.frameIndex++

	if a.frameIndex >= len(a.sprites) {
		a.frameIndex = 0
	}
	a.FrameTimeLeft = a.spriteDurations[a.frameIndex]
}

func (a *AnimatedSprite) ActiveSprite() *Sprite {
	return a.sprites[a.frameIndex]
}

type Entity interface {
	Draw(*ebiten.Image)
	Update()
	Layer() *Layer
}
