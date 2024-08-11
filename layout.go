package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {

	if g.scaledBackground == nil || g.scaledActiveGameScreen == nil {
		g.scaledBackground, g.scaledActiveGameScreen = createScaledImages(outsideWidth, outsideHeight, g.originalBackground, g.activeGameScreen)
		g.summonerButtons = initializeSummonerButtons(outsideWidth, outsideHeight)
	}

	return outsideWidth, outsideHeight
}

func createScaledImages(outsideWidth, outsideHeight int, originalBackground, activeGameScreen *ebiten.Image) (*ebiten.Image, *ebiten.Image) {
	scaledBackground := ebiten.NewImage(outsideWidth, outsideHeight)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(float64(outsideWidth)/float64(originalBackground.Bounds().Dx()), float64(outsideHeight)/float64(originalBackground.Bounds().Dy()))
	scaledBackground.DrawImage(originalBackground, op)

	// Scale activeGameScreen to fit the screen
	scaledActiveGameScreen := ebiten.NewImage(outsideWidth, outsideHeight)
	op1 := &ebiten.DrawImageOptions{}
	op1.GeoM.Scale(float64(outsideWidth)/float64(activeGameScreen.Bounds().Dx()), float64(outsideHeight)/float64(activeGameScreen.Bounds().Dy()))
	scaledActiveGameScreen.DrawImage(activeGameScreen, op1)

	return scaledBackground, scaledActiveGameScreen
}

func initializeSummonerButtons(outsideWidth, outsideHeight int) []*SummonerButton {
	var summonerButtons []*SummonerButton

	squareWidth := outsideWidth / 20

	buttonImage1 := ebiten.NewImage(squareWidth, squareWidth) // Create a button image placeholder
	buttonImage1.Fill(color.RGBA{255, 0, 0, 255})             // Fill the button with red color for demo purposes

	buttonImage2 := ebiten.NewImage(squareWidth, squareWidth) // Create a button image placeholder
	buttonImage2.Fill(color.RGBA{0, 255, 0, 255})             // Fill the button with red color for demo purposes

	buttonImage3 := ebiten.NewImage(squareWidth, squareWidth) // Create a button image placeholder
	buttonImage3.Fill(color.RGBA{0, 0, 255, 255})             // Fill the button with red color for demo purposes

	for i := 1; i < 6; i++ {
		summonerSpell1 := Button{x: outsideWidth / 6 * i, y: outsideHeight / 6, width: squareWidth, height: squareWidth, image: buttonImage1, originalImage: buttonImage1}
		summonerSpell2 := Button{x: outsideWidth/6*i + squareWidth, y: outsideHeight / 6, width: squareWidth, height: squareWidth, image: buttonImage2, originalImage: buttonImage2}
		championPortrait := Button{x: outsideWidth/6*i + squareWidth/2, y: outsideHeight/6 - squareWidth, width: squareWidth, height: squareWidth, image: buttonImage3, originalImage: buttonImage3}

		summonerButton := SummonerButton{summonerSpell1: summonerSpell1, summonerSpell2: summonerSpell2, championPortrait: championPortrait}
		summonerButtons = append(summonerButtons, &summonerButton)
	}

	return summonerButtons
}
