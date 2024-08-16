package main

import (
	"image"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

// Define a custom type for the enumeration
type GameState int

const (
	Inactive = iota
	Found
	Active
)

// Game is an empty struct that represents your game state.
type Game struct {
	originalBackground     *ebiten.Image
	scaledBackground       *ebiten.Image
	activeGameScreen       *ebiten.Image
	scaledActiveGameScreen *ebiten.Image
	championAsset          map[int]image.Image
	summonerAsset          map[int]SummonerInfo
	apiRequestChannel      chan bool
	apiResponseChannel     chan []EnemyAssets
	apiPeriodicTimer       *time.Timer
	GameState              GameState
	summonerButtons        []*SummonerButton
}

func RenderDisplayScreen(championAsset map[int]image.Image, summonerAsset map[int]SummonerInfo) {

	ebiten.SetFullscreen(true)

	// Create a new Game instance with the background image
	game := &Game{
		originalBackground: LoadEbitenImage("nightsky.png"),
		activeGameScreen:   LoadEbitenImage("activeGame.png"),
		championAsset:      championAsset,
		summonerAsset:      summonerAsset,
		apiRequestChannel:  make(chan bool, 1),
		apiResponseChannel: make(chan []EnemyAssets),
		GameState:          Inactive,
	}

	game.apiRequestChannel <- true

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func (g *Game) Update() error {

	select {
	case response := <-g.apiResponseChannel:
		_ = response
		if g.GameState == Found {
			g.SetSummonerButtonsToAsset(response)
			g.GameState = Active
		}

		g.apiPeriodicTimer = time.AfterFunc(10*time.Second, func() {
			g.apiRequestChannel <- true
		})
	case <-g.apiRequestChannel:
		go ObtainEnemyAssets(g)
	default:
		// Continue running the game without blocking
	}

	if g.GameState == Active {
		listenToLeftClicksOnScreen(g)
	}

	// Update buttons (cooldown timers)
	for _, button := range g.summonerButtons {
		button.summonerSpell1.Update()
		button.summonerSpell2.Update()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	if g.GameState == Active {
		if g.scaledActiveGameScreen != nil {
			screen.DrawImage(g.scaledActiveGameScreen, nil)

			for _, summonerButton := range g.summonerButtons {
				summonerButton.Draw(screen)
			}
		}
	} else {
		if g.scaledBackground != nil {
			screen.DrawImage(g.scaledBackground, nil)
		}
	}
}

func (g *Game) SetSummonerButtonsToAsset(enemyAssetSlice []EnemyAssets) {
	for i, enemyAsset := range enemyAssetSlice {
		summonerButton := g.summonerButtons[i]

		// Process champion portrait
		championIcon := createScaledImage(summonerButton.championPortrait.width, summonerButton.championPortrait.height, ebiten.NewImageFromImage(enemyAsset.ChampionAsset))
		summonerButton.championPortrait.image = championIcon
		summonerButton.championPortrait.originalImage = championIcon

		// Process summoner spell 1
		summoner1Icon := createScaledImage(summonerButton.summonerSpell1.width, summonerButton.summonerSpell1.height, ebiten.NewImageFromImage(enemyAsset.SummonerSpell1Info.Icon))
		summonerButton.summonerSpell1.image = summoner1Icon
		summonerButton.summonerSpell1.originalImage = summoner1Icon
		summonerButton.summonerSpell1.cooldown = int(enemyAsset.SummonerSpell1Info.Cooldown)
		summonerButton.summonerSpell1.cooldownTimer = int(enemyAsset.SummonerSpell1Info.Cooldown)

		// Process summoner spell 2
		summoner2Icon := createScaledImage(summonerButton.summonerSpell2.width, summonerButton.summonerSpell2.height, ebiten.NewImageFromImage(enemyAsset.SummonerSpell2Info.Icon))
		summonerButton.summonerSpell2.image = summoner2Icon
		summonerButton.summonerSpell2.originalImage = summoner2Icon
		summonerButton.summonerSpell2.cooldown = int(enemyAsset.SummonerSpell2Info.Cooldown)
		summonerButton.summonerSpell2.cooldownTimer = int(enemyAsset.SummonerSpell2Info.Cooldown)
	}
}

// Helper function to create and scale an image
func createScaledImage(targetWidth, targetHeight int, sourceImage *ebiten.Image) *ebiten.Image {
	// Create a new image with target dimensions
	scaledImage := ebiten.NewImage(targetWidth, targetHeight)

	// Calculate scaling factors
	scaleX := float64(targetWidth) / float64(sourceImage.Bounds().Dx())
	scaleY := float64(targetHeight) / float64(sourceImage.Bounds().Dy())

	// Create DrawImageOptions with scaling
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scaleX, scaleY)

	// Draw the source image onto the new scaled image
	scaledImage.DrawImage(sourceImage, op)

	return scaledImage
}

func listenToLeftClicksOnScreen(g *Game) {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		for _, button := range g.summonerButtons {
			checkAndHandleClick(x, y, &button.summonerSpell1)
			checkAndHandleClick(x, y, &button.summonerSpell2)
		}
	}
}

func checkAndHandleClick(x, y int, spell *Button) {
	// Check if within bounds of a button
	if x >= spell.x && x <= spell.x+spell.width && y >= spell.y && y <= spell.y+spell.height {
		if !spell.isCoolingDown {
			spell.isCoolingDown = true
		}
	}
}
