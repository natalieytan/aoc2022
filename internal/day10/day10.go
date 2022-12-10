package day10

import (
	"fmt"
	"strings"
)

const startingRegister = 1
const crtWidth = 40
const litPixel = "#"
const darkPixel = "."

func Part1(operations []Operation) int {
	cylesToTallySignalStrength := map[int]bool{
		20:  true,
		60:  true,
		100: true,
		140: true,
		180: true,
		220: true,
	}
	totalSignalStrength := 0
	updateTotalSignalStrength := func(register int, cyclesCompleted int) {
		if cylesToTallySignalStrength[cyclesCompleted] {
			totalSignalStrength += cyclesCompleted * register
		}
	}

	runOperationsWithEffect(operations, updateTotalSignalStrength)
	return totalSignalStrength
}

func Part2(operations []Operation) []string {
	crtScreen := []string{}
	currentCrtRow := make([]string, crtWidth)
	updateCrtScreen := func(register int, cyclesCompleted int) {
		crtIndex := (cyclesCompleted - 1) % crtWidth
		if crtIndex >= register-1 && crtIndex <= register+1 {
			currentCrtRow[crtIndex] = litPixel
		} else {
			currentCrtRow[crtIndex] = darkPixel
		}
		if crtIndex == crtWidth-1 {
			crtScreen = append(crtScreen, strings.Join(currentCrtRow, ""))
			currentCrtRow = make([]string, crtWidth)
		}
	}

	runOperationsWithEffect(operations, updateCrtScreen)
	return crtScreen
}

func PrintPart2(operations []Operation) {
	crtScreen := Part2(operations)
	for _, crtRow := range crtScreen {
		fmt.Printf("%+v\n", crtRow)
	}
}

func runOperationsWithEffect(operations []Operation, effect func(register int, cyclesCompleted int)) {
	register := startingRegister
	cyclesCompleted := 0

	for _, operation := range operations {
		for cyclesForOperation := 1; cyclesForOperation <= operation.cyclesToComplete(); cyclesForOperation++ {
			cyclesCompleted += 1

			effect(register, cyclesCompleted)

			if operation.operationType == Addx && cyclesForOperation == operation.cyclesToComplete() {
				register += operation.value
			}
		}
	}
}
