package Color

import (
	"fmt"
	"testing"
)

func TestColors(t *testing.T) {
	println(fmt.Sprintf(Black, "hello Black"))
	println(fmt.Sprintf(Blue, "hello Blue"))
	println(fmt.Sprintf(Green, "hello Green"))
	println(fmt.Sprintf(Red, "hello Red"))
	println(fmt.Sprintf(Orange, "hello Orange"))
	println(fmt.Sprintf(White, "hello White"))
	println(fmt.Sprintf(Yellow, "hello Yellow"))
}
