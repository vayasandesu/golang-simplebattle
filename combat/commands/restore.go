package combat

import (
	"simplebattle/combat"
)

// RestoreCommand ... restore self hitpoint
type RestoreCommand struct {
	Self *combat.IRestorer
}

// Execute ... execute action
func (cmd RestoreCommand) Execute() {
	(*cmd.Self).Restore()
}
