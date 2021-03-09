package combat

type ICombatable interface {
	Attack(target ICombatable)
	TakeDamage(attacker ICombatable, damage int)
}

type ILogable interface {
	GetInformation() string
}
