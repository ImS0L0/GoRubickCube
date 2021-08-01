package Cube

func (C *Cube) moveYLayersTo() {
	frontList := C.faceFront[Layer]

	if Prime {
		C.faceFront[Layer] = C.faceLeft[Layer]
		C.faceLeft[Layer] = C.faceBack[Layer]
		C.faceBack[Layer] = C.faceRight[Layer]
		C.faceRight[Layer] = frontList
	} else {
		C.faceFront[Layer] = C.faceRight[Layer]
		C.faceRight[Layer] = C.faceBack[Layer]
		C.faceBack[Layer] = C.faceLeft[Layer]
		C.faceLeft[Layer] = frontList
	}
}

func (C *Cube) rotateYFace() {
	if Layer == 0 {
		C.faceUp.RotateTo(&Prime)
	} else if Layer == C.width-1 {
		C.faceDown.RotateTo(&noPrime)
	}
}

func (C *Cube) MoveYLayer(layer, times *uint8, primeRotation *bool) {
	Layer, Prime, noPrime = *layer, *primeRotation, !*primeRotation

	for i := uint8(0); i < *times; i++ {
		C.rotateYFace()
		C.moveYLayersTo()
	}
}

func (C *Cube) MoveYwLayer(layer, nextLayer, times *uint8, primeRotation *bool) {
	for i := *layer; i <= *nextLayer; i++ {
		C.MoveYLayer(&i, times, primeRotation)
	}
}
