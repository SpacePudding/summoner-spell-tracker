package main

import (
	"fmt"
	"image"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

const (
	textFontSize = 24
)

type SummonerButton struct {
	summonerSpell1   Button
	summonerSpell2   Button
	championPortrait Button
}

type Button struct {
	x, y, width, height int
	image               *ebiten.Image
	cooldown            int
	cooldownTimer       int
	isCoolingDown       bool
	originalImage       *ebiten.Image
}

func (b *SummonerButton) Draw(screen *ebiten.Image) {
	// Draw all buttons
	b.drawButton(screen, &b.summonerSpell1)
	b.drawButton(screen, &b.summonerSpell2)
	b.drawButton(screen, &b.championPortrait)

	// Draw cooldown masks for buttons with active cooldowns
	if b.summonerSpell1.isCoolingDown {
		b.drawCooldownMask(screen, &b.summonerSpell1)
		b.drawCooldownTimer(screen, &b.summonerSpell1)
	}
	if b.summonerSpell2.isCoolingDown {
		b.drawCooldownMask(screen, &b.summonerSpell2)
		b.drawCooldownTimer(screen, &b.summonerSpell2)
	}
}

func (b *SummonerButton) drawButton(screen *ebiten.Image, btn *Button) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(btn.x), float64(btn.y))
	screen.DrawImage(btn.image, op)
}

func (b *SummonerButton) drawCooldownMask(screen *ebiten.Image, btn *Button) {
	cooldownFraction := float64(btn.cooldownTimer) / float64(btn.cooldown)

	// Create a mask that covers the button
	mask := ebiten.NewImage(btn.width, btn.height)
	mask.Fill(color.RGBA{0, 0, 0, 128}) // Fully shaded

	// Calculate the angle for revealing the button
	arcRadius := float64(btn.width) / 2
	arcAngle := 2 * math.Pi * (1 - cooldownFraction) // Clockwise angle in radians

	for y := 0; y < btn.height; y++ {
		for x := 0; x < btn.width; x++ {
			dx := float64(x) - arcRadius
			dy := float64(y) - arcRadius
			angle := math.Atan2(dy, dx) + math.Pi/2 // Start from 12 o'clock

			// Normalize the angle to be between 0 and 2*math.Pi
			if angle < 0 {
				angle += 2 * math.Pi
			}

			// Reveal the section of the button based on the cooldown progress
			if angle <= arcAngle {
				mask.Set(x, y, color.RGBA{0, 0, 0, 0}) // Transparent (revealing)
			}
		}
	}

	// Draw the mask on top of the button
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(btn.x), float64(btn.y))
	screen.DrawImage(mask, op)
}

func (b *SummonerButton) drawCooldownTimer(screen *ebiten.Image, btn *Button) {
	// Create a new image for the text
	text := fmt.Sprintf("%d", btn.cooldownTimer)
	textImage := ebiten.NewImage(textFontSize*len(text), textFontSize)
	textImage.Fill(color.Transparent)

	// Draw the text
	face := basicfont.Face7x13
	d := &font.Drawer{
		Dst:  textImage,
		Src:  image.NewUniform(color.White),
		Face: face,
		Dot:  fixed.Point26_6{fixed.Int26_6(0), fixed.Int26_6(textFontSize)},
	}
	d.DrawString(text)

	// Center the text image on the button
	textOp := &ebiten.DrawImageOptions{}
	textOp.GeoM.Translate(float64(btn.x+btn.width/2)-float64(textFontSize*len(text))/2, float64(btn.y+btn.height/2)-float64(textFontSize)/2)

	// Draw the text image on the screen
	screen.DrawImage(textImage, textOp)
}

func (b *Button) Update() {
	if b.isCoolingDown {
		b.cooldownTimer--
		if b.cooldownTimer <= 0 {
			b.isCoolingDown = false
			b.cooldownTimer = b.cooldown
		}
	}
}
