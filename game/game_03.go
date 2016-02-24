package main

import tl "github.com/JoelOtter/termloop"

func main() {
	game := tl.NewGame()
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorGreen,
		Fg: tl.ColorBlack,
		Ch: 'v',
	})
	lake := tl.NewRectangle(10, 10, 50, 20, tl.ColorBlue)
	level.AddEntity(lake)
	game.Screen().SetLevel(level)
	game.Start()
}
