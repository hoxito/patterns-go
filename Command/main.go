package command

func main() {
	sword := &weapon{}

	SellCommand := &SellCommand{
		item: sword,
	}

	SellButton := &Button{
		command: SellCommand,
	}
	SellButton.press()

}
