package game

import (
	"fmt"
	"github.com/rileya/flamingo/gfx"
	"time"
)

// Main object representing a game.
type Game struct {
	currentState State
	nextState    State
	gfxContext   *gfx.Context
}

func CreateGame() *Game {
	return &Game{nil, nil, nil}
}

// Start the main loop with |state|.
func (game *Game) Start(state State) error {
	context, err := gfx.CreateContext(800, 600)
	if err != nil {
		return err
	}
	game.gfxContext = context
	game.currentState = state
	game.currentState.Init(game)
	game.mainLoop()
	return nil
}

func (game *Game) mainLoop() {
	startTime := time.Now()
	lastFrameTime := startTime
	for {
		now := time.Now()
		delta := time.Since(lastFrameTime)
		lastFrameTime = now

		keepLooping, err := game.currentState.Update(delta)

		if err != nil {
			fmt.Println(err)
			break
		}

		if !keepLooping {
			game.currentState.Deinit()
			game.currentState = game.nextState
			if game.currentState == nil {
				break
			}
			game.currentState.Init(game)
		}

		game.gfxContext.SwapBuffers()
	}
	game.gfxContext.Release()
}

func (game *Game) GetGfxContext() *gfx.Context {
	return game.gfxContext
}
