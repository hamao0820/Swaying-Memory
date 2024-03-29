package swayingmemory

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

// Input manages the input state including gamepads and keyboards.
type Input struct {
	gamepadIDs []ebiten.GamepadID
}

// GamepadIDButtonPressed returns a gamepad ID where at least one button is pressed.
// If no button is pressed, GamepadIDButtonPressed returns -1.
func (i *Input) GamepadIDButtonPressed() ebiten.GamepadID {
	i.gamepadIDs = ebiten.AppendGamepadIDs(i.gamepadIDs[:0])
	for _, id := range i.gamepadIDs {
		for b := ebiten.GamepadButton(0); b <= ebiten.GamepadButtonMax; b++ {
			if ebiten.IsGamepadButtonPressed(id, b) {
				return id
			}
		}
	}
	return -1
}

func (i *Input) IsClicked() bool {
	return inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)
}

func (i *Input) IsPressedSpace() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeySpace)
}
