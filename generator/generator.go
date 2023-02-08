package generator

import (
	"math/rand"
)

type User struct {
	ID    int
	Name1 string
	Name2 string
	Name3 string
	Name4 string
}

func (User) TableName() string {
	return "Test"
}

const DATA_LIMIT = 10000

var Generator = func(c chan<- User) {
	for i := 0; i < DATA_LIMIT; i++ {
		alph := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
		rand.Shuffle(len(alph), func(i, j int) {
			alph[i], alph[j] = alph[j], alph[i]
		})
		u := User{
			ID:    rand.Int(),
			Name1: string(alph[:10]),
			Name2: string(alph[10:20]),
			Name3: string(alph[20:30]),
			Name4: string(alph[30:40]),
		}
		c <- u
	}
}
