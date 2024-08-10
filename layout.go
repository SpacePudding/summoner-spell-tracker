package main

import "github.com/hajimehoshi/ebiten/v2"

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {

	if g.scaledBackground == nil || g.scaledActiveGameScreen == nil {
		g.scaledBackground, g.scaledActiveGameScreen = createScaledImages(outsideWidth, outsideHeight, g.originalBackground, g.activeGameScreen)
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
