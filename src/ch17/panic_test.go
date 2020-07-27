package panic_test

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanic(t *testing.T) {
	fmt.Println("haha")
	//os.Exit(10)
	defer func() {
		fmt.Println("finally")
		if err := recover(); err != nil {
			fmt.Println("panic info are ", err)
		}
	}()
	panic(errors.New("panic 了"))
}
