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
}

type Sprite struct {
	image *ebiten.Image
	rect  image.Rectangle
}

func NewSprite(path string) (*Sprite, error) {
	img, imgimg, err := ebitenutil.NewImageFromFile(fmt.Sprintf("assets/%s%d.png", path, i+1)) //fire1.png, fire2.png, ...
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
	return anim
}

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

func (a *AnimatedSprite) Draw(screen *ebiten.Image) {
	sprite := a.ActiveSprite()
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(sprite.rect.Min.X), float64(sprite.rect.Min.Y))
	screen.DrawImage(sprite.image, opts)
}

type Entity interface {
	Draw(*ebiten.Image)
	Update()
}
