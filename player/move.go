package player

type Move struct {
	FromCol int
	FromRow int

	ToCol int
	ToRow int

	WhitePlayer bool
}