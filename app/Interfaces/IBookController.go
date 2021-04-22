package interfaces

type IBookController interface {
	Create(player1Name string, player2Name string) (string, error)
}