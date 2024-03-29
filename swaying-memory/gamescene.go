package swayingmemory

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type GameMode int

const (
	ModeZeroFlipped GameMode = iota
	ModeOneFlipped
	ModeTwoFlipped
	ModeClear
)

const poseTime = 30

var cardCounts = 10

type GameScene struct {
	cards         []*Card
	fillipedCards [2]*Card
	tickCounts    int
	matchedCounts int
	mode          GameMode
	startTime     time.Time
	time          int64
}

func NewGameScene() *GameScene {
	cards := make([]*Card, 0)
	for _, i := range sample(len(CardTypes), cardCounts) {
		cards = append(cards, NewCard(CardTypes[i], rand.Float64()*(ScreenWidth-CardWidth), rand.Float64()*(ScreenHeight-CardHeight)))
		cards = append(cards, NewCard(CardTypes[i], rand.Float64()*(ScreenWidth-CardWidth), rand.Float64()*(ScreenHeight-CardHeight)))
	}

	return &GameScene{
		cards:     cards,
		startTime: time.Now(),
	}
}

func (s *GameScene) Update(state *GameState) error {
	for _, card := range s.cards {
		card.Update()
	}

	if s.mode == ModeClear {
		if state.Input.IsPressedSpace() {
			state.SceneManager.GoTo(&TitleScene{})
		}
		return nil
	}
	s.time = time.Since(s.startTime).Milliseconds()

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
				s.matchedCounts++
				if s.matchedCounts == cardCounts {
					s.mode = ModeClear
					return nil
				}
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
	for _, card := range s.cards {
		card.Draw(screen)
	}

	time := fmt.Sprintf("time: %03.3f", float64(s.time)/1000)
	if s.mode == ModeClear {
		message := "Clear!"
		drawBigText(screen, message, ScreenWidth/2-getBigTextWidth(message)/2, ScreenHeight/2-20, color.White)
		drawNormalText(screen, time, ScreenWidth/2-getNormalTextWidth(time)/2, ScreenHeight/2+50, color.White)
		drawNormalText(screen, "Press Space to Title", ScreenWidth/2-getNormalTextWidth("PPress Space to Title")/2, ScreenHeight/2+100, color.White)
	} else {
		drawNormalText(screen, time, 0, 20, color.White)
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
