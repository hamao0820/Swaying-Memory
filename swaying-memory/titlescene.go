package swayingmemory

import (
	"image"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var gopherImage *ebiten.Image
var gopher *Gopher

type Gopher struct {
	x, y   float64
	dx, dy float64
}

func (g *Gopher) Update() {
	g.x += g.dx
	g.y += g.dy
	if g.x < 0 || ScreenWidth < g.x+float64(gopherImage.Bounds().Dx()) {
		g.dx *= -1
	}
	if g.y < 0 || ScreenHeight < g.y+float64(gopherImage.Bounds().Dy()) {
		g.dy *= -1
	}
}

func (g *Gopher) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.x, g.y)
	screen.DrawImage(gopherImage, op)
}

func init() {
	img, _, err := ebitenutil.NewImageFromFile("resources/images/faint.png")
	if err != nil {
		panic(err)
	}
	// 絵の部分を切り取る
	gopherImage = ebiten.NewImageFromImage(img.SubImage(image.Rect(70, 0, 300, 370)))

	gopher = &Gopher{
		x:  rand.Float64() * (ScreenWidth - float64(gopherImage.Bounds().Dx())),
		y:  rand.Float64() * (ScreenHeight - float64(gopherImage.Bounds().Dy())),
		dx: 2,
		dy: 2,
	}
}

type TitleScene struct{}

func (s *TitleScene) drawBackground(screen *ebiten.Image) {
	gopher.Draw(screen)
}

func (s *TitleScene) Update(state *GameState) error {
	gopher.Update()
	return nil
}

func (s *TitleScene) Draw(screen *ebiten.Image) {
	s.drawBackground(screen)

	title := "Swaying Memory"
	drawBigText(screen, title, ScreenWidth/2-getBigTextWidth(title)/2, 100, color.White)

	subtitle := "Press Enter to start"
	drawNormalText(screen, subtitle, ScreenWidth/2-getNormalTextWidth(subtitle)/2, 200, color.White)
}
