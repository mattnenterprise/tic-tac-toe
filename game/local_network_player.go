package game

import (
	"fmt"
	"net"
	"strconv"
)

type LocalNetworkPlayer struct {
	conn net.Conn
	*BasePlayer
}

func NewLocalNetworkPlayer(value SlotValue, conn net.Conn) *LocalNetworkPlayer {
	return &LocalNetworkPlayer{
		conn:       conn,
		BasePlayer: NewBasePlayer(value),
	}
}

func (h *LocalNetworkPlayer) MakeMove(game *TicTacToeGame) {
	var slot int
	success := false
	for !success {
		fmt.Println("Please enter a value 1-9 to represent the slot to make your move")
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
	fmt.Fprintf(h.conn, strconv.Itoa(slot)+"\n")
}
