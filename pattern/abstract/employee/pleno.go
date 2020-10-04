package employee

type Pleno struct {
}

func (p *Pleno) Name() string {
	return "pleno"

}

func (p *Pleno) Salary() float32 {
	return 2.500
}

func NewPleno() Employee{
	return &Pleno{}
}
