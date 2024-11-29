package core

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type AssetManager struct{
	images map[string]*ebiten.Image
}

func NewAssetManager() *AssetManager {
	return &AssetManager{
		images: make(map[string]*ebiten.Image),
	}
}

// TODO: Load a JSON file and use goroutines to load stuff in parallel

func (am *AssetManager) LoadImage(path string) (*ebiten.Image, error) {
	if img, ok := am.images[path]; ok {
		return img, nil
	}

	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		return nil, err
	}

	am.images[path] = img

	return img, nil
}
