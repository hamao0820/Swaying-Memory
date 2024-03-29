package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
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
	for i, cardType := range CardTypes {
		cards = append(cards, NewCard(cardType, i%5*(CardWidth+5), +i/5*(CardHeight+5)))
	}
	return &Game{
		cards: cards,
	}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
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
