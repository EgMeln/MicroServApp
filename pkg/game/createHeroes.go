package game

import (
	"math/rand"
	"strconv"
	"time"
)

func CreateRandomHeroes(heroes *[]Hero) {
	rand.Seed(time.Now().UnixNano())
	for i := range *heroes {
		num := rand.Intn(3)
		switch num {
		case 0:
			(*heroes)[i] = &Warrior{Id: i, Name: "name" + strconv.FormatInt(int64(i), 10), Healths: 100, Power: 3 + rand.Intn(8), Defense: 3 + rand.Intn(8), Rage: 0}
		case 1:
			(*heroes)[i] = &Mage{Id: i, Name: "name" + strconv.FormatInt(int64(i), 10), Healths: 100, Power: 4 + rand.Intn(7), Defense: 4 + rand.Intn(7), Mana: 0}

		case 2:
			(*heroes)[i] = &Hunter{Id: i, Name: "name" + strconv.FormatInt(int64(i), 10), Healths: 100, Power: 5 + rand.Intn(6), Defense: 5 + rand.Intn(6), Energy: 0}
		}
	}
}
