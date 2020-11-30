package command

import "github.dxc.com/projects/golang-exemplos/pattern/commands/receiver"

type OnCommand struct {
	Device receiver.Device
}

func (c *OnCommand) execute()  {
	c.Device.On()
	
}
