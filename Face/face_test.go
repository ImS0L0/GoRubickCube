package Face

import (
	"fmt"
	"testing"
)

func TestNewFace(t *testing.T) {
	var F Face

	width := uint8(3)

	F.New(&width, 1)

	fmt.Println("Face", F)
	fmt.Print("\n------\n")
}

func TestRotates(t *testing.T) {
	matrix := [][]uint8{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}

	F := Face(matrix)

	F.Rotate()
	fmt.Println("Rotate", F)

	F.Rotate()
	fmt.Println("Rotate", F)

	F.CounterRotate()
	fmt.Println("CounterRotate", F)

	fmt.Print("\n------\n")
}

func TestReplaceCol(t *testing.T) {
	matrix := [][]uint8{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}

	F := Face(matrix)

	F.RepalceColumn(1, []uint8{2, 2, 2})

	fmt.Println(F)

	F.RepalceColumn(2, []uint8{2, 2, 2})

	fmt.Println(F)

	F.RepalceColumn(0, []uint8{2, 2, 2})

	fmt.Println(F)
	fmt.Print("\n------\n")
}

func TestGetCol(t *testing.T) {
	matrix := [][]uint8{{6, 3, 1}, {7, 2, 2}, {4, 1, 3}}

	F := Face(matrix)

	matrixFace := [][]uint8{F.GetColumn(0), F.GetColumn(1), F.GetColumn(2)}

	result := [][]uint8{{6, 7, 4}, {3, 2, 1}, {1, 2, 3}}

	for i := 0; i < len(matrixFace); i++ {
		for j := 0; j < len(matrixFace[i]); j++ {
			if matrixFace[i][j] != result[i][j] {
				t.Error("[", i, "][", j, "]diferent", matrixFace[i][j], result[i][j])
			}
		}
	}

	fmt.Println("Done")

	fmt.Print("\n------\n")
}
