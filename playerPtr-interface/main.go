package main

import (
	"fmt"

	"golang.org/x/exp/rand"
)

type Player interface {
	KickBall() (name string, dmg int)
}

type footbalPlayer struct {
	name                string
	shot, control, fame int
}

type Ronaldo struct {
	name                        string
	shot, control, fame, number int
}

func (f *footbalPlayer) KickBall() (string, int) {
	dmg := (f.control + f.shot) * (f.fame)
	return f.name, dmg
}

func (r *Ronaldo) KickBall() (string, int) {
	dmg := (r.control + r.shot) * (r.fame + r.number)
	return r.name, dmg
}

func main() {
	team := make([]Player, 11)
	for i := 0; i < len(team)-1; i++ {
		team[i] = &footbalPlayer{
			name:    "random",
			shot:    rand.Intn(10),
			control: rand.Intn(10),
			fame:    rand.Intn(10),
		}
	}
	team[len(team)-1] = &Ronaldo{
		name:    "ronaldo",
		shot:    10,
		control: 9,
		fame:    10,
		number:  7,
	}

	for _, p := range team {
		name, dmg := p.KickBall()
		fmt.Printf("%s is kicking the ball: %v\n", name, dmg)
	}
}
