package main

import (
	"image"
	_ "image/png"
	"os"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
)

type ImageCache struct {
	mu     sync.RWMutex
	images map[string]*ebiten.Image
}

var GlobalImageCache = NewImageCache()

func NewImageCache() *ImageCache {
	return &ImageCache{
		images: make(map[string]*ebiten.Image),
	}
}

func (ic *ImageCache) Load(path string) (*ebiten.Image, error) {
	ic.mu.Lock()
	defer ic.mu.Unlock()

	if img, exists := ic.images[path]; exists {
		return img, nil
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	ebitenImg := ebiten.NewImageFromImage(img)
	ic.images[path] = ebitenImg

	return ebitenImg, nil
}
