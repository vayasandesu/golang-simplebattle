package combat

import (
	"simplebattle/combat"
)

// DefenseCommand ... use for increase self defense power
type DefenseCommand struct {
	Self *combat.IDefender
}

// Execute ... execute action
func (cmd DefenseCommand) Execute() {
	(*cmd.Self).Defense()
}
