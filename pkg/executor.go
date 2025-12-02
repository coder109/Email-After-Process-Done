package pkg

func BuildCommand(args []string, info Information) []string {
	if len(args) == 1 {
		panic("No Command")
	}

	command := append([]string{}, args[1:]...)

	return command
}
