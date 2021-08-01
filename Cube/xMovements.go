package Cube

import "github.com/ImS0L0/GoRubikCube/Face"

func (C *Cube) moveXLayersTo() {
	frontList := C.faceFront.GetColumn(&Layer)

	counterLayer := uint8((C.width - Layer) - 1)

	if Prime {
		C.faceFront.RepalceColumn(&Layer, C.faceUp.GetColumn(&Layer))
		C.faceUp.RepalceColumn(&Layer, Face.InvertList(C.faceBack.GetColumn(&counterLayer)))
		C.faceBack.RepalceColumn(&counterLayer, Face.InvertList(C.faceDown.GetColumn(&Layer)))
		C.faceDown.RepalceColumn(&Layer, frontList)
	} else {
		C.faceFront.RepalceColumn(&Layer, C.faceDown.GetColumn(&Layer))
		C.faceDown.RepalceColumn(&Layer, Face.InvertList(C.faceBack.GetColumn(&counterLayer)))
		C.faceBack.RepalceColumn(&counterLayer, Face.InvertList(C.faceUp.GetColumn(&Layer)))
		C.faceUp.RepalceColumn(&Layer, frontList)
	}
}

func (C *Cube) rotateXFace() {
	if Layer == 0 {
		C.faceLeft.RotateTo(&noPrime)
	} else if Layer == C.width-1 {
		C.faceRight.RotateTo(&Prime)
	}
}

func (C *Cube) MoveXLayer(layer, times *uint8, primeRotation *bool) {
	Layer, Prime, noPrime = *layer, *primeRotation, !*primeRotation

	for i := uint8(0); i < *times; i++ {
		C.rotateXFace()
		C.moveXLayersTo()
	}
}

func (C *Cube) MoveXwLayer(layer, nextLayer, times *uint8, primeRotation *bool) {
	for i := *layer; i <= *nextLayer; i++ {
		C.MoveXLayer(&i, times, primeRotation)
	}
}
