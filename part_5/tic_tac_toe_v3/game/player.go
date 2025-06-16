package game

type Player struct {
	Figure BoardField `json:"figure"`
}

func NewPlayer() *Player {
	return &Player{Figure: cross}
}

func (p *Player) switchPlayer() {
	if p.Figure == cross {
		p.Figure = nought
	} else {
		p.Figure = cross
	}
}

func (p *Player) getSymbol() string {
	if p.Figure == cross {
		return "X"
	}
	return "O"
}
