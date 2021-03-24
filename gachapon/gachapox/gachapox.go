package gachapox

import (
	"math/rand"
	"simplebattle/gachapon"
)

type GachapoxSlot struct {
	ItemName string
	Amount   int
}

type Gachapox struct {
	Data  []GachapoxSlot
	Pools []GachapoxSlot
}

func (gacha *Gachapox) Random() gachapon.Item {
	size := len(gacha.Pools)
	index := rand.Intn(size)

	item := gacha.Pools[index]

	return gachapon.Item{
		Name:   item.ItemName,
		Amount: item.Amount,
	}
}

func (gacha *Gachapox) Reset() {
	copy(gacha.Pools, gacha.Data)
}
