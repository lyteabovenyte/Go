package main

import (
	"fmt"
	"math/rand"
)

type Player interface {
	kickBall() (string, int)
}

type fotballPlayer struct {
	name          string
	shot, control int
	fame          int
}

type Ronaldo struct {
	name          string
	shot, control int
	fame          int
	number        int
}

type Messi struct {
	name          string
	shot, control int
	fame          int
	number        int
}

func (m Messi) kickBall() (string, int) {
	dmg := (m.control + m.shot) * (m.fame + m.number)
	return m.name, dmg
}

func (r Ronaldo) kickBall() (string, int) {
	dmg := (r.shot + r.control) * (r.fame + r.number)
	return r.name, dmg
}

func (p fotballPlayer) kickBall() (string, int) {
	dmg := (p.shot + p.control) * p.fame
	return p.name, dmg
}

func main() {
	team := make([]Player, 11)
	for i := 0; i < len(team)-2; i++ {
		team[i] = fotballPlayer{
			name:    "random",
			shot:    rand.Intn(10),
			control: rand.Intn(10),
			fame:    rand.Intn(10),
		}
	}
	team[len(team)-1] = Ronaldo{
		name:    "ronaldo",
		shot:    10,
		control: 9,
		fame:    10,
		number:  7,
	}
	team[len(team)-2] = Messi{
		name:    "messi",
		shot:    9,
		control: 10,
		fame:    10,
		number:  10,
	}

	for _, p := range team {
		name, dmg := p.kickBall()
		{
			fmt.Printf("%s kicking the ball: %v\n", name, dmg)
		}
	}
}
