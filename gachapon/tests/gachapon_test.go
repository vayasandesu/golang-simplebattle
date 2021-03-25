package test

import (
	"fmt"
	"math/rand"
	"simplebattle/gachapon"
	"simplebattle/gachapon/gachapox"
	"testing"
)

func generateGachapox(amount int) *gachapox.Gachapox {
	itembox := make([]gachapox.GachapoxSlot, 10)
	for i := 0; i < amount; i++ {
		itembox[i] = gachapox.GachapoxSlot{
			ItemName: fmt.Sprintf("Potion %d", i),
			Amount:   rand.Intn(20),
		}
	}

	gacha := gachapox.Gachapox{
		Data: itembox,
	}
	return &gacha
}

// folder : gachapon/tests
// package : test
func Test_GachapoxShouldReturnDefaultWhenEmpty(t *testing.T) {
	amount := 10
	var item gachapon.Item

	// init box
	box := generateGachapox(amount)
	box.Reset()

	for i := 0; i < amount; i++ {
		item = box.Random()
		if item.Name == "" {
			t.Fatalf(`First %d roll should have item`, amount)
		}
	}

	item = box.Random()
	if item.Name != "" {
		t.Fatalf(`After %d roll should not get item but get %s`, amount, item.Name)
	}
}
