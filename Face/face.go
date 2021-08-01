package Face

type Face [][]uint8

func (F *Face) New(width *uint8, color uint8) {
	matrix := make([][]uint8, *width)

	for i := uint8(0); i < *width; i++ {
		matrix[i] = make([]uint8, *width)

		for j := uint8(0); j < *width; j++ {
			matrix[i][j] = color
		}
	}

	*F = Face(matrix)
}

func (F *Face) RepalceColumn(index *uint8, list []uint8) {
	matrix := [][]uint8(*F)

	listLen := len(list)

	for i := 0; i < listLen; i++ {
		matrix[i][*index] = list[i]
	}

	*F = Face(matrix)
}

func (F *Face) GetColumn(index *uint8) (list []uint8) {
	faceLen, matrix := len(*F), [][]uint8(*F)

	list = make([]uint8, faceLen)

	for i := 0; i < faceLen; i++ {
		list[i] = matrix[i][*index]
	}

	return list
}

func (F *Face) rotate() {
	oldMatrix := [][]uint8(*F)

	width := len(oldMatrix)

	newMatirx := make([][]uint8, width)

	for i := 0; i < width; i++ {
		newMatirx[i] = make([]uint8, width)
	}

	newFace := Face(newMatirx)

	i, j := 0, uint8(width-1)

	for i < width {
		newFace.RepalceColumn(&j, oldMatrix[i])
		i++
		j--
	}

	*F = newFace
}

func (F *Face) primeRotate() {
	oldMatrix := [][]uint8(*F)

	width := uint8(len(oldMatrix))

	newMatirx := make([][]uint8, width)

	for i := uint8(0); i < width; i++ {
		newMatirx[i] = make([]uint8, width)
	}

	newFace := Face(newMatirx)

	for i := uint8(0); i < width; i++ {
		newFace.RepalceColumn(&i, InvertList(oldMatrix[i]))
	}

	*F = newFace
}

func (F *Face) RotateTo(prime *bool) {
	if *prime {
		F.primeRotate()
	} else {
		F.rotate()
	}
}

func InvertList(oldList []uint8) []uint8 {
	listLen := len(oldList)
	newList := make([]uint8, listLen)

	for i, j := 0, listLen-1; i < listLen; i++ {
		newList[i] = oldList[j]
		j--
	}

	return newList
}
