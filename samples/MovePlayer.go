package main

import (
	"fmt"
	"github.com/martamius/mcpigo"
)

func main() {
	fmt.Printf("Connecting");
	minecraft := mcpigo.Start();
	
	fmt.Printf("Connected");
	minecraft.Chat.Message("Getting positon...")
	pos := minecraft.Player.Position()
	
	pos.Y += 50
	
	minecraft.Chat.Message("Changing position...")
	minecraft.Player.SetPosition(pos)

	//pos.X -= 5
	//id := mc.World.BlockAtPosition(pos)
	//fmt.Println(id)

	//mc.World.SetBlockAtCoordinates(x, y+5, z, 1)

	//pos1 := pos
	//pos2 := pos
	//pos2.X += 5

	//mc.World.SetBlocksInPositionRange(pos1, pos2, 1)
}
