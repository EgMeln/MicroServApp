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
func Run32() string {
	heroes := make([]Hero, 32)
	var resultStr string
	rand.Seed(time.Now().UnixNano())
	CreateRandomHeroes32(&heroes)
	c0 := make(chan Hero)
	c1 := make(chan Hero)

	for len(heroes) != 1 {
		go ToFight(&heroes, c0, c1)
		go MakeFight(&heroes, &resultStr, c0, c1)
		time.Sleep(time.Millisecond)
	}
	fmt.Println("IN THIS FIGHT, ", heroes[0], " WON!!!")
	resultStr = resultStr + "IN THIS FIGHT, " + string(heroes[0].getName()) + " WON!!!"
	return resultStr
}
func Run64() string {
	heroes := make([]Hero, 64)
	var resultStr string
	rand.Seed(time.Now().UnixNano())
	CreateRandomHeroes64(&heroes)
	c0 := make(chan Hero)
	c1 := make(chan Hero)

	for len(heroes) != 1 {
		go ToFight(&heroes, c0, c1)
		go MakeFight(&heroes, &resultStr, c0, c1)
		time.Sleep(time.Millisecond)
	}
	fmt.Println("IN THIS FIGHT, ", heroes[0], " WON!!!")
	resultStr = resultStr + "IN THIS FIGHT, " + string(heroes[0].getName()) + " WON!!!"
	return resultStr
}
func Run128() string {
	heroes := make([]Hero, 128)
	var resultStr string
	rand.Seed(time.Now().UnixNano())
	CreateRandomHeroes128(&heroes)
	c0 := make(chan Hero)
	c1 := make(chan Hero)

	for len(heroes) != 1 {
		go ToFight(&heroes, c0, c1)
		go MakeFight(&heroes, &resultStr, c0, c1)
		time.Sleep(time.Millisecond)
	}
	fmt.Println("IN THIS FIGHT, ", heroes[0], " WON!!!")
	resultStr = resultStr + "IN THIS FIGHT, " + string(heroes[0].getName()) + " WON!!!"
	return resultStr
}

func ToFight(heroes *[]Hero, downstream, downstream2 chan Hero) {
	if len(*heroes) == 0 || len(*heroes) == 1 {
		return
	}

	first := 0 + rand.Intn(len(*heroes))
	downstream <- (*heroes)[first]
	remove(heroes, first)

	second := 0 + rand.Intn(len(*heroes))
	downstream2 <- (*heroes)[second]
	remove(heroes, second)
}

func MakeFight(heroes *[]Hero, resultStr *string, upstream, upstream2 chan Hero) {
	for v := range upstream {
		for p := range upstream2 {
			for {
				*resultStr = *resultStr + "This battle between " + string(v.getName()) + " and " + string(p.getName()) + "\n"
				if v.amountStamina() > 20 {
					v.SecondSkill(p)
				} else if v.amountStamina() > 10 {
					v.FirstSkill(p)
				} else {
					choose := 1 + rand.Intn(2)
					if choose == 1 {
						fmt.Println(v, "Attack")
						v.Attack(p)
						*resultStr = *resultStr + string(v.getName()) + " Attack\n"
					} else {
						fmt.Println(v, "Defend")
						*resultStr = *resultStr + string(v.getName()) + " Defend\n"
						v.Defend()
					}
				}
				if p.Health() <= 0 {
					fmt.Println("In this battle winner ", v)
					fmt.Println("lose ", p)
					*resultStr = *resultStr + "In this battle win " + string(v.getName()) + "\n"
					*resultStr = *resultStr + "In this battle lose " + string(p.getName()) + "\n"
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
						*resultStr = *resultStr + string(p.getName()) + " Attack\n"
						p.Attack(v)
					} else {
						fmt.Println(p, "Defend")
						*resultStr = *resultStr + string(p.getName()) + " Defend\n"
						p.Defend()
					}
				}
				if v.Health() <= 0 {
					fmt.Println("In this battle winner ", p)
					fmt.Println("lose ", v)
					*resultStr = *resultStr + " In this battle win " + string(p.getName()) + "\n"
					*resultStr = *resultStr + " In this battle win " + string(v.getName()) + "\n"
					*heroes = append(*heroes, p)
					return
				}
			}
		}
	}
}
