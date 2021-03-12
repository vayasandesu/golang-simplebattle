package action

import components "simplebattle/combat/components"

// AttackCommand ... use for execute attack of Attacker to Target
type AttackCommand struct {
	Attacker *components.IAttacker
	Target   *components.IDamagable
}

// Execute ... execute action
func (cmd AttackCommand) Execute() {
	(*cmd.Attacker).Attack(*cmd.Target)
}
