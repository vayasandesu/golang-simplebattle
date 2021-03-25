package scenario

import (
	"fmt"
	"math/rand"
	"simplebattle/gachapon"
	"simplebattle/gachapon/gachapin"
	"simplebattle/gachapon/gachapox"
)

func rng(gacha gachapon.IGachaponer) gachapon.Item {
	return gacha.Random()
}

func generateGachapin() gachapon.IGachaponer {
	slots := make([]gachapin.GachapinSlot, 5)
	slots[0] = gachapin.GachapinSlot{
		ProductName: "BITCOIN",
		Amount:      1,
		Weigth:      1,
	}
	slots[1] = gachapin.GachapinSlot{
		ProductName: "BANANA",
		Amount:      5,
		Weigth:      100,
	}
	slots[2] = gachapin.GachapinSlot{
		ProductName: "PANCAKE",
		Amount:      5,
		Weigth:      100,
	}
	slots[3] = gachapin.GachapinSlot{
		ProductName: "USD",
		Amount:      20,
		Weigth:      200,
	}
	slots[4] = gachapin.GachapinSlot{
		ProductName: "THB",
		Amount:      100,
		Weigth:      1000,
	}

	gacha := gachapin.Gachapin{
		Slots: slots,
	}

	return &gacha
}

func generateGachapox() gachapon.IGachaponer {
	itembox := make([]gachapox.GachapoxSlot, 10)
	for i := 0; i < 10; i++ {
		itembox[i] = gachapox.GachapoxSlot{
			ItemName: fmt.Sprintf("Potion %d", i),
			Amount:   rand.Intn(20),
		}
	}

	gacha := gachapox.Gachapox{
		Data: itembox,
	}
	gacha.Reset()

	return &gacha
}

func GachaponRunner() {
	fmt.Println("Gachapox generator")
	fmt.Println("Item should be empty after 10 roll")
	gachaponer := scenario.GenerateGachapox()
	for i := 1; i <= 15; i++ {
		item := scenario.RNG(gachaponer)
		fmt.Printf("Gachapon random %d got %+v\n", i, item)
	}

	fmt.Println("Gachapin generator")
	fmt.Println("Alway get item")
	gachaponer = scenario.GenerateGachapin()
	for i := 1; i <= 15; i++ {
		item := scenario.RNG(gachaponer)
		fmt.Printf("Gachapin random %d got %+v\n", i, item)
	}
}
