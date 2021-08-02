package Cube

import (
	"fmt"
	"math/rand"
	"regexp"
	"strconv"
	"time"
)

var moves = []string{"U", "D", "R", "L", "F", "B"}

// moves     = "U", "D", "R", "L", "F", "B"
// rotations = "2 rotations", "1 rotation"

// -- if rotations == 1

//	 rotationType = "Prime", "noPrime"

// -- if width > 3

//	 layer = (C.Width / 2) - 1

func (C *Cube) CreateScramble(length uint8) {
	beforeIndex, xIndex, yIndex, zIndex := -1, uint(0), uint(0), uint(0)

	rand.Seed(time.Now().Unix())
	C.scramble = []string{}

	for i := uint8(0); i < length; i++ {
		for {
			if index := rand.Intn(6); index != beforeIndex {
				movement := moves[index]

				if index < 2 {
					yIndex++
					xIndex, zIndex = 0, 0
				} else if index < 4 {
					xIndex++
					yIndex, zIndex = 0, 0
				} else if index < 6 {
					zIndex++
					xIndex, yIndex = 0, 0
				}

				if xIndex == 2 || yIndex == 2 || zIndex == 2 {
					continue
				}

				if C.width > 3 && (C.width != 4 || i > 18) {
					layerMax := int(C.width / 2)

					layerNum := rand.Intn(layerMax + 1)

					if layerNum > 0 {
						movement += "w"
					}
					if layerNum > 2 {
						movement = fmt.Sprint(layerNum) + movement
					}
				}

				if rotaionNum := rand.Intn(2); rotaionNum == 1 {
					movement += "2"
				} else if primeRotarion := rand.Intn(2); primeRotarion == 1 {
					movement += "'"
				}

				C.scramble = append(C.scramble, movement)

				beforeIndex = index
				break
			}
		}
	}
	C.MoveCube(&C.scramble)
}

func (C *Cube) MoveCube(scramble *[]string) {

	xLayer, _ := regexp.Compile(`(^\d?[RL]{1}w[2\']?$)|(^[RML]{1}[2\']?$)`)
	yLayer, _ := regexp.Compile(`(^\d?[UD]{1}w[2\']?$)|(^[UED]{1}[2\']?$)`)
	zLayer, _ := regexp.Compile(`(^\d?[FB]{1}w[2\']?$)|(^[FSB]{1}[2\']?$)`)

	rotation, _ := regexp.Compile(`[MLEDSB]`)
	primeRotation, _ := regexp.Compile(`\'`)

	rotationNum, _ := regexp.Compile(`2$`)

	middleLayer, _ := regexp.Compile(`[MES]`)
	// firstLayer, _ := regexp.Compile(`[ULB]`)
	lastLayer, _ := regexp.Compile(`[DRF]`)

	wLayer, _ := regexp.Compile(`w`)
	getWLayer, _ := regexp.Compile(`^\d?`)

	layerNum, nextLayerNum, times, prime := uint8(0), uint8(0), uint8(1), false

	for _, movement := range *scramble {

		// fmt.Print(i, ") ", movement)

		if rotation.MatchString(movement) {
			prime = true
		}

		if primeRotation.MatchString(movement) {
			prime = !prime
		}

		if rotationNum.MatchString(movement) {
			times = 2
		}

		if middleLayer.MatchString(movement) {
			if C.width%2 == 0 {
				continue
			}
			layerNum = C.width / 2
			nextLayerNum = layerNum
		}

		if wLayer.MatchString(movement) {
			WLayer := getWLayer.FindString(movement)
			num, _ := strconv.Atoi(WLayer)

			if num < 3 {
				nextLayerNum = 1
			} else if num > int(C.width)-1 {
				nextLayerNum = C.width - 1
			} else if num > 2 {
				nextLayerNum = uint8(num) - 1
			}
		}

		if lastLayer.MatchString(movement) {
			layerNum = (C.width - 1) - layerNum
			nextLayerNum = (C.width - 1) - nextLayerNum

			// fmt.Println(layerNum, nextLayerNum)
		}

		if xLayer.MatchString(movement) {
			// fmt.Print(" : X")
			if layerNum > nextLayerNum {
				// fmt.Println(" : 1Case")
				C.MoveXwLayer(&nextLayerNum, &layerNum, &times, &prime)
			} else if nextLayerNum > layerNum {
				// fmt.Println(" : 2Case")
				C.MoveXwLayer(&layerNum, &nextLayerNum, &times, &prime)
			} else {
				// fmt.Println(" : 3Case")
				C.MoveXLayer(&layerNum, &times, &prime)
			}
		} else if yLayer.MatchString(movement) {
			// fmt.Print(" : Y")
			if layerNum > nextLayerNum {
				// fmt.Println(" : 1Case")
				C.MoveYwLayer(&nextLayerNum, &layerNum, &times, &prime)
			} else if nextLayerNum > layerNum {
				// fmt.Println(" : 2Case")
				C.MoveYwLayer(&layerNum, &nextLayerNum, &times, &prime)
			} else {
				// fmt.Println(" : 3Case")
				C.MoveYLayer(&layerNum, &times, &prime)
			}
		} else if zLayer.MatchString(movement) {
			// fmt.Print(" : Z")
			if layerNum > nextLayerNum {
				// fmt.Println(" : 1Case")
				C.MoveZwLayer(&nextLayerNum, &layerNum, &times, &prime)
			} else if nextLayerNum > layerNum {
				// fmt.Println(" : 2Case")
				C.MoveZwLayer(&layerNum, &nextLayerNum, &times, &prime)
			} else {
				// fmt.Println(" : 3Case")
				C.MoveZLayer(&layerNum, &times, &prime)
			}
		} else {
			// fmt.Println("invalid Scrable")
			return
		}
		layerNum, nextLayerNum, times, prime = uint8(0), uint8(0), uint8(1), false
	}
}
