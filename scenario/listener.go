package scenario

import (
	"fmt"
	"math/rand"
	"simplebattle/observer"
)

func createCharacter(name string, hp int, atk int, def int) observer.Character {
	character := observer.Character{
		Name: name,
		HP:   hp,
		ATK:  atk,
		DEF:  def,
	}

	return character
}

func displayDamage(channel chan observer.DamageEvent) {
	for {
		event := <-channel
		fmt.Printf("!> [%s] deals [%d] damage to [%s]\n",
			event.Attacker.Name,
			event.Damage,
			event.Attacked.Name)
	}
}

func displayHealth(channel chan observer.DamageEvent) {
	for {
		event := <-channel
		fmt.Printf("!> [%s] attacked by [%s] and has remaining [%d] hp\n",
			event.Attacked.Name,
			event.Attacker.Name,
			event.Attacked.HP)
	}
}

func addHealthDisplayListener(characters ...*observer.Character) {
	channel := make(chan observer.DamageEvent)
	size := len(characters)

	for i := 0; i < size; i++ {
		characters[i].AddDamageEventListener(channel)
	}

	go displayHealth(channel)
}

func addDamageDisplayListener(characters ...*observer.Character) {
	channel := make(chan observer.DamageEvent)
	size := len(characters)
	for i := 0; i < size; i++ {
		characters[i].AddDamageEventListener(channel)
	}

	go displayDamage(channel)
}

func autoAttack(self *observer.Character, onDone chan bool, enemies ...*observer.Character) {
	for self.HP > 0 && len(enemies) > 0 {
		enemiesLen := len(enemies)
		index := rand.Intn(enemiesLen)

		target := enemies[index]
		//fmt.Printf("ADDRESS target : %p\n", target)
		//fmt.Printf("ADDRESS enemies index : %p\n", &enemies[index])
		fmt.Printf("%s [%d] attack %s\n", self.Name, self.HP, target.Name)
		target.TakeDamage(*self, self.ATK)
		if target.HP <= 0 {
			enemies = append(enemies[:index], enemies[index+1:]...)
		}
		//time.Sleep(1000 * time.Millisecond)
	}
	if len(enemies) == 0 {
		fmt.Printf("%s has no target anymore\n", self.Name)
	} else {
		fmt.Printf("%s is dead\n", self.Name)
	}
	onDone <- true
}

func EventListenerRunner() {
	player := createCharacter("Hero", 100, 5, 2)
	monsterA := createCharacter("Slime", 10, 5, 3)
	monsterB := createCharacter("Bean", 10, 4, 2)
	monsterC := createCharacter("Cacao", 10, 3, 1)

	addDamageDisplayListener(&player, &monsterA, &monsterB, &monsterC)
	addHealthDisplayListener(&player)

	// number of character -1
	alivePlayer := 4
	deadChan := make(chan bool, 3)
	go autoAttack(&player, deadChan, &monsterA, &monsterB, &monsterC)
	go autoAttack(&monsterA, deadChan, &player)
	go autoAttack(&monsterB, deadChan, &player)
	go autoAttack(&monsterC, deadChan, &player)

	for alivePlayer > 1 {
		<-deadChan
		alivePlayer--
	}

	fmt.Println("END game")
}
