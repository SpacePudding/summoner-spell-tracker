package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type SummonerButton struct {
	summonerSpell1   Button
	summonerSpell2   Button
	championPortrait Button
}

type Button struct {
	x, y, width, height int
	image               *ebiten.Image
	cooldownTimer       int // in seconds
	isCoolingDown       bool
	originalImage       *ebiten.Image
}

func (b *SummonerButton) Draw(screen *ebiten.Image) {

	// Draw the button
	op1 := &ebiten.DrawImageOptions{}
	op1.GeoM.Translate(float64(b.summonerSpell1.x), float64(b.summonerSpell1.y))
	screen.DrawImage(b.summonerSpell1.image, op1)

	op2 := &ebiten.DrawImageOptions{}
	op2.GeoM.Translate(float64(b.summonerSpell2.x), float64(b.summonerSpell2.y))
	screen.DrawImage(b.summonerSpell2.image, op2)

	op3 := &ebiten.DrawImageOptions{}
	op3.GeoM.Translate(float64(b.championPortrait.x), float64(b.championPortrait.y))
	screen.DrawImage(b.championPortrait.image, op3)

	// // If the button is cooling down, overlay a darker shade and draw the cooldown timer
	// if b.isCoolingDown {
	// 	// Darken the button
	// 	darkOverlay := ebiten.NewImage(b.width, b.height)
	// 	darkOverlay.Fill(color.RGBA{0, 0, 0, 128}) // semi-transparent black overlay
	// 	screen.DrawImage(darkOverlay, op)

	// 	// Draw the cooldown timer as text
	// 	// ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%d", b.cooldownTimer), b.x+10, b.y+10)
	// }
}

func (b *Button) Update() {
	if b.isCoolingDown {
		b.cooldownTimer--
		if b.cooldownTimer <= 0 {
			b.isCoolingDown = false
			b.image = b.originalImage // Reset to original image
		}
	}
}
