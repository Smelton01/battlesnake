package game

import (
	"log"
	"math/rand"
)

//go:generate go run golang.org/x/tools/cmd/stringer -type=move
type move int8

const (
	up move = iota
	down
	left
	right
)

// Info controls the battlensnake appearance and other permissions
func Info() BattlesnakeInfoResponse {
	log.Printf("INFO")

	return BattlesnakeInfoResponse{
		APIVersion: "1",
		Author:     "Smelton01",
		Color:      string(blue),
		Head:       string(blue),
		Tail:       string(blue),
	}
}

// Start is called every time the battlesnake enters a game.
// The state param contains information about the game that is a bout to be played.
func Start(state *GameState) {
	log.Printf("START GAME: %s\n", state.Game.ID)
}

// End is called when a game your battlesnake was participating in ends.
func End(state *GameState) {
	log.Printf("ENDING GAME: %s\n", state.Game.ID)
}

func Move(state *GameState) BattlesnakeMoveResponse {
	moves := map[move]bool{
		up:    true,
		down:  true,
		left:  true,
		right: true,
	}

	// Step 0: Don't let your battlesnake move back on its own neck
	head := state.You.Body[0]
	neck := state.You.Body[1]

	boardWidth := state.Board.Width
	boardHeight := state.Board.Height

	if neck.X < head.X || head.X == 0 {
		log.Println("not left")
		moves[left] = false
	}
	if neck.X > head.X || head.X == boardWidth-1 {
		log.Println("not right")
		moves[right] = false
	}
	if neck.Y < head.Y || head.Y == 0 {
		log.Println("not down")
		moves[down] = false
	}
	if neck.Y > head.Y || head.Y == boardHeight-1 {
		log.Println("not up")
		moves[up] = false
	}

	var nextMove move

	safeMoves := []move{}
	for m, isSafe := range moves {
		if isSafe {
			safeMoves = append(safeMoves, m)
		}
	}

	if len(safeMoves) == 0 {
		nextMove = down
		log.Printf("%s MOVE %d: No safe moves found! Moving %s\n", state.Game.ID, state.Turn, nextMove)
	} else {
		nextMove = safeMoves[rand.Intn(len(safeMoves))]
	}

	log.Println("MOVING: ", nextMove)

	return BattlesnakeMoveResponse{
		Move: nextMove.String(),
	}
}

// func (m move) String() string {
// 	switch m {
// 	case 0:
// 		return "up"
// 	case 1:
// 		return "down"
// 	case 2:
// 		return "left"
// 	case 3:
// 		return "right"
// 	default:
// 		return ""
// 	}
// }
