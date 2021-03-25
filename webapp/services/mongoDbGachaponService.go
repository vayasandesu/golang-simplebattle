package services

import "simplebattle/gachapon"

type MongoDbGachaponService struct {
}

func (service *MongoDbGachaponService) Roll(user string, times int) ([]gachapon.Item, error) {

	return []gachapon.Item{}, nil
}

func (service *MongoDbGachaponService) Recharge(user string, recharge int) {

}
