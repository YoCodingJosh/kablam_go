package core

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	json "github.com/goccy/go-json"
)

type AssetManager struct{
	Images map[string]*ebiten.Image
	imagesMu sync.Mutex

	// TODO: other asset types
}

func NewAssetManager() *AssetManager {
	return &AssetManager{
		Images: make(map[string]*ebiten.Image),
	}
}

func (am *AssetManager) LoadFromJSON(path string) error {
	// Open the JSON file
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open file %q: %w", path, err)
	}
	defer file.Close()

	// Read the file content
	fileContent, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("failed to read file %q: %w", path, err)
	}

	// Temporary structure to hold the parsed JSON
	parsedData := struct {
		Images map[string]string `json:"images"`
	}{}

	// Parse the JSON file into the temporary structure
	err = json.Unmarshal(fileContent, &parsedData)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON from file %q: %w", path, err)
	}

	// Create wait group to wait for all images to load
	var wg sync.WaitGroup
	errorCh := make(chan error)
	done := make(chan struct{})

	// Error handling goroutine
	go func() {
    var err error
    for e := range errorCh { // Reads errors from the channel
        if err == nil {      // Collect only the first error
            err = e
        }
    }
    done <- struct{}{} // Signal completion
	}()

	wg.Add(1) // TODO: Increment for other asset types (ie: sounds, fonts, etc)

	// Image loading goroutine
	go func() {
		defer wg.Done()

		// Load the images into the AssetManager
		// TODO: investigate if we can load images concurrently
		for key, value := range parsedData.Images {
			image, err := am.LoadImage(fmt.Sprintf("resources/images/%s", value))

			if err != nil {
					errorCh <- fmt.Errorf("failed to load image %q: %w", value, err)
					continue
			}

			am.imagesMu.Lock()

			am.Images[key] = image

			am.imagesMu.Unlock()
		}
	}()

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

	return img, nil
}
