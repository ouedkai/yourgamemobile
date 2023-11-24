package example

import (
	"github.com/hajimehoshi/ebiten/v2/mobile"
	"github.com/ouedkai/yourgame"
)

func init() {
	// yourgame.Game must implement ebiten.Game interface.
	// For more details, see
	// * https://pkg.go.dev/github.com/hajimehoshi/ebiten/v2#Game
	// Create an instance of GameImpl, which implements the Game interface
	game := &yourgame.Game{}
	// Set the game for mobile
	mobile.SetGame(game)
}

// Dummy is a dummy exported function.
//
// gomobile doesn't compile a package that doesn't include any exported function.
// Dummy forces gomobile to compile this package.
func Dummy() {}
