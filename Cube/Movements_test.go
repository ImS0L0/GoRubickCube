package Cube

import (
	"fmt"
	"testing"
)

func init() {
	C.New(10)
}

func TestMove(t *testing.T) {

	C.AutoScramble()
	C.PrintCube()

	fmt.Println(C.scramble)

	// // scramble : U R U' R' D L D' L' F R F' L B U B'

	// fmt.Println("U")
	// C.MoveYLayer(0, false)

	// fmt.Println("R")
	// C.MoveXLayer(2, false)

	// fmt.Println("U'")
	// C.MoveYLayer(0, true)

	// fmt.Println("R'")
	// C.MoveXLayer(2, true)

	// fmt.Println("D")
	// C.MoveYLayer(2, true)

	// fmt.Println("L")
	// C.MoveXLayer(0, true)

	// fmt.Println("D'")
	// C.MoveYLayer(2, false)

	// fmt.Println("L'")
	// C.MoveXLayer(0, false)

	// fmt.Println("F")
	// C.MoveZLayer(2, false)

	// fmt.Println("R")
	// C.MoveXLayer(2, false)

	// fmt.Println("F'")
	// C.MoveZLayer(2, true)

	// fmt.Println("L")
	// C.MoveXLayer(0, true)

	// fmt.Println("B")
	// C.MoveZLayer(0, true)

	// fmt.Println("U")
	// C.MoveYLayer(0, false)

	// fmt.Println("B'")
	// C.MoveZLayer(0, false)
	// C.PrintCube()
}

// --- Y

// U (0, false)
// U' (0, true)

// E (1, true)
// E' (1, false)

// D (2, true)
// D' (2, false)

// --- X

// R (2, false)
// R' (2, true)

// M (1, true)
// M' (1, false)

// L (0, true)
// L'(0, false)

// --- Z

// F (2, false)
// F' (2, true)

// S (1, true)
// S' (1, false)

// B (0, true)
// B' (0, false)
