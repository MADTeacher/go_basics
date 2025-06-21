package client

type State int

const (
	waitNickNameConfirm State = iota
	mainMenu
	waitRoomJoin
	playing
	playerMove
	opponentMove
	endGame
	waitingOpponentInRoom
	waitResponseFromServer
)
