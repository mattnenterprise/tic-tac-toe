package game

import (
	"bufio"
	"net"
	"strconv"
	"strings"
)

type RemoteNetworkPlayer struct {
	conn net.Conn
	*BasePlayer
}

func NewRemoteNetworkPlayer(value SlotValue, conn net.Conn) *RemoteNetworkPlayer {
	return &RemoteNetworkPlayer{
		conn:       conn,
		BasePlayer: NewBasePlayer(value),
	}
}

func (h *RemoteNetworkPlayer) MakeMove(game *TicTacToeGame) {
	slot, err := bufio.NewReader(h.conn).ReadString('\n')
	if err != nil {
		// TODO remove this later
		panic(err)
	}
	strSlot, err := strconv.Atoi(strings.Trim(slot, " \n"))
	if err != nil {
		// TODO remove this later
		panic(err)
	}
	game.MakeMove(strSlot-1, h.Value())
}
