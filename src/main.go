package main

import (
	"fmt"
	//内部引入的路径也是全路径
	"github.com/gitdebugdemo/go/src/funcs"
)

func main() {
	say := funcs.Say()
	fmt.Printf(say)
}
