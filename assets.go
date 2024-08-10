package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// assetData is a private struct that represents the structure of the JSON data
type assetData struct {
	Data map[string]asset `json:"data"`
}

// asset is a private struct that holds the key, image, and cooldown information
type asset struct {
	Key      string    `json:"key"`
	Image    Image     `json:"image"`
	Cooldown []float64 `json:"cooldown,omitempty"` // Cooldown field is optional and will be omitted if not present
}

// image is a private struct that holds the image URL information
type Image struct {
	ImageUrl string `json:"full"`
}

const (
	baseURL              = "https://ddragon.leagueoflegends.com/cdn/"
	dataLocale           = "/data/en_US"
	championAssetPartURL = "/img/champion/"
	summonerAssetPartURL = "/img/spell/"
	championURL          = "/champion.json"
	summonerURL          = "/summoner.json"
)

// SummonerInfo struct contains the information for summoner spells
type SummonerInfo struct {
	Cooldown float64
	URL      string
}

// FetchChampionAssetURL is a public function to fetch champion assets and returns map[int]string
func FetchChampionAssetURL() (map[int]string, error) {
	return fetchAssetURL(championURL, func(a asset) string {
		lolVersion := os.Getenv("LOLVERSION")
		return fmt.Sprint(baseURL, lolVersion, championAssetPartURL, a.Image.ImageUrl)
	})
}

// FetchSummonerSpellsAssetURL is a public function to fetch summoner spell assets and returns map[int]SummonerInfo
func FetchSummonerSpellsAssetURL() (map[int]SummonerInfo, error) {
	return fetchAssetURL(summonerURL, func(a asset) SummonerInfo {
		lolVersion := os.Getenv("LOLVERSION")
		return SummonerInfo{
			Cooldown: a.Cooldown[0], // Fetches the first element in the Cooldown array
			URL:      fmt.Sprint(baseURL, lolVersion, summonerAssetPartURL, a.Image.ImageUrl),
		}
	})
}

// fetchAssetURL is a private function that fetches and processes data based on the provided processing function
func fetchAssetURL[T any](dataURL string, process func(asset) T) (map[int]T, error) {
	lolVersion := os.Getenv("LOLVERSION")
	url := fmt.Sprint(baseURL, lolVersion, dataLocale, dataURL)

	jsonData, err := FetchJSON(url, nil)
	if err != nil {
		return nil, err
	}

	var data assetData
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, err
	}

	result := make(map[int]T)
	for _, a := range data.Data {
		num, _ := strconv.Atoi(a.Key)
		result[num] = process(a)
	}

	return result, nil
}
