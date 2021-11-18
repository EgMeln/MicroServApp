package game

import (
	"fmt"
	"math/rand"
	"time"
)

func remove(s *[]Hero, i int) []Hero {
	(*s)[i] = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return *s
}

func Run() string {
	heroes := make([]Hero, 32)
	rand.Seed(time.Now().UnixNano())
	CreateRandomHeroes32(&heroes)

	c0 := make(chan Hero)
	c1 := make(chan Hero)

	for len(heroes) != 1 {
		go ToFight(&heroes, c0, c1)
		go MakeFight(&heroes, c0, c1)
	}
	fmt.Println("\nIN THIS FIGHT, ", heroes[0], " WON!!!")
	str := "IN THIS FIGHT, " + string(heroes[0].getName()) + " WON!!!"
	return str
}

func ToFight(heroes *[]Hero, downstream, downstream2 chan Hero) {
	if len(*heroes) == 0 || len(*heroes) == 1 {
		return
	}

	first := 0 + rand.Intn(len(*heroes))
	if first <= len(*heroes) {
		downstream <- (*heroes)[first]
		remove(heroes, first)
	} else {
		return
	}
	second := 0 + rand.Intn(len(*heroes))
	if second <= len(*heroes) {
		downstream2 <- (*heroes)[second]
		remove(heroes, second)
	} else {
		return
	}
}

func MakeFight(heroes *[]Hero, upstream, upstream2 chan Hero) {
	for v := range upstream {
		for p := range upstream2 {
			for {
				if v.amountStamina() > 20 {
					v.SecondSkill(p)
				} else if v.amountStamina() > 10 {
					v.FirstSkill(p)
				} else {
					choose := 1 + rand.Intn(2)
					if choose == 1 {
						fmt.Println(v, "Attack")
						v.Attack(p)
					} else {
						fmt.Println(v, "Defend")
						v.Defend()
					}
				}
				if p.Health() <= 0 {
					fmt.Println("In this battle winner ", v)
					fmt.Println("lose ", p)
					*heroes = append(*heroes, v)
					return
				}
				if p.amountStamina() > 20 {
					p.SecondSkill(v)
				} else if p.amountStamina() > 10 {
					p.FirstSkill(v)
				} else {
					choose := 1 + rand.Intn(2)
					if choose == 1 {
						fmt.Println(p, "Attack")
						p.Attack(v)
					} else {
						fmt.Println(p, "Defend")
						p.Defend()
					}
				}
				if v.Health() <= 0 {
					fmt.Println("In this battle winner ", p)
					fmt.Println("lose ", v)
					*heroes = append(*heroes, p)
					return
				}
			}
		}
	}
}
