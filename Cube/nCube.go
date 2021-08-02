package Cube

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ImS0L0/GoRubikCube/Color"
	"github.com/ImS0L0/GoRubikCube/Face"
)

type Cube struct {
	width    uint8
	scramble []string

	faceUp    Face.Face
	faceDown  Face.Face
	faceRight Face.Face
	faceLeft  Face.Face
	faceFront Face.Face
	faceBack  Face.Face
	voidFace  Face.Face
}

func (C *Cube) New(width uint8) {

	if width > 12 {
		C.width = 12
	} else if width <= 0 {
		C.width = 2
	} else {
		C.width = width
	}

	C.voidFace.New(&C.width, 0)
	C.faceUp.New(&C.width, 1)
	C.faceDown.New(&C.width, 2)
	C.faceLeft.New(&C.width, 3)
	C.faceRight.New(&C.width, 4)
	C.faceFront.New(&C.width, 5)
	C.faceBack.New(&C.width, 6)
}

func (C *Cube) GetScramble() []string {
	return C.scramble
}

func (C *Cube) PrintCube() {
	cube := [][]Face.Face{
		{C.voidFace, C.faceUp},
		{C.faceLeft, C.faceFront, C.faceRight, C.faceBack},
		{C.voidFace, C.faceDown}}

	const cubeLen = 3

	for faceList := 0; faceList < cubeLen; faceList++ {
		faceHeight := C.width // len(cube[0][0])

		for listFromFace := uint8(0); listFromFace < faceHeight; listFromFace++ {
			listFromCubeLen := len(cube[faceList])

			for face := 0; face < listFromCubeLen; face++ {
				listLen := C.width

				for square := uint8(0); square < listLen; square++ {
					i := cube[faceList][face][listFromFace][square]
					fmt.Printf(Color.Arr[i], " ■")
				}
				fmt.Print("  ")
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func (C *Cube) AutoScramble() {
	scrambleLen := 11

	rand.Seed(time.Now().Unix())

	if C.width > 2 {
		scrambleLen = rand.Intn(4) + 18
	}

	if C.width > 3 && C.width < 12 {
		scrambleLen = 0
		for i := uint8(0); i < C.width; i++ {
			scrambleLen += rand.Intn(1+int(i)) + (9 + int(i))
		}
	}

	if C.width > 11 {
		scrambleLen = 205
	}

	C.CreateScramble(uint8(scrambleLen))
}

// 1x1x1 >=

// X Y Z
//

// 2x2x2 >=

// U D R L F B
//
// U' D' R' L' F' B'
//

// 3x3x3 >=

// Uw Dw Rw Lw Fw Bw
//
// Uw' Dw' Rw' Lw' Fw' Bw'
//

// M E S
//
// M' E' S'
//
