package action

import (
	components "simplebattle/combat/components"
)

// RestoreCommand ... restore self hitpoint
type RestoreCommand struct {
	Self *components.IRestorer
}

// Execute ... execute action
func (cmd RestoreCommand) Execute() {
	(*cmd.Self).Restore()
}
