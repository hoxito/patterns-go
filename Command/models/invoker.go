package models

type Button struct {
	Command Command
}

func (b *Button) Press() {
	b.Command.Execute()
}
