package game

import (
	"math/rand"
	"strconv"
	"time"
)

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
func CreateRandomHeroes64(heroes *[]Hero) {
	rand.Seed(time.Now().UnixNano())
	warriors := 12 + rand.Intn(9)
	mages := 12 + rand.Intn(9)
	hunters := 64 - (warriors + mages)
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
func CreateRandomHeroes128(heroes *[]Hero) {
	rand.Seed(time.Now().UnixNano())
	warriors := 24 + rand.Intn(18)
	mages := 24 + rand.Intn(18)
	hunters := 128 - (warriors + mages)
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
