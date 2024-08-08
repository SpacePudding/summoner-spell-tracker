package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type ChampionData struct {
	Data map[string]Champion `json:"data"`
}

type Champion struct {
	Key   string `json:"key"`
	Image Image  `json:"image"`
}

type Image struct {
	ImageUrl string `json:"full"`
}

const (
	baseURL  = "https://ddragon.leagueoflegends.com/cdn/"
	locale   = "en_US"
	endpoint = "/data/" + locale + "/champion.json"
)

// CreateMapping creates a mapping from champion IDs to their names
func fetchChampionAssetURL() (map[int]string, error) {

	lolVersion := os.Getenv("LOLVERSION")
	url := fmt.Sprint(baseURL, lolVersion, endpoint)

	jsonData, err := FetchJSON(url, nil)
	if err != nil {
		return nil, err
	}

	var championData ChampionData

	err = json.Unmarshal(jsonData, &championData)
	if err != nil {
		return nil, err
	}

	idToUrl := make(map[int]string)
	for _, champ := range championData.Data {
		num, _ := strconv.Atoi(champ.Key)
		idToUrl[num] = champ.Image.ImageUrl
	}

	return idToUrl, nil
}

// // GetChampionName gets the champion name by ID
// func getChampionName(idToName map[int]string, championID int) string {
// 	if name, ok := idToName[championID]; ok {
// 		return name
// 	}
// 	return "Champion ID not found"
// }

// func fetchSummonerInfoAndAssetURL() error {
// 	return nil
// }
