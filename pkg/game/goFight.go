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

	//var wg sync.WaitGroup

	for len(heroes) != 1 {
		go ToFight(&heroes, c0, c1)
		go MakeFight(&heroes, c0, c1)
		//time.Sleep(time.Millisecond)
	}
	//fmt.Println("IN THIS FIGHT, ", heroes[0], " WON!!!")
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
		go MakeFight(&heroes, c0, c1)
	}
	//fmt.Println("IN THIS FIGHT, ", heroes[0], " WON!!!")
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
		go MakeFight(&heroes, c0, c1)
		time.Sleep(time.Millisecond)
	}
	//fmt.Println("IN THIS FIGHT, ", heroes[0], " WON!!!")
	resultStr = resultStr + "IN THIS FIGHT, " + string(heroes[0].getName()) + " WON!!!"
	return resultStr
}

func ToFight(heroes *[]Hero, downstream, downstream2 chan Hero) {
	if len(*heroes) == 0 || len(*heroes) == 1 {
		return
	}

	first := 0 + rand.Intn(len(*heroes))
	second := 0 + rand.Intn(len(*heroes))
	if first <= len(*heroes) && second <= len(*heroes) {
		downstream <- (*heroes)[first]
		downstream2 <- (*heroes)[second]
		remove(heroes, first)
		remove(heroes, second)
	} else {
		return
	}
	//if second <= len(*heroes) {
	//	downstream2 <- (*heroes)[second]
	//	remove(heroes, second)
	//} else {
	//	return
	//}
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

//func ToFight(heroes *[]Hero, downstream, downstream2 chan Hero) {
//	if len(*heroes) == 0 || len(*heroes) == 1 {
//		return
//	}
//
//	first := 0 + rand.Intn(len(*heroes))
//	if first <= len(*heroes) {
//		downstream <- (*heroes)[first]
//		remove(heroes, first)
//	} else {
//		return
//	}
//	second := 0 + rand.Intn(len(*heroes))
//	if second <= len(*heroes) {
//		downstream2 <- (*heroes)[second]
//		remove(heroes, second)
//	} else {
//		return
//	}
//}
//
//func MakeFight(heroes *[]Hero, resultStr *string, upstream, upstream2 chan Hero) {
//	for v := range upstream {
//		for p := range upstream2 {
//			*resultStr = *resultStr + "This battle between " + string(v.getName()) + " and " + string(p.getName()) + "\n"
//			for {
//				if v.amountStamina() > 20 {
//					v.SecondSkill(p)
//					*resultStr = v.SecondSkill(p) + "\n"
//				} else if v.amountStamina() > 10 {
//					v.FirstSkill(p)
//					*resultStr = v.FirstSkill(p) + "\n"
//				} else {
//					choose := 1 + rand.Intn(2)
//					if choose == 1 {
//						//fmt.Println(v, "Attack")
//						*resultStr = *resultStr + string(v.getName()) + " Attack\n"
//						v.Attack(p)
//						*resultStr = v.Attack(p) + "\n"
//					} else {
//						//fmt.Println(v, "Defend")
//						*resultStr = *resultStr + string(v.getName()) + " Defend\n"
//						v.Defend()
//						*resultStr = v.Defend() + "\n"
//					}
//				}
//				if p.Health() <= 0 {
//					//fmt.Println("In this battle winner ", v)
//					//fmt.Println("lose ", p)
//					*resultStr = *resultStr + "In this battle win " + string(v.getName()) + "\n"
//					*resultStr = *resultStr + "In this battle lose " + string(p.getName()) + "\n"
//					*heroes = append(*heroes, v)
//					return
//				}
//				if p.amountStamina() > 20 {
//					p.SecondSkill(v)
//					*resultStr = p.SecondSkill(v) + "\n"
//				} else if p.amountStamina() > 10 {
//					p.FirstSkill(v)
//					*resultStr = p.FirstSkill(v) + "\n"
//				} else {
//					choose := 1 + rand.Intn(2)
//					if choose == 1 {
//						//fmt.Println(p, "Attack")
//						*resultStr = *resultStr + string(p.getName()) + " Attack\n"
//						p.Attack(v)
//						*resultStr = p.Attack(v) + "\n"
//					} else {
//						//fmt.Println(p, "Defend")
//						*resultStr = *resultStr + string(p.getName()) + " Defend\n"
//						p.Defend()
//						*resultStr = p.Defend() + "\n"
//					}
//				}
//				if v.Health() <= 0 {
//					//fmt.Println("In this battle winner ", p)
//					//fmt.Println("lose ", v)
//					*resultStr = *resultStr + " In this battle win " + string(p.getName()) + "\n"
//					*resultStr = *resultStr + " In this battle win " + string(v.getName()) + "\n"
//					*heroes = append(*heroes, p)
//					return
//				}
//			}
//		}
//	}
//}
