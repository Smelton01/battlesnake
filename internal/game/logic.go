package game

import (
	"log"
	"math/rand"
)

type move int8

const (
	up move = iota
	down
	left
	right
)

// Info controls the battlensnake appearance and other permissions
// See [Personalization]: https://docs.battlesnake.com/references/personalization
func Info() BattlesnakeInfoResponse {
	log.Printf("INFO")
	return BattlesnakeInfoResponse{
		APIVersion: "0.1",
		Author:     "Smelton01",
		Color:      string(blue),
		Head:       string(blue),
		Tail:       string(blue),
	}
}

// Start is called every time the battlesnake enters a game.
// The state param contains information about the game that is a bout to be played.
func Start(state GameState) {
	log.Printf("%s START\n", state.Game.ID)
}

// End is called when a game your battlesnake was participating in ends.
func End(state GameState) {
	log.Printf("%s END\n", state.Game.ID)
}

func Move(state GameState) BattlesnakeMoveResponse {
	moves := map[move]bool{
		up:    true,
		down:  true,
		left:  true,
		right: true,
	}

	// Step 0: Don't let your battlesnake move back on its own neck
	head := state.You.Body[0]
	neck := state.You.Body[1]

	if neck.X < head.X {
		moves[left] = false
	}
	if neck.X > head.X {
		moves[right] = false
	}
	if neck.Y < head.Y {
		moves[down] = false
	}
	if neck.Y < head.Y {
		moves[up] = false
	}

	var nextMove string

	safeMoves := []string{}
	// TODO use stringer to make string from move
	for m, isSafe := range moves {
		if isSafe {
			safeMoves = append(safeMoves, m.String())
		}
	}

	if len(safeMoves) == 0 {
		nextMove = down.String()
		log.Printf("%s MOVE %d: No safe moves found! Moving %s\n", state.Game.ID, state.Turn, nextMove)
	} else {
		nextMove = safeMoves[rand.Intn(len(safeMoves))]
	}

	return BattlesnakeMoveResponse{
		Move: nextMove,
	}
}

func (m move) String() string {
	switch m {
	case 0:
		return "up"
	case 1:
		return "down"
	case 2:
		return "left"
	case 3:
		return "right"
	default:
		return ""
	}
}
