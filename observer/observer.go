package observer

import (
	"simplebattle/utils"
)

type Character struct {
	Name                string
	HP                  int
	ATK                 int
	DEF                 int
	damageEventListener []chan DamageEvent // also use map[string][]chan TYPE to use event as key but in scenario the struct is match better
}

type DamageEvent struct {
	Attacked Character
	Attacker Character
	Damage   int
}

func EmitDamageEvent(listeners []chan DamageEvent, event DamageEvent) {
	for _, listener := range listeners {
		listener <- event
	}
}

func (character *Character) TakeDamage(attacker Character, damage int) {
	finalDamage := damage - character.DEF
	finalDamage = utils.Max(0, finalDamage)
	damageEvent := DamageEvent{
		Attacked: *character,
		Attacker: attacker,
		Damage:   finalDamage,
	}

	character.HP -= finalDamage

	EmitDamageEvent(character.damageEventListener, damageEvent)
}

func (character *Character) AddDamageEventListener(observer chan DamageEvent) {
	if character.damageEventListener == nil {
		character.damageEventListener = []chan DamageEvent{}
	}

	character.damageEventListener = append(character.damageEventListener, observer)
}

func (character *Character) RemoveDamageEventListener(observer chan DamageEvent) {
	if character.damageEventListener == nil {
		return
	}

	listeners := character.damageEventListener
	for i, lisenter := range listeners {
		if lisenter == observer {
			listeners = append(listeners[:i], listeners[i+1:]...)
		}
	}

	character.damageEventListener = listeners
}
