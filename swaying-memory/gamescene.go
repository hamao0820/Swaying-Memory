package swayingmemory

import (
	"fmt"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type GameMode int

const (
	ModeZeroFlipped GameMode = iota
	ModeOneFlipped
	ModeTwoFlipped
)

const poseTime = 30

type GameScene struct {
	cards         []*Card
	fillipedCards [2]*Card
	tickCounts    int
	mode          GameMode
}

func NewGameScene() *GameScene {
	cards := make([]*Card, 0)
	for _, i := range sample(len(CardTypes), 3) {
		cards = append(cards, NewCard(CardTypes[i], rand.Float64()*(ScreenWidth-CardWidth), rand.Float64()*(ScreenHeight-CardHeight)))
		cards = append(cards, NewCard(CardTypes[i], rand.Float64()*(ScreenWidth-CardWidth), rand.Float64()*(ScreenHeight-CardHeight)))
	}
	sample := sample(len(cards), len(cards))
	for i, j := range sample {
		cards[i], cards[j] = cards[j], cards[i]
	}

	return &GameScene{
		cards: cards,
	}
}

func (s *GameScene) Update(state *GameState) error {
	for _, card := range s.cards {
		card.Update()
	}

	if s.mode == ModeTwoFlipped {
		ebiten.SetCursorShape(ebiten.CursorShapeNotAllowed)
		if s.tick() {
			return nil
		}

		ebiten.SetCursorShape(ebiten.CursorShapeDefault)
		if s.fillipedCards[0] != nil && s.fillipedCards[1] != nil {
			if s.fillipedCards[0].Type == s.fillipedCards[1].Type {
				s.fillipedCards[0].matched = true
				s.fillipedCards[1].matched = true
			} else {
				s.fillipedCards[0].flipped = false
				s.fillipedCards[1].flipped = false
			}
			s.fillipedCards[0] = nil
			s.fillipedCards[1] = nil
		}

		s.tickCounts = 0
		s.mode = ModeZeroFlipped
	}

	x, y := ebiten.CursorPosition()
	var hoveredCard *Card
	for _, card := range s.cards {
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

	if s.mode == ModeZeroFlipped {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			if !hoveredCard.flipped {
				hoveredCard.flipped = true
				s.fillipedCards[0] = hoveredCard
				s.mode = ModeOneFlipped
			}
		}
	}

	if s.mode == ModeOneFlipped {
		if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
			if !hoveredCard.flipped {
				hoveredCard.flipped = true
				s.fillipedCards[1] = hoveredCard
				s.mode = ModeTwoFlipped
			}
		}
	}

	return nil
}

func (s *GameScene) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0x80, 0xff, 0x80, 0xff})
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%f", ebiten.ActualFPS()))
	for _, card := range s.cards {
		card.Draw(screen)
	}
}

func (s *GameScene) tick() bool {
	s.tickCounts++
	return s.tickCounts < poseTime
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
