package main

import (
	"github.com/joho/godotenv"
)

func main() {
	championAssetURL, summonerAsset, err := initializeLOLAssets()
	if err != nil {
		panic(err)
	}

	RenderDisplayScreen(championAssetURL, summonerAsset)
}

func initializeLOLAssets() (map[int]string, map[int]SummonerInfo, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, nil, err
	}

	championAssetURL, err := FetchChampionAssetURL()
	if err != nil {
		return nil, nil, err
	}

	summonerAsset, err := FetchSummonerSpellsAssetURL()
	if err != nil {
		return nil, nil, err
	}

	return championAssetURL, summonerAsset, nil
}
