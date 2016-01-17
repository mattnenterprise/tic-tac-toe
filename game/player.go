package game

type Player interface {
	MakeMove(game *TicTacToeGame)
	Value() SlotValue
}

type BasePlayer struct {
	value SlotValue
}

func NewBasePlayer(v SlotValue) *BasePlayer {
	return &BasePlayer{
		value: v,
	}
}

func (b *BasePlayer) Value() SlotValue {
	return b.value
}
