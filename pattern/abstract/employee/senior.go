package employee

type Senior struct {
}

func (p *Senior) Name() string {
	return "senior"

}

func (p *Senior) Salary() float32 {
	return 8.450
}

func NewSenior() Employee{
	return &Senior{}
}
