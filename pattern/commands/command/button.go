package command

type Button  struct {
	Command Command
}

func(b *Button) Press()  {
	b.Command.execute()
}