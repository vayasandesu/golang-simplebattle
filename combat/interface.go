package combat

/* Best practies : interface should suffix with -er */

type IDamagable interface {
	TakeDamage(attacker IDamagable, damage int)
}

type IDefender interface {
	Defense()
}

type IAttacker interface {
	Attack(target IDamagable)
}

type IRestorer interface {
	Restore()
}

type ILogable interface {
	GetInformation() string
}
