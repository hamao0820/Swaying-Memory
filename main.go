package main

import (
	"fmt"
	"image/color"
	"log"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

type GameMode int

const (
	ModeZeroFlipped GameMode = iota
	ModeOneFlipped
	ModeTwoFlipped
)

const poseTime = 30

type Game struct {
	cards         []*Card
	fillipedCards [2]*Card
	t             int
	mode          GameMode
}

func newGame() *Game {
	cards := []*Card{}
	for _, i := range sample(16, 3) {
		cards = append(cards, NewCard(CardTypes[i], rand.Float64()*float64(screenWidth-CardWidth), rand.Float64()*float64(screenHeight-CardHeight)))
		cards = append(cards, NewCard(CardTypes[i], rand.Float64()*float64(screenWidth-CardWidth), rand.Float64()*float64(screenHeight-CardHeight)))
	}

	return &Game{
		cards: cards,
		mode:  ModeZeroFlipped,
	}
}

func (g *Game) Update() error {
	for _, card := range g.cards {
		card.Update()
	}

	if g.mode == ModeTwoFlipped {
		g.t++
		ebiten.SetCursorShape(ebiten.CursorShapeNotAllowed)
		if g.t < poseTime {
			return nil
		}

		ebiten.SetCursorShape(ebiten.CursorShapeDefault)
		if g.fillipedCards[0] != nil && g.fillipedCards[1] != nil {
			if g.fillipedCards[0].Type == g.fillipedCards[1].Type {
				g.fillipedCards[0].matched = true
				g.fillipedCards[1].matched = true
			} else {
				g.fillipedCards[0].flipped = false
				g.fillipedCards[1].flipped = false
			}
			g.fillipedCards[0] = nil
			g.fillipedCards[1] = nil
		}

		g.t = 0
		g.mode = ModeZeroFlipped
	}

	x, y := ebiten.CursorPosition()
	var hoveredCard *Card
	for _, card := range g.cards {
		if card.matched {
			continue
		}
		card.hovered = false
		if card.In(x, y) {
			if hoveredCard == nil || card.zIndex > hoveredCard.zIndex {
				hoveredCard = card
			}
		}
	}

	ebiten.SetCursorShape(ebiten.CursorShapeDefault)
	if hoveredCard == nil {
		return nil
	}
	ebiten.SetCursorShape(ebiten.CursorShapePointer)
	hoveredCard.hovered = true

	if g.mode == ModeZeroFlipped {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			if !hoveredCard.flipped {
				hoveredCard.flipped = true
				g.fillipedCards[0] = hoveredCard
				g.mode = ModeOneFlipped
			}
		}
	}

	if g.mode == ModeOneFlipped {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			if !hoveredCard.flipped {
				hoveredCard.flipped = true
				g.fillipedCards[1] = hoveredCard
				g.mode = ModeTwoFlipped
			}
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x80, 0xff, 0x80, 0xff})
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%f", ebiten.ActualFPS()))
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

func sample(n int, r int) []int {
	sample := make([]int, n)
	for i := 0; i < n; i++ {
		sample[i] = i
	}
	rand.Shuffle(n, func(i, j int) {
		sample[i], sample[j] = sample[j], sample[i]
	})
	return sample[:r]
}
