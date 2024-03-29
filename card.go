package main

import (
	"fmt"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type CardType string

const (
	CardTypeAngry       CardType = "angry"
	CardTypeAutumn      CardType = "autumn"
	CardTypeAwake       CardType = "awake"
	CardTypeBalloon     CardType = "balloon"
	CardTypeBaseball    CardType = "baseball"
	CardTypeBeer        CardType = "beer"
	CardTypeBye         CardType = "bye"
	CardTypeCheer       CardType = "cheer"
	CardTypeCold        CardType = "cold"
	CardTypeCook        CardType = "cook"
	CardTypeCry         CardType = "cry"
	CardTypeEmbarrass   CardType = "embarrass"
	CardTypeFaint       CardType = "faint"
	CardTypeGoodMorning CardType = "good_morning"
	CardTypeHi          CardType = "hi"
	CardTypeHide        CardType = "hide"
	CardTypeHideAway    CardType = "hide_away"
	CardTypeHot         CardType = "hot"
	CardTypeHotSpring   CardType = "hot_spring"
	CardTypeHungry      CardType = "hungry"
	CardTypeLovely      CardType = "lovely"
	CardTypeNinja       CardType = "ninja"
	CardTypeNo          CardType = "no"
	CardTypeOk          CardType = "ok"
	CardTypeQuestion    CardType = "question"
	CardTypeRunAway     CardType = "run_away"
	CardTypeScare       CardType = "scare"
	CardTypeSigh        CardType = "sigh"
	CardTypeSleeping    CardType = "sleeping"
	CardTypeSleepy      CardType = "sleepy"
	CardTypeSorry       CardType = "sorry"
	CardTypeSpring      CardType = "spring"
	CardTypeSurprise    CardType = "surprise"
	CardTypeTehepero    CardType = "tehepero"
	CardTypeThankYou    CardType = "thank_you"
	CardTypeWork        CardType = "work"
)

var CardTypes = []CardType{
	CardTypeAngry,
	CardTypeAutumn,
	CardTypeAwake,
	CardTypeBalloon,
	CardTypeBaseball,
	CardTypeBeer,
	CardTypeBye,
	CardTypeCheer,
	CardTypeCold,
	CardTypeCook,
	CardTypeCry,
	CardTypeEmbarrass,
	CardTypeFaint,
	CardTypeGoodMorning,
	CardTypeHi,
	CardTypeHide,
	CardTypeHideAway,
	CardTypeHot,
	CardTypeHotSpring,
	CardTypeHungry,
	CardTypeLovely,
	CardTypeNinja,
	CardTypeNo,
	CardTypeOk,
	CardTypeQuestion,
	CardTypeRunAway,
	CardTypeScare,
	CardTypeSigh,
	CardTypeSleeping,
	CardTypeSleepy,
	CardTypeSorry,
	CardTypeSpring,
	CardTypeSurprise,
	CardTypeTehepero,
	CardTypeThankYou,
	CardTypeWork,
}

func loadGopher(t CardType) (*ebiten.Image, error) {
	gopher, _, err := ebitenutil.NewImageFromFile(fmt.Sprintf("resources/images/%s.png", t))
	if err != nil {
		return nil, err
	}
	return gopher, nil
}

const (
	CardWidth  = 60
	CardHeight = 80
)

type Card struct {
	Type       CardType
	frontImage *ebiten.Image
	backImage  *ebiten.Image
	x, y       float64
	dx, dy     float64
}

func NewCard(t CardType, x, y float64) *Card {
	gopher, err := loadGopher(t)
	if err != nil {
		panic(err)
	}
	front := ebiten.NewImage(CardWidth, CardHeight)
	front.Fill(color.White)
	vector.StrokeLine(front, 0, 1, CardWidth, 1, 1, color.Black, false)
	vector.StrokeLine(front, CardWidth, 0, CardWidth, CardHeight, 1, color.Black, false)
	vector.StrokeLine(front, CardWidth, CardHeight, 0, CardHeight, 1, color.Black, false)
	vector.StrokeLine(front, 1, 0, 1, CardHeight, 1, color.Black, false)

	scale := 0.15
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(float64(CardWidth/2-float64(gopher.Bounds().Dx())*scale/2), float64(CardHeight/2-float64(gopher.Bounds().Dy())*scale/2))
	front.DrawImage(gopher, op)
	back := ebiten.NewImage(CardWidth, CardHeight)
	back.Fill(color.RGBA{0x80, 0x80, 0x80, 0xff})

	if err != nil {
		panic(err)
	}
	return &Card{
		Type:       t,
		frontImage: front,
		backImage:  back,
		x:          x,
		y:          y,
		dx:         rand.Float64()*2 - 1,
		dy:         rand.Float64()*2 - 1,
	}
}

func (c *Card) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(c.x), float64(c.y))
	screen.DrawImage(c.frontImage, op)
}

func (c *Card) Update() {
	c.x += c.dx
	c.y += c.dy
	if c.x < 0 || c.x > screenWidth-CardWidth {
		c.dx *= -1
	}
	if c.y < 0 || c.y > screenHeight-CardHeight {
		c.dy *= -1
	}
}
