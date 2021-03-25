package webapp

import (
	"simplebattle/gachapon"
	"simplebattle/gachapon/gachapin"
	"simplebattle/webapp/routes"
	"simplebattle/webapp/services"
)

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

func Start() {
	// can switch between gachapin and gachapox
	service := services.MemoryGachaponService{
		Gachapon: generateGachapin(),
		Wallet:   make(map[string]int),
	}

	// can select framework type eg. echo / gin (must reimplement ... no choice)
	// can switch memoryService and mongoDbService (not implement yet)
	// without effected the routes code
	route := routes.EchoRoute{
		Service: &service,
		Port:    "3000",
	}

	route.Start()
}
