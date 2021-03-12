package hero

import (
	"fmt"
	components "simplebattle/combat/components"
	"simplebattle/utils"
)

type Tanker struct {
	Name   string
	Hp     int
	Atk    int
	Def    int
	Immune bool
}

// Attack ... may mutant character data then we use pointer reciever
func (character *Tanker) Attack(target components.IDamagable) {
	target.TakeDamage(character, character.Atk)
}

// Attack ... may mutant character data then we use pointer reciever
func (character *Tanker) Defense() {
	character.Immune = true
	fmt.Println("cast defense")
}

// TakeDamage ... may mutant character data then we use pointer reciever
func (character *Tanker) TakeDamage(attacker components.IDamagable, damage int) {
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
func (character Tanker) GetInformation() string {
	return fmt.Sprintf("%s\t | HP: %d, ATK: %d, DEF: %d, immune: %t",
		character.Name,
		character.Hp,
		character.Atk,
		character.Def,
		character.Immune)
}
