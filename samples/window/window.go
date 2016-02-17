package main

import (
	"fmt"
	"github.com/rileya/flamingo/game"
	"github.com/rileya/flamingo/gfx"
	"runtime"
	"time"
)

type SampleState struct {
	game        *game.Game
	timeElapsed float64
}

func (state *SampleState) Init(game *game.Game) {
	fmt.Println("Initializing...")
	state.game = game
}

func (state *SampleState) Update(delta time.Duration) (bool, error) {
	state.timeElapsed += delta.Seconds()
	state.game.GetGfxContext().SetBackgroundColor(
		&gfx.Colorf{1.0, 0.0, 0.0, 1.0})
	state.game.GetGfxContext().Clear()
	return state.timeElapsed < 5.0, nil
}

func (state *SampleState) Deinit() {
	fmt.Println("De-Initializing.")
}

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

func main() {
	gameObject := game.CreateGame()
	gameObject.Start(&SampleState{})
}
