package main

import (
	"fmt"

	"github.com/pearish/modular_cloud/common"
)

func main() {
	c := common.NewCloud("first", "second")
	fmt.Println(c.A())
	fmt.Println(c.B())
	fmt.Println(c.C())
	fmt.Println(c.D())
	fmt.Println(c.E())
}
