package action

import (
	components "simplebattle/combat/components"
)

// DefenseCommand ... use for increase self defense power
type DefenseCommand struct {
	Self *components.IDefender
}

// Execute ... execute action
func (cmd DefenseCommand) Execute() {
	(*cmd.Self).Defense()
}
