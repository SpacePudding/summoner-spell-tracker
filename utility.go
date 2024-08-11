package main

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

// FetchJSON fetches the JSON data from the URL
func FetchJSON(url string, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: %s", resp.Status)
	}

	return io.ReadAll(resp.Body)
}

func FetchPng(url string, client *http.Client) image.Image {

	resp, err := client.Get(url)
	if err != nil {
		log.Fatalf("Failed to fetch image: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to fetch image: %v", resp.Status)
	}

	// Decode the PNG image
	img, err := png.Decode(resp.Body)
	if err != nil {
		log.Fatalf("Failed to decode image: %v", err)
	}

	return img
}

func LoadEbitenImage(url string) *ebiten.Image {
	image, err := loadImage(url)
	if err != nil {
		log.Fatalf("Failed to load image: %v", err)
	}

	return ebiten.NewImageFromImage(image)
}

func loadImage(filename string) (image.Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	img, err := png.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}
