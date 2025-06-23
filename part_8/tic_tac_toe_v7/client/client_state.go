package client

type State int

const (
	// ожидание подтверждения никнейма от сервера
	waitNickNameConfirm State = iota
	// главное меню
	mainMenu
	// ход игрока
	playerMove
	// ход оппонента
	opponentMove
	// конец игры
	endGame
	// ожидание присоединения оппонента
	waitingOpponentInRoom
	// ожидание ответа от сервера
	waitResponseFromServer
)
