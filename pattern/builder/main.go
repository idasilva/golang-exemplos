package main

import (
	"fmt"
	"github.dxc.com/projects/golang-exemplos/pattern/builder/build"
)

func main(){
	director := build.NewDiretor()
	build.HandlerBuilder(director)


	House := director.BuildHouse()

	for _, i := range(House) {
		fmt.Printf("Normal House Door Type: %s\n", i.DoorType)
		fmt.Printf("Normal House Window Type: %s\n", i.WindowType)
		fmt.Printf("Normal House Num Floor: %d\n", i.Floor)

	}
}