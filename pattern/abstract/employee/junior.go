package employee

type Junior struct {
}

func (p *Junior) Name() string {
	return "junior"

}

func (p *Junior) Salary() float32 {
	return 2.300
}

func NewJunior() Employee{
	return &Junior{}
}
