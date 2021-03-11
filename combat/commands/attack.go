package command

import "simplebattle/combat"

// AttackCommand ... use for execute attack of Attacker to Target
type AttackCommand struct {
	Attacker *combat.IAttacker
	Target   *combat.IDamagable
}

// Execute ... execute action
func (cmd AttackCommand) Execute() {
	(*cmd.Attacker).Attack(*cmd.Target)
}
