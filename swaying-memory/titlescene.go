package swayingmemory

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type TitleScene struct {
}

func (s *TitleScene) Update(state *GameState) error {
	return nil
}

func (s *TitleScene) Draw(screen *ebiten.Image) {
	title := "Swaying Memory"
	drawBigText(screen, title, ScreenWidth/2-getBigTextWidth(title)/2, 100, color.White)

	subtitle := "Press Enter to start"
	drawNormalText(screen, subtitle, ScreenWidth/2-getNormalTextWidth(subtitle)/2, 200, color.White)
}
