package combat

import (
	"fmt"
	"simplebattle/utils"
)

type Character struct {
	Name   string
	Hp     int
	Atk    int
	Def    int
	Immune bool
}

// Attack ... may mutant character data then we use pointer reciever
func (character *Character) Attack(target IDamagable) {
	target.TakeDamage(character, character.Atk)
}

// TakeDamage ... may mutant character data then we use pointer reciever
func (character *Character) TakeDamage(attacker IDamagable, damage int) {
	finalDamage := damage - character.Def
	finalDamage = utils.Max(finalDamage, 0)
	if character.Immune && finalDamage > 0 {
		finalDamage = 1
		character.Immune = false
	}

	character.Hp -= finalDamage
	if finalDamage > 0 {
		character.Immune = true
	}
}

// GetInformation ...  no need to mutant data then we use value reciever
func (character Character) GetInformation() string {
	return fmt.Sprintf("%s\t | HP: %d, ATK: %d, DEF: %d",
		character.Name,
		character.Hp,
		character.Atk,
		character.Def)
}
