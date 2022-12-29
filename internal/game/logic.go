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
	mybody := state.You.Body
	avoid := NewSet(mybody...)

	// Add other snakes to the avoid list
	for _, snake := range state.Board.Snakes {
		avoid.add(snake.Body...)
	}

	boardWidth := state.Board.Width
	boardHeight := state.Board.Height
	food := state.Board.Food

	if neck.X < head.X || head.X == 0 || avoid.exits(Coord{X: head.X - 1, Y: head.Y}) {
		log.Println("not left")
		moves[left] = false
	}
	if neck.X > head.X || head.X == boardWidth-1 || avoid.exits(Coord{X: head.X + 1, Y: head.Y}) {
		log.Println("not right")
		moves[right] = false
	}
	if neck.Y < head.Y || head.Y == 0 || avoid.exits(Coord{X: head.X, Y: head.Y - 1}) {
		log.Println("not down")
		moves[down] = false
	}
	if neck.Y > head.Y || head.Y == boardHeight-1 || avoid.exits(Coord{X: head.X, Y: head.Y + 1}) {
		log.Println("not up")
		moves[up] = false
	}

	// choos a food at random
	safeMoves := []move{}
	for m, isSafe := range moves {
		if isSafe {
			safeMoves = append(safeMoves, m)
		}
	}
	nextMove := move(-1)

	// 0,0 point is located in the bottom left corner for some reason, therefore positive Y is going up
	if len(safeMoves) == 0 {
		nextMove = down
		log.Printf("%s MOVE %d: No safe moves found! Moving %s\n", state.Game.ID, state.Turn, nextMove)
	} else if len(food) > 0 {
		for _, f := range food {
			target := f
			log.Println("Going for target at: ", target)

			// for _, move := range [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			// }

			if target.X < head.X {
				for _, move := range safeMoves {
					if move == left {
						nextMove = left
					}
				}
			} else if target.X > head.X {
				for _, move := range safeMoves {
					if move == right {
						nextMove = right
					}
				}
			} else if target.Y > head.Y {
				for _, move := range safeMoves {
					if move == up {
						nextMove = up
					}
				}
			} else if target.Y < head.Y {
				for _, move := range safeMoves {
					if move == down {
						nextMove = down
					}
				}
			}

			if nextMove != -1 {
				break
			}
		}
	} else {
		nextMove = safeMoves[rand.Intn(len(safeMoves))]
	}

	if nextMove == -1 {
		nextMove = safeMoves[rand.Intn(len(safeMoves))]
	}

	log.Println("MOVING: ", nextMove)

	return BattlesnakeMoveResponse{
		Move: nextMove.String(),
	}
}
