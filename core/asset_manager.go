package core

import (
	"io"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	json "github.com/goccy/go-json"
)

type AssetManager struct{
	Images map[string]*ebiten.Image `json:"images"`
}

func NewAssetManager() *AssetManager {
	return &AssetManager{
		Images: make(map[string]*ebiten.Image),
	}
}

func (am *AssetManager) LoadFromJSON(path string) error {
	// Load up the JSON file
	file, err := os.Open(path)
	if err != nil {
			log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	// Actually read the file
	fileContent, err := io.ReadAll(file)
	if err != nil {
			log.Fatalf("Failed to read file: %v", err)
	}

	// Structure to hold the marshalled JSON
	marshalledJson := struct {
		Images map[string]string `json:"images"`
	}{}

	// Parse the JSON file
	err = json.Unmarshal(fileContent, &marshalledJson)
	if err != nil {
		return err
	}

	// TODO: Actually use goroutines to load everything in parallel

	// print the images for now
	for k, v := range marshalledJson.Images {
		log.Printf("Key: %s, Value: %v", k, v)
	}

	return nil
}

func (am *AssetManager) LoadImage(path string) (*ebiten.Image, error) {
	if img, ok := am.Images[path]; ok {
		return img, nil
	}

	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		return nil, err
	}

	am.Images[path] = img

	return img, nil
}
