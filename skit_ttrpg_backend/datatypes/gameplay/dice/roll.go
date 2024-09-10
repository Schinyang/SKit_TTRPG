package roll

import "math/rand"

func rollDice(sides int) int {
	return rand.Intn(sides) + 1
}

func RollMultiple(times int, sides int) (int, []int) {
	rolls := make([]int, times)
	result := 0
	for i := 0; i < times; i++ {
		roll := rollDice(sides)
		result += roll
		rolls[i] = roll
	}
	return result, rolls
}
