package main

import "fmt"

/*
There is a new mobile game that starts with consecutively numbered clouds. Some of the clouds are thunderheads and others are cumulus. The player can jump on any cumulus cloud having a number that is equal to the number of the current cloud plus 1 or 2. The player must avoid the thunderheads. Determine the minimum number of jumps it will take to jump from the starting postion to the last cloud. It is always possible to win the game.

For each game, you will get an array of clouds numbered 0 if they are safe or 1 if they must be avoided.
*/

func main() {
	data := []int{0, 0, 0, 0, 1, 0}
	fmt.Println(jumpingOnCloud(data))
}

//input = binary integers i.e 0 or 1.returns the minimum no. of jumps
func jumpingOnCloud(clouds []int) int {

	var jump int
	for i := 0; i < len(clouds)-1; i++ {
		if clouds[i] == 0 {
			i++
		}
		jump++
	}
	return jump
}
