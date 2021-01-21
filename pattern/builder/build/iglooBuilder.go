package build

type iglooBuilder struct {
	windowType string
	doorType string
	floor int
}

func(n *iglooBuilder) setWindowType(){
	n.windowType = "Wooden Window 2"

}

func(n *iglooBuilder) setDoorType(){
	n.doorType = "Wooden Door 2"

}
func (n *iglooBuilder) setNumFloor(){
	n.floor  = 4

}
func(n *iglooBuilder) getHouse() house {
	return house{
		n.windowType,
		n.doorType,
		n.floor,

	}

}

func NewIglooBUilder() *iglooBuilder{
	return &iglooBuilder{}
}