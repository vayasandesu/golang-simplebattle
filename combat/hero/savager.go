package hero

import (
	"fmt"
	components "simplebattle/combat/components"
	"simplebattle/utils"
)

type Savager struct {
	Name  string
	Hp    int
	Atk   int
	Def   int
	Stack int
}

// Attack ... may mutant character data then we use pointer reciever
func (character *Savager) Attack(target components.IDamagable) {
	finalAttack := character.Atk + character.Stack
	target.TakeDamage(character, finalAttack)
}

// TakeDamage ... may mutant character data then we use pointer reciever
func (character *Savager) TakeDamage(attacker components.IDamagable, damage int) {
	finalDamage := damage - character.Def
	finalDamage = utils.Max(finalDamage, 0)

	character.Hp -= finalDamage + int(character.Stack/2.0)
}

// Restore ... may mutant character data then we use pointer reciever
func (character *Savager) Restore(attacker components.IDamagable, damage int) {
	character.Hp += 2
	character.Stack = 0
}

// GetInformation ...  no need to mutant data then we use value reciever
func (character Savager) GetInformation() string {
	return fmt.Sprintf("%s\t | HP: %d, ATK: %d, DEF: %d",
		character.Name,
		character.Hp,
		character.Atk,
		character.Def)
}
