package game

import "fmt"

type HumanPlayer struct {
	*BasePlayer
}

func NewHumanPlayer(value SlotValue) *HumanPlayer {
	return &HumanPlayer{
		BasePlayer: NewBasePlayer(value),
	}
}

func (h *HumanPlayer) MakeMove(game *TicTacToeGame) {
	success := false
	for !success {
		fmt.Println("Please enter a value 1-9 to represent the slot to make your move")
		var slot int
		_, err := fmt.Scanf("%d", &slot)
		if err != nil {
			fmt.Println("Invalid slot selected")
			break
		}
		success = game.MakeMove(slot-1, h.value)
		if !success {
			fmt.Println("Cannot make a move in that slot")
		}
	}
}
