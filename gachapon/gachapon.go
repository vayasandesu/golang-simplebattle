package gachapon

type Item struct {
	Name   string
	Amount int
}

type IGachaponer interface {
	Random() Item
}
