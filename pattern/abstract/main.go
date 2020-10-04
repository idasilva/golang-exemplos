package main

import (
	"fmt"
	"github.dxc.com/projects/golang-exemplos/pattern/abstract/factory"
)

func main()  {

	employee := factory.CreateANewEmployee(factory.Senior)
	info := fmt.Sprintf("Name:%s, Salary:%v",employee.Name(),employee.Salary())
	fmt.Println(info)

}
