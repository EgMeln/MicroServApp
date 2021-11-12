11package game

import (
	"fmt"
	"math/rand"
)

type Hero interface {
	Attack(hero Hero)
	Defend()
	FirstSkill(hero Hero)
	SecondSkill(hero Hero)
	Health() int
	setHealth(int)
	setDefense(int)
	amountStamina() int
	amountDefense() int
	getID() int
}

type Warrior struct {
	Id      int
	Name    string
	Healths int
	Power   int
	Defense int
	Rage    int
}
type Mage struct {
	Id      int
	Name    string
	Healths int
	Power   int
	Defense int
	Mana    int
}
type Hunter struct {
	Id      int
	Name    string
	Healths int
	Power   int
	Defense int
	Energy  int
}

func (war *Warrior) Attack(hero Hero) {
	war.Rage += 10
	var damage int
	if war.Power+(10+rand.Intn(11))-hero.amountDefense() <= 0 {
		damage = 0
		hero.setHealth(damage)
		hero.setDefense(war.Power + (10 + rand.Intn(11)))
	} else {
		damage = war.Power + (10 + rand.Intn(11)) - war.Defense
		hero.setHealth(damage)
		hero.setDefense(damage)
	}
	fmt.Println("You have dealt ", damage, " damage")
}
func (war *Warrior) Defend() {
	war.Rage += 5
	war.Defense = war.Defense + (8 + rand.Intn(9))
}
func (war *Warrior) FirstSkill(hero Hero) {
	if war.Rage < 10 {
		fmt.Println("You're not angry enough! Be angrier!!!Arrrrrr")
		return
	}
	fmt.Println("YOU want to be like your hero Dovakin, so you scream at the enemy. You're stunning him!")
	war.Power += 1
	war.Rage -= 10
	var damage int
	if war.Power+5+rand.Intn(6)-hero.amountDefense() <= 0 {
		damage = 0
		hero.setHealth(damage)
		hero.setDefense(war.Power + 5 + rand.Intn(6))
	} else {
		damage = war.Power + 5 + rand.Intn(6) - war.Defense
		hero.setHealth(damage)
		hero.setDefense(damage)
	}
	fmt.Println("You have dealt ", damage, " damage")
}
func (war *Warrior) SecondSkill(hero Hero) {
	if war.Rage < 20 {
		fmt.Println("You're not angry enough! Be angrier!!!Arrrr")
		return
	}
	fmt.Println("You take out the peasant pitchfork that you stole from your father on the farm and rush with it to the enemy, imagining that you are a Warsong")
	war.Power += 3
	war.Rage -= 20
	var damage int
	if war.Power+20+rand.Intn(21)-hero.amountDefense() <= 0 {
		damage = 0
		hero.setHealth(damage)
		hero.setDefense(war.Power + 20 + rand.Intn(21))
	} else {
		damage = war.Power + 20 + rand.Intn(21) - war.Defense
		hero.setHealth(damage)
		hero.setDefense(damage)
	}
	fmt.Println("You have dealt ", damage, " damage")

}
func (war *Warrior) Health() int {
	return war.Healths
}
func (war *Warrior) amountStamina() int {
	return war.Rage
}
func (war *Warrior) amountDefense() int {
	return war.Defense
}
func (war *Warrior) setHealth(points int) {
	war.Healths = war.Healths - points
}
func (war *Warrior) setDefense(point int) {
	if war.Defense-point <= 0 {
		war.Defense = 0
		return
	}
	war.Defense = war.Defense - point
}
func (war *Warrior) getID() int {
	return war.Id
}

func (mag *Mage) Attack(hero Hero) {
	mag.Mana += 3
	var damage int
	if mag.Power+7+rand.Intn(8)-hero.amountDefense() <= 0 {
		damage = 0
		hero.setHealth(damage)
		hero.setDefense(mag.Power + 7 + rand.Intn(8))
	} else {
		damage = mag.Power + 7 + rand.Intn(8) - mag.Defense
		hero.setHealth(damage)
		hero.setDefense(damage)
	}
	fmt.Println("You have dealt ", damage, " damage")

}
func (mag *Mage) Defend() {
	mag.Mana += 12
	mag.Defense = mag.Defense + (13 + rand.Intn(14))
}
func (mag *Mage) FirstSkill(hero Hero) {
	if mag.Mana < 10 {
		fmt.Println("You have too little mana. Draw it harder!")
		return
	}
	fmt.Println("You turn the enemy into the animal that he is most afraid of. And what is your surprise when the enemy became a sheep.The enemy is stunned.")
	mag.Power += 4
	mag.Mana -= 10
	var damage int
	if mag.Power+(4+rand.Intn(5))-hero.amountDefense() <= 0 {
		damage = 0
		hero.setHealth(damage)
		hero.setDefense(mag.Power + (4 + rand.Intn(5)))
	} else {
		damage = mag.Power + (4 + rand.Intn(5)) - mag.Defense
		hero.setHealth(damage)
		hero.setDefense(damage)
	}

	fmt.Println("You have dealt ", damage, " damage")

}
func (mag *Mage) SecondSkill(hero Hero) {
	if mag.Mana < 20 {
		fmt.Println("You have too little mana. Draw it harder!")
		return
	}
	fmt.Println("YYou summon a huge fireball as if you are from the Uchiha clan and throw it at the enemy")
	mag.Mana -= 20
	mag.Power += 3
	var damage int
	if mag.Power+(22+rand.Intn(23))-hero.amountDefense() <= 0 {
		damage = 0
		hero.setHealth(damage)
		hero.setDefense(mag.Power + (22 + rand.Intn(23)))
	} else {
		damage = mag.Power + (22 + rand.Intn(23)) - mag.Defense
		hero.setHealth(damage)
		hero.setDefense(damage)
	}
	fmt.Println("You have dealt ", damage, " damage")
}
func (mag *Mage) Health() int {
	return mag.Healths
}
func (mag *Mage) amountStamina() int {
	return mag.Mana
}
func (mag *Mage) amountDefense() int {
	return mag.Defense
}
func (mag *Mage) setHealth(points int) {
	mag.Healths = mag.Healths - points
}
func (mag *Mage) setDefense(point int) {
	if mag.Defense-point <= 0 {
		mag.Defense = 0
		return
	}
	mag.Defense = mag.Defense - point
}
func (mag *Mage) getID() int {
	return mag.Id
}

func (hun *Hunter) Attack(hero Hero) {
	hun.Energy += 12
	var damage int
	if hun.Power+(11+rand.Intn(12))-hero.amountDefense() <= 0 {
		damage = 0
		hero.setHealth(damage)
		hero.setDefense(hun.Power + (11 + rand.Intn(12)))
	} else {
		damage = hun.Power + (11 + rand.Intn(12)) - hun.Defense
		hero.setHealth(damage)
		hero.setDefense(damage)
	}
	fmt.Println("You have dealt ", damage, " damage")

}
func (hun *Hunter) Defend() {
	hun.Energy += 7
	hun.Defense += 10 + rand.Intn(11)
}
func (hun *Hunter) FirstSkill(hero Hero) {
	if hun.Energy < 10 {
		fmt.Println("You have too little Energy, you need to pet some animal urgently")
		return
	}
	fmt.Println("You send your pet to walk between the enemy's legs and rub against them.You stunned him.")
	hun.Power += 2
	hun.Energy -= 10
	var damage int
	if hun.Power+(7+rand.Intn(8))-hero.amountDefense() <= 0 {
		damage = 0
		hero.setHealth(damage)
		hero.setDefense(hun.Power + (7 + rand.Intn(8)))
	} else {
		damage = hun.Power + (7 + rand.Intn(8)) - hun.Defense
		hero.setHealth(damage)
		hero.setDefense(damage)
	}
	fmt.Println("You have dealt ", damage, " damage")
}
func (hun *Hunter) SecondSkill(hero Hero) {
	if hun.Energy < 20 {
		fmt.Println("You have too little Energy, you need to pet some animal urgently")
		return
	}
	fmt.Println("You imagine that you are a Rexar, and your cat is a bear. And gather all your strength in this shot!")
	hun.Energy -= 20
	hun.Power += 3
	var damage int
	if hun.Power+(25+rand.Intn(26))-hero.amountDefense() <= 0 {
		damage = 0
		hero.setHealth(damage)
		hero.setDefense(hun.Power + (25 + rand.Intn(26)))
	} else {
		damage = hun.Power + (25 + rand.Intn(26)) - hun.Defense
		hero.setHealth(damage)
		hero.setDefense(damage)
	}
	fmt.Println("You have dealt ", damage, " damage")
}
func (hun *Hunter) Health() int {
	return hun.Healths
}
func (hun *Hunter) amountStamina() int {
	return hun.Energy
}
func (hun *Hunter) amountDefense() int {
	return hun.Defense
}
func (hun *Hunter) setHealth(points int) {
	hun.Healths = hun.Healths - points
}
func (hun *Hunter) setDefense(point int) {
	if hun.Defense-point <= 0 {
		hun.Defense = 0
		return
	}
	hun.Defense = hun.Defense - point
}
func (hun *Hunter) getID() int {
	return hun.Id
}
