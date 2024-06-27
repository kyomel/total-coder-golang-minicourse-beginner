package main

import "fmt"

// 1.7 INTERFACES

// ABSTRACT TYPE
type Player interface {
	KickBall() int
}

// CONCRETE TYPE
type BadPlayer struct {
	Power int
}

type GoodPlayer struct {
	Power int
}

func (badPlayer BadPlayer) KickBall() int {
	// Will always execute logic:
	// 1.
	// 2.
	// 3.
	// fmt.Println("Bad player is kicking the ball with power", badPlayer.Power)
	return badPlayer.Power
}

func (goodPlayer GoodPlayer) KickBall() int {
	// Will always execute logic:
	// 1.
	// 2.
	// 3.
	multiplier := 10
	// fmt.Println("GOOD player is kicking the ball with power", goodPlayer.Power*multiplier)
	return goodPlayer.Power * multiplier
}

func main() {
	badPlayer := BadPlayer{
		Power: 1,
	}

	goodPlayer := GoodPlayer{
		Power: 10,
	}

	TeamPlayerKicksBall(badPlayer)
	TeamPlayerKicksBall(goodPlayer)
}

func TeamPlayerKicksBall(player Player) {
	player.KickBall()
	// switch p := player.(type) {
	// case BadPlayer:
	// 	fmt.Println("Bad player is kicking the ball with power", p.Power)
	// case GoodPlayer:
	// 	fmt.Println("Good player is kicking the ball with power", p.Power)
	// }

	if p, ok := player.(GoodPlayer); ok {
		fmt.Println("this is the good player attackPower", p.Power)
	}
}
