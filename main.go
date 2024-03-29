package main

import (
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type Game struct {
	cards []*Card
}

func newGame() *Game {
	cards := []*Card{}
	cards = append(cards, NewCard(CardTypeBeer, rand.Float64()*(screenWidth-CardWidth), rand.Float64()*(screenHeight-CardHeight)))
	cards = append(cards, NewCard(CardTypeCook, rand.Float64()*(screenWidth-CardWidth), rand.Float64()*(screenHeight-CardHeight)))
	cards = append(cards, NewCard(CardTypeBalloon, rand.Float64()*(screenWidth-CardWidth), rand.Float64()*(screenHeight-CardHeight)))
	cards = append(cards, NewCard(CardTypeSleepy, rand.Float64()*(screenWidth-CardWidth), rand.Float64()*(screenHeight-CardHeight)))

	return &Game{
		cards: cards,
	}
}

func (g *Game) Update() error {
	for _, card := range g.cards {
		card.Update()
	}

	x, y := ebiten.CursorPosition()
	var hoveredCard *Card
	for _, card := range g.cards {
		card.hovered = false
		if card.In(x, y) {
			if hoveredCard == nil || card.zIndex > hoveredCard.zIndex {
				hoveredCard = card
			}
		}
	}
	if hoveredCard == nil {
		return nil
	}
	hoveredCard.hovered = true
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		if !hoveredCard.flipped {
			hoveredCard.flipped = !hoveredCard.flipped
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x80, 0xff, 0x80, 0xff})
	for _, card := range g.cards {
		card.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Swaying Memory")
	g := newGame()
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
