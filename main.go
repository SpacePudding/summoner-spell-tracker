package main

import (
	"fmt"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// I want to render "No active game currently found"
	// renderInactiveScreen()

	championAssetURL, _ := fetchChampionAssetURL()
	_ = championAssetURL

	// summoner, _ := fetchSummonerInfoAndAssetURL()

	for {

		// Use API call to fetch the game. The return from that call should be the five enemy champions and their png
		enemyData, err := obtainEnemyTeamAssets()

		if err == nil {

			// Otherwise if I get HTTP Status 200 I want to render the screen I want.
			// renderScreen(summonerAssets, enemyChampions)
			fmt.Println(enemyData)

		}

		fmt.Println("Waiting for 10 seconds...")
		time.Sleep(10 * time.Second)
		fmt.Println("Done waiting!")
	}
}
