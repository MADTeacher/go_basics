package game

type Player struct {
	figure BoardField
}

func NewPlayer() *Player {
	return &Player{figure: cross}
}

func (p *Player) switchPlayer() {
	if p.figure == cross {
		p.figure = nought
	} else {
		p.figure = cross
	}
}

func (p *Player) getSymbol() string {
	if p.figure == cross {
		return "X"
	}
	return "O"
}
