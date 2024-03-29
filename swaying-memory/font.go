package swayingmemory

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	mplusNormalFont font.Face
	mplusBigFont    font.Face
	jaKanjis        = []rune{}
)

func init() {
	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		log.Fatal(err)
	}
	mplusBigFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    48,
		DPI:     dpi,
		Hinting: font.HintingFull, // Use quantization to save glyph cache images.
	})
	if err != nil {
		log.Fatal(err)
	}

	// Adjust the line height.
	mplusBigFont = text.FaceWithLineHeight(mplusBigFont, 54)
}

func drawNormalText(screen *ebiten.Image, str string, x, y int, clr color.Color) {
	text.Draw(screen, str, mplusNormalFont, x, y, clr)
}

func drawBigText(screen *ebiten.Image, str string, x, y int, clr color.Color) {
	text.Draw(screen, str, mplusBigFont, x, y, clr)
}

func getNormalTextWidth(str string) int {
	b, _ := font.BoundString(mplusNormalFont, str)
	return (b.Max.X - b.Min.X).Ceil()
}

func getBigTextWidth(str string) int {
	b, _ := font.BoundString(mplusBigFont, str)
	return (b.Max.X - b.Min.X).Ceil()
}
