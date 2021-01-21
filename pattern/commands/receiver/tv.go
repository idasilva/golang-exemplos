package receiver

import "fmt"

type Tv struct {
	isRunning bool
}

func (t *Tv)On(){
	t.isRunning = true
	fmt.Println("Turning on")
}

func (t *Tv) Off(){
	t.isRunning = false
	fmt.Println("Turning of")
}
