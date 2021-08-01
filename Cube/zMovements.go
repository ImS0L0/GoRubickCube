package Cube

import (
	"github.com/ImS0L0/GoRubikCube/Face"
)

func (C *Cube) moveZLayersTo() {
	upList := C.faceUp[Layer]

	counterLayer := uint8((C.width - Layer) - 1)

	if Prime {
		C.faceUp[Layer] = C.faceRight.GetColumn(&counterLayer)
		C.faceRight.RepalceColumn(&counterLayer, Face.InvertList(C.faceDown[counterLayer]))
		C.faceDown[counterLayer] = C.faceLeft.GetColumn(&Layer)
		C.faceLeft.RepalceColumn(&Layer, Face.InvertList(upList))
	} else {
		C.faceUp[Layer] = Face.InvertList(C.faceLeft.GetColumn(&Layer))
		C.faceLeft.RepalceColumn(&Layer, C.faceDown[counterLayer])
		C.faceDown[counterLayer] = Face.InvertList(C.faceRight.GetColumn(&counterLayer))
		C.faceRight.RepalceColumn(&counterLayer, upList)
	}
}

func (C *Cube) rotateZFace() {
	if Layer == 0 {
		C.faceBack.RotateTo(&noPrime)
	} else if Layer == C.width-1 {
		C.faceFront.RotateTo(&Prime)
	}
}

func (C *Cube) MoveZLayer(layer, times *uint8, primeRotation *bool) {
	Layer, Prime, noPrime = *layer, *primeRotation, !*primeRotation

	for i := uint8(0); i < *times; i++ {
		C.rotateZFace()
		C.moveZLayersTo()
	}
}

func (C *Cube) MoveZwLayer(layer, nextLayer, times *uint8, primeRotation *bool) {
	for i := *layer; i <= *nextLayer; i++ {
		C.MoveZLayer(&i, times, primeRotation)
	}
}
