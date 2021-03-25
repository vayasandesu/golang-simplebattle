package gachapox

import (
	"fmt"
	"math/rand"
	"simplebattle/gachapon"
)

type GachapoxSlot struct {
	ItemName string
	Amount   int
}

type Gachapox struct {
	Data  []GachapoxSlot
	pools []GachapoxSlot
}

func RemoveAt(array []GachapoxSlot, index int) []GachapoxSlot {
	output := append(array[:index], array[index+1:]...)

	return output
}

func (gacha *Gachapox) Random() gachapon.Item {
	size := len(gacha.pools)
	if size == 0 {
		fmt.Println("[GACHAPOX] item pool is empty, cannot random")
		return gachapon.Item{}
	}

	index := rand.Intn(size)

	retrieve := gacha.pools[index]
	gacha.pools = RemoveAt(gacha.pools, index)

	return gachapon.Item{
		Name:   retrieve.ItemName,
		Amount: retrieve.Amount,
	}
}

func (gacha *Gachapox) Reset() {
	gacha.pools = make([]GachapoxSlot, len(gacha.Data))
	copy(gacha.pools, gacha.Data)
}

func (gacha Gachapox) GetRemaining() string {
	output := ""
	for _, item := range gacha.pools {
		output += fmt.Sprintf("Name : %s, Quantity %d", item.ItemName, item.Amount)
	}

	return output
}
