package game

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func Run() {
	heroes := make([]Hero, 30)
	rand.Seed(time.Now().UnixNano())
	CreateRandomHeroes32(&heroes)
	c0 := make(chan Hero)
	c1 := make(chan Hero)

	for len(heroes) != 1 {
		go ToFight(&heroes, c0, c1)
		go MakeFight(&heroes, c0, c1)
		time.Sleep(time.Millisecond)
	}
	fmt.Println("IN THIS FIGHT, ", heroes[0], " WON!!!")
}

func CreateRandomHeroes32(heroes *[]Hero) {
	rand.Seed(time.Now().UnixNano())
	warriors := 6 + rand.Intn(9)
	mages := 6 + rand.Intn(9)
	hunters := 32 - (warriors + mages)
	i := 0
	for warriors != 0 {
		(*heroes)[i] = &Warrior{Id: i, Name: "name" + strconv.FormatInt(int64(i), 10), Healths: 100, Power: 3 + rand.Intn(8), Defense: 3 + rand.Intn(8), Rage: 0}
		i++
		warriors--
	}
	for mages != 0 {
		(*heroes)[i] = &Mage{Id: i, Name: "name" + strconv.FormatInt(int64(i), 10), Healths: 100, Power: 4 + rand.Intn(7), Defense: 4 + rand.Intn(7), Mana: 0}
		i++
		mages--
	}
	for hunters != 0 {
		(*heroes)[i] = &Hunter{Id: i, Name: "name" + strconv.FormatInt(int64(i), 10), Healths: 100, Power: 5 + rand.Intn(6), Defense: 5 + rand.Intn(6), Energy: 0}
		i++
		hunters--
	}
}
