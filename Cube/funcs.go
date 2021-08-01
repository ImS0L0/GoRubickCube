package Cube

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var (
	Layer   uint8
	Prime   bool
	noPrime bool
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

// func InvertList(oldList []uint8) []uint8 {
// 	listLen := len(oldList)
// 	newList := make([]uint8, listLen)

// 	for i, j := 0, listLen-1; i < listLen; i++ {
// 		newList[i] = oldList[j]
// 		j--
// 	}

// 	return newList
// }
