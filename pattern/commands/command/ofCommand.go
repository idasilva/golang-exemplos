package command

import "github.dxc.com/projects/golang-exemplos/pattern/commands/receiver"

type OffCommand struct {
	Device receiver.Device
}

func (o *OffCommand) execute()  {
	o.Device.Off()
	
}