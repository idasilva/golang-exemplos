package main
import (
	"github.dxc.com/projects/golang-exemplos/pattern/commands/command"
	"github.dxc.com/projects/golang-exemplos/pattern/commands/receiver"
)

func main (){
	tv := &receiver.Tv{}
	onCommand := &command.OnCommand{
		Device: tv,
	}
	offCommand := &command.OffCommand{
		Device: tv,
	}
	onButton := &command.Button{
		Command: onCommand,
	}
	onButton.Press()
	offButton := &command.Button{
		Command: offCommand,
	}
	offButton.Press()
}
