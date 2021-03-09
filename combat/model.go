package combat

import (
	"fmt"
	"simplebattle/utils"
)

type Character struct {
	Name string
	Hp   int
	Atk  int
	Def  int
}

func (character *Character) Attack(target ICombatable) {
	target.TakeDamage(character, character.Atk)
}

func (character *Character) TakeDamage(attacker ICombatable, damage int) {
	finalDamage := damage - character.Def
	finalDamage = utils.Max(finalDamage, 0)

	character.Hp -= finalDamage
}

func (character Character) GetInformation() string {
	return fmt.Sprintf("%s\t | HP: %d, ATK: %d, DEF: %d",
		character.Name,
		character.Hp,
		character.Atk,
		character.Def)
}
