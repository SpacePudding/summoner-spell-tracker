package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

type OngoingMatch struct {
	Participants []Participants `json:"participants"`
}

type Participants struct {
	PuuId      string `json:"puuid"`
	TeamId     int    `json:"teamId"`
	Spell1Id   int    `json:"spell1Id"`
	Spell2Id   int    `json:"spell2Id"`
	ChampionId int    `json:"championId"`
}

type EnemyData struct {
	Spell1Id   int
	Spell2Id   int
	ChampionId int
}

func obtainEnemyTeamAssets() ([]EnemyData, error) {
	ongoingMatch, err := obtainOngoingMatch()
	if err != nil {
		return []EnemyData{}, err
	}
	return extractEnemyData(ongoingMatch.Participants)
}

func obtainOngoingMatch() (OngoingMatch, error) {
	encryptedPUUID := os.Getenv("PUUID")
	urlEndpoint := fmt.Sprintf("https://euw1.api.riotgames.com/lol/spectator/v5/active-games/by-summoner/%s", encryptedPUUID)

	// Create a GET request
	req, err := http.NewRequest("GET", urlEndpoint, nil)
	if err != nil {
		panic(err)
	}

	// Add headers if required (e.g., API key)
	req.Header.Set("X-Riot-Token", os.Getenv("API_KEY")) // Replace with actual header name if needed

	// Send HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return OngoingMatch{}, errors.New("game not found")
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Parse the JSON response into a struct
	var ongoingMatch OngoingMatch
	err = json.Unmarshal(body, &ongoingMatch)
	if err != nil {
		panic(err)
	}

	return ongoingMatch, nil
}

func extractEnemyData(participants []Participants) ([]EnemyData, error) {
	enemyTeamId, err := determineEnemyTeamId(participants)
	if err != nil {
		return []EnemyData{}, err
	}

	enemyDataArr := []EnemyData{}

	for i := 0; i < len(participants); i++ {
		if participants[i].TeamId == enemyTeamId {
			enemyData := EnemyData{
				Spell1Id:   participants[i].Spell1Id,
				Spell2Id:   participants[i].Spell2Id,
				ChampionId: participants[i].ChampionId,
			}
			enemyDataArr = append(enemyDataArr, enemyData)
		}
	}
	return enemyDataArr, nil
}

func determineEnemyTeamId(participants []Participants) (int, error) {
	const blueSide = 100
	const redSide = 200

	for i := 0; i < len(participants); i++ {
		if participants[i].PuuId == os.Getenv("PUUID") {
			allyTeamId := participants[i].TeamId

			if allyTeamId == blueSide {
				return redSide, nil
			} else if allyTeamId == redSide {
				return blueSide, nil
			} else {
				return 0, fmt.Errorf("unknown teamId %d", allyTeamId)
			}
		}
	}

	return 0, errors.New("your PUUID couldn't be found in the active game")
}
