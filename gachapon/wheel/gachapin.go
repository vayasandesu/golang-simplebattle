package gachapin

import (
	"math/rand"
	"simplebattle/gachapon"
)

type Gachapin struct {
	Slots []GachapinSlot
}

func (box Gachapin) GetMaxWeight() int {
	totalWeight := 0
	for _, item := range box.Slots {
		totalWeight += item.Weigth
	}

	return totalWeight
}

func (box *Gachapin) Random() gachapon.Item {
	maxWeight := box.GetMaxWeight()
	rng := rand.Intn(maxWeight)

	for _, item := range box.Slots {
		rng -= item.Weigth
		if rng <= 0 {
			return gachapon.Item{
				Name:   item.ProductName,
				Amount: item.Amount,
			}
		}
	}

	return gachapon.Item{}
}
