package main

import (
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// Game is an empty struct that represents your game state.
type Game struct {
	originalBackground     *ebiten.Image
	scaledBackground       *ebiten.Image
	activeGameScreen       *ebiten.Image
	scaledActiveGameScreen *ebiten.Image
	championAssetURL       map[int]string
	summonerAsset          map[int]SummonerInfo
	apiRequestChannel      chan bool
	apiResponseChannel     chan []EnemyAssets
	apiPeriodicTimer       *time.Timer
	enemyAssets            []EnemyAssets
	isGameActive           bool
}

func RenderDisplayScreen(championAssetURL map[int]string, summonerAsset map[int]SummonerInfo) {

	ebiten.SetFullscreen(true)

	// Create a new Game instance with the background image
	game := &Game{
		originalBackground: LoadEbitenImage("nightsky.png"),
		activeGameScreen:   LoadEbitenImage("activeGame.png"),
		championAssetURL:   championAssetURL,
		summonerAsset:      summonerAsset,
		apiRequestChannel:  make(chan bool, 1),
		apiResponseChannel: make(chan []EnemyAssets),
		isGameActive:       false,
	}

	game.apiRequestChannel <- true

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Update() error {

	select {
	case response := <-g.apiResponseChannel:
		g.enemyAssets = response

		if g.isGameActive {
			// Do something
		}

		g.apiPeriodicTimer = time.AfterFunc(10*time.Second, func() {
			g.apiRequestChannel <- true
		})
	case <-g.apiRequestChannel:
		go ObtainEnemyAssets(g)
	default:
		// Continue running the game without blocking
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	if g.isGameActive {
		if g.scaledActiveGameScreen != nil {
			screen.DrawImage(g.scaledActiveGameScreen, nil)
		}
	} else {
		if g.scaledBackground != nil {
			screen.DrawImage(g.scaledBackground, nil)
		}
	}
}
