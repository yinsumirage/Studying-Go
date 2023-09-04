package move_goyacc

import (
	"fmt"
	"time"
)

func Zmove(z int) error {
	fmt.Println("moving z to ", z)
	time.Sleep(time.Second)
	return nil
}

func Xmove(x int) error {
	fmt.Println("moving x to ", x)
	time.Sleep(time.Second)
	return nil
}

func Ymove(y int) error {
	fmt.Println("moving y to ", y)
	time.Sleep(time.Second)
	return nil
}

func Sleep(sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
}
