package build

type director struct {
	builder []iBuilder
}
func(d *director) setBuilder(b ...iBuilder){
	d.builder = b
}

func(d *director) BuildHouse() []house {
	var houses = make([]house, 0)
	for _, i := range(d.builder){
		i.setDoorType()
		i.setWindowType()
		i.setNumFloor()
		houses = append(houses,i.getHouse())

	}
	return houses
}
func NewDiretor() *director {
	return &director{
	}
}
