package combat

import (
	"fmt"
	"simplebattle/utils"
)

type Monster struct {
	Name string
	Hp   int
	Atk  int
	Def  int
}

// Attack ... may mutant character data then we use pointer reciever
func (character *Monster) Attack(target IDamagable) {
	target.TakeDamage(character, character.Atk)
}

// TakeDamage ... may mutant character data then we use pointer reciever
func (character *Monster) TakeDamage(attacker IDamagable, damage int) {
	finalDamage := damage - character.Def
	finalDamage = utils.Max(finalDamage, 0)

	character.Hp -= finalDamage
}

// GetInformation ...  no need to mutant data then we use value reciever
func (character Monster) GetInformation() string {
	return fmt.Sprintf("%s\t | HP: %d, ATK: %d, DEF: %d",
		character.Name,
		character.Hp,
		character.Atk,
		character.Def)
}
