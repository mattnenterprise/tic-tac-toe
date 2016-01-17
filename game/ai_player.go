package game

type AiPlayer struct {
	*BasePlayer
}

func NewAiPlayer(value SlotValue) *AiPlayer {
	return &AiPlayer{
		BasePlayer: NewBasePlayer(value),
	}
}

func (a *AiPlayer) MakeMove(game *TicTacToeGame) {
	// TODO make the AI not stupid using minimax
	for i := 0; i < BOARDHEIGHT*BOARDWIDTH; i++ {
		success := game.MakeMove(i, a.value)
		if success {
			return
		}
	}
}
