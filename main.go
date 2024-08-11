package main

import (
	"fmt"
	"image"
	"net/http"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	start := time.Now() // Record the start time
	championAssetURL, summonerAsset, err := initializeLOLAssets()
	if err != nil {
		panic(err)
	}

	elapsed := time.Since(start) // Calculate the elapsed time

	fmt.Printf("initializeLOLAssets took %s\n", elapsed)

	RenderDisplayScreen(championAssetURL, summonerAsset)
}

func initializeLOLAssets() (map[int]image.Image, map[int]SummonerInfo, error) {
	if err := godotenv.Load(); err != nil {
		return nil, nil, err
	}

	var championAssetURL map[int]image.Image
	var summonerAsset map[int]SummonerInfo
	errCh := make(chan error, 2)
	client := &http.Client{}

	go func() {
		var err error
		championAssetURL, err = FetchChampionAssetURL(client)
		errCh <- err
	}()

	go func() {
		var err error
		summonerAsset, err = FetchSummonerSpellsAssetURL(client)
		errCh <- err
	}()

	// Wait for both goroutines to finish and check for errors
	if err := <-errCh; err != nil {
		return nil, nil, err
	}
	if err := <-errCh; err != nil {
		return nil, nil, err
	}

	return championAssetURL, summonerAsset, nil
}
