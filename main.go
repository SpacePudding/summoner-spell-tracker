package main

import (
	"fmt"
	"time"

	"github.com/joho/godotenv"
)

type SummonerAssets struct {
}

// type OngoingMatch struct {
// 	GameId int `json:"gameId"`
// }

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// I want to render "No active game currently found"
	// renderInactiveScreen()

	// I want to fetch the asset of summoners from a specific endpoint to keep things up to date.
	// summonerAssets, err = fetchAssets()

	for {

		// Use API call to fetch the game. The return from that call should be the five enemy champions and their png

		ongoingMatch, err := obtainEnemyTeamAssets()

		if err == nil {

			// Otherwise if I get HTTP Status 200 I want to render the screen I want.
			// renderScreen(summonerAssets, enemyChampions)
			fmt.Println(ongoingMatch)

		}

		fmt.Println("Waiting for 10 seconds...")
		time.Sleep(10 * time.Second)
		fmt.Println("Done waiting!")
	}
}

// func fetchAssets() (SummonerAssets, error) {

// }
