package services

import "simplebattle/gachapon"

type IGachaService interface {
	Roll(user string, times int) ([]gachapon.Item, error)
	Recharge(user string, recharge int) int
}
