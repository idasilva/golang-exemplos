package factory

import "github.dxc.com/projects/golang-exemplos/pattern/abstract/employee"

func CreateANewEmployee(employe string) employee.Employee {

	switch employe {
	case Junior:
		return employee.NewJunior()
	case Pleno:
		return employee.NewPleno()
	default:
		return employee.NewSenior()

	}
}
