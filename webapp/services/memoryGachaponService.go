package services

import (
	"errors"
	"simplebattle/gachapon"
)

type MemoryGachaponService struct {
	Gachapon gachapon.IGachaponer
	Wallet   map[string]int
}

func (service *MemoryGachaponService) Roll(user string, times int) ([]gachapon.Item, error) {
	wallet, ok := service.Wallet[user]
	if !ok {
		return nil, errors.New("no user wallet just recharge first")
	}

	if wallet < times {
		return nil, errors.New("point not enough")
	}
	service.Wallet[user] = wallet - times

	output := make([]gachapon.Item, times)
	for i := 0; i < times; i++ {
		output[i] = service.Gachapon.Random()
	}

	return output, nil
}

func (service *MemoryGachaponService) Recharge(user string, recharge int) int {
	wallet, ok := service.Wallet[user]
	if !ok {
		wallet = 0
	}

	result := wallet + recharge
	service.Wallet[user] = result

	return result
}
