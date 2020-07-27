package main

import (
	"mygo/player"
	"mygo/referee"
)

func main() {
	Grade,Poker,Array:=player.Player("小明")
	Grade2,Poker2,Array2:=player.Player("小刚")
	referee.Referee(Grade,Poker,Array,Grade2,Poker2,Array2)
}
