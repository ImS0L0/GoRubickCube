package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/ImS0L0/GoRubikCube/Color"
	"github.com/ImS0L0/GoRubikCube/Cube"
)

func clearTerminal() {

	var Terminal *exec.Cmd

	if runtime.GOOS == "windows" {
		Terminal = exec.Command("cmd", "/c", "cls")
	} else {
		Terminal = exec.Command("clear")
	}
	Terminal.Stdout = os.Stdout
	Terminal.Run()
	fmt.Println()
}

func askQuestionY(s *bufio.Scanner, yes, no func()) {
	s.Scan()
	input := s.Text()

	if input == "" || strings.ToLower(input) == "y" {
		yes()
	} else {
		no()
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var cube Cube.Cube

	msg := ""

	for {
		clearTerminal()
		fmt.Printf(Color.Yellow, "WELCOME!!\n\n")
		fmt.Printf(Color.White, "Enter Cube size 2 to 12")
		fmt.Printf(Color.Red, " (type \"exit\" to leave)")
		fmt.Printf(Color.Red, msg)
		fmt.Printf(Color.Blue, "\n ƒ: ")
		scanner.Scan()

		input := scanner.Text()

		if strings.ToLower(input) == "exit" {
			fmt.Printf(Color.Yellow, "\n\n Good Bye ;)\n\n")
			break
		}

		if num, err := strconv.Atoi(input); err != nil {
			msg = " (\"" + input + "\" is not a number)"
			continue
		} else {
			cube.New(uint8(num))

			clearTerminal()
			cube.PrintCube()

			fmt.Printf(Color.White, "\nScramble Cube? (Y/n)")
			fmt.Printf(Color.Blue, "\n ƒ: ")

			askQuestionY(scanner, func() {
				clearTerminal()
				cube.PrintCube()

				fmt.Printf(Color.White, "\nScramble Size  Automatic? (Y/n)")
				fmt.Printf(Color.Blue, "\n ƒ: ")

				askQuestionY(scanner, func() {
					cube.AutoScramble()
				}, func() {
					msg = ""
					for {
						clearTerminal()
						cube.PrintCube()

						fmt.Printf(Color.White, "\nEnter Size")
						fmt.Printf(Color.Red, msg)
						fmt.Printf(Color.Blue, "\n ƒ: ")

						scanner.Scan()

						input = scanner.Text()

						if num, err = strconv.Atoi(input); err != nil {
							msg = " (\"" + input + "\" is not a number)"
							continue
						} else {
							cube.CreateScramble(uint8(num))
							break
						}
					}
				})
			}, func() {})

			for {
				clearTerminal()
				if len(cube.GetScramble()) > 0 {
					fmt.Printf(Color.Blue, "\n Scramble:  ")
					fmt.Printf(Color.Green, cube.GetScramble())
					fmt.Println()
				}
				cube.PrintCube()

				fmt.Printf(Color.White, "\nEnter Movements")
				fmt.Printf(Color.Red, "  (type \"exit\" to leave of Cube)")
				fmt.Printf(Color.Orange, "  (type \"help\" for show the help)")
				fmt.Printf(Color.Blue, "\n ƒ: ")

				scanner.Scan()

				input = scanner.Text()

				inputLowerCase := strings.ToLower(input)

				if inputLowerCase == "exit" {
					break
				} else if inputLowerCase == "help" {
					clearTerminal()
					fmt.Printf(Color.White, "Notation:\n\n")

					fmt.Printf(Color.White, "All notations usually have a capital letter (example: U, D, R, L ...) and each letter represents a movement:\n\n")

					fmt.Printf(Color.Green, "\tU = Up layer\n\tD = Down layer\n\tR = Right layer\n\tL = Left layer\n\tF = Front layer\n\tB = Back layer\n\n")

					fmt.Printf(Color.White, "Too exist middle movements\n\n")

					fmt.Printf(Color.Green, "\tE = between U and D (E follows D)\n\tM = between L and R (M follows L)\n\tS = between F and B (S follows F)\n\n")

					fmt.Printf(Color.White, "All this movements can move clockwise and counterclockwise, represented by a simple quotation mark, like this: U' F' L' ...\n\n")

					fmt.Printf(Color.White, "To this movements are called \"prime\" (or opposide) movements example: U' is U prime.\n\n")

					fmt.Printf(Color.White, "The \"prime\" movements are like the nomal movement, but on the contrary, namely, if U movement moves to the left, U Prime moves to the right.\n\n")

					fmt.Printf(Color.White, "And finally, the \"twice\" movements is like a normal movement but this be repeat 2 times, and this is representated like this U2 or F2 or R2 ...\n\n\n")

					fmt.Printf(Color.Green, " (Enter to quit) ")

					scanner.Scan()
				} else {
					arr := strings.Split(input, " ")
					cube.MoveCube(&arr)
				}
			}

		}
	}

}
