package game

import (
	"math/rand"
	"time"
)

func createAndShowAllHeroes(heroes *[]Hero) string {
	CreateRandomHeroes(heroes)
	var str string
	for _, v := range *heroes {
		str = v.PrintListHeroes() + "\n"
	}
	return str
}
func Run32() string {
	heroes := make([]Hero, 32)
	roundHeroes := make([]Hero, 32)
	var resultStr string
	rand.Seed(time.Now().UnixNano())
	//CreateRandomHeroes(&heroes)
	createAndShowAllHeroes(&heroes)
	c0 := make(chan Hero)
	c1 := make(chan Hero)
	c3 := make(chan Hero)
	for len(heroes) != 1 {
		roundHeroes = nil
		for i := 0; i < len(heroes)/2; i++ {
			go ToFight(i, &heroes, c0, c1)
			go MakeFight(&resultStr, c0, c1, c3)
		}
		for i := 0; i < len(heroes)/2; i++ {
			roundHeroes = append(roundHeroes, <-c3)
		}
		heroes = roundHeroes
	}
	resultStr = resultStr + "IN THIS FIGHT, " + heroes[0].getName() + " WON!!!"
	return resultStr
}

func Run64() string {
	heroes := make([]Hero, 64)
	roundHeroes := make([]Hero, 64)
	var resultStr string
	rand.Seed(time.Now().UnixNano())
	CreateRandomHeroes(&heroes)
	c0 := make(chan Hero)
	c1 := make(chan Hero)
	c3 := make(chan Hero)
	for len(heroes) != 1 {
		roundHeroes = nil
		for i := 0; i < len(heroes)/2; i++ {
			go ToFight(i, &heroes, c0, c1)
			go MakeFight(&resultStr, c0, c1, c3)
		}
		for i := 0; i < len(heroes)/2; i++ {
			roundHeroes = append(roundHeroes, <-c3)
		}
		heroes = roundHeroes
	}
	resultStr = resultStr + "IN THIS FIGHT, " + heroes[0].getName() + " WON!!!"
	return resultStr
}
func Run128() string {
	heroes := make([]Hero, 128)
	roundHeroes := make([]Hero, 128)
	var resultStr string
	rand.Seed(time.Now().UnixNano())
	CreateRandomHeroes(&heroes)
	c0 := make(chan Hero)
	c1 := make(chan Hero)
	c3 := make(chan Hero)
	for len(heroes) != 1 {
		roundHeroes = nil
		for i := 0; i < len(heroes)/2; i++ {
			go ToFight(i, &heroes, c0, c1)
			go MakeFight(&resultStr, c0, c1, c3)
		}
		for i := 0; i < len(heroes)/2; i++ {
			roundHeroes = append(roundHeroes, <-c3)
		}
		heroes = roundHeroes
	}
	resultStr = resultStr + "IN THIS FIGHT, " + heroes[0].getName() + " WON!!!"
	return resultStr
}

func ToFight(i int, heroes *[]Hero, downstream, downstream2 chan Hero) {
	first := 2 * i
	if first < len(*heroes) && (len(*heroes) != 1 || len(*heroes) != 0) {
		downstream <- (*heroes)[first]
	} else {
		return
	}
	second := 2*i + 1
	if second < len(*heroes) && (len(*heroes) != 1 || len(*heroes) != 0) {
		downstream2 <- (*heroes)[second]
	} else {
		return
	}
}

func MakeFight(resultStr *string, upstream, upstream2, downstream chan Hero) {
	for v := range upstream {
		for p := range upstream2 {
			*resultStr = *resultStr + "This battle between " + v.getName() + " and " + p.getName() + "\n"
			for {
				if v.amountStamina() > 20 {
					*resultStr = v.SecondSkill(p) + "\n"
				} else if v.amountStamina() > 10 {
					*resultStr = v.FirstSkill(p) + "\n"
				} else {
					choose := 1 + rand.Intn(2)
					if choose == 1 {
						*resultStr = *resultStr + v.getName() + " Attack\n"
						*resultStr = v.Attack(p) + "\n"
					} else {
						*resultStr = *resultStr + v.getName() + " Defend\n"
						*resultStr = v.Defend() + "\n"
					}
				}
				if p.Health() <= 0 {
					*resultStr = *resultStr + "In this battle win " + v.getName() + "\n"
					*resultStr = *resultStr + "In this battle lose " + p.getName() + "\n"
					downstream <- v
					return
				}
				if p.amountStamina() > 20 {
					*resultStr = p.SecondSkill(v) + "\n"
				} else if p.amountStamina() > 10 {
					*resultStr = p.FirstSkill(v) + "\n"
				} else {
					choose := 1 + rand.Intn(2)
					if choose == 1 {
						*resultStr = *resultStr + p.getName() + " Attack\n"
						*resultStr = p.Attack(v) + "\n"
					} else {
						*resultStr = *resultStr + p.getName() + " Defend\n"
						*resultStr = p.Defend() + "\n"
					}
				}
				if v.Health() <= 0 {
					*resultStr = *resultStr + " In this battle win " + p.getName() + "\n"
					*resultStr = *resultStr + " In this battle win " + v.getName() + "\n"
					downstream <- p
					return
				}
			}
		}
	}
}
