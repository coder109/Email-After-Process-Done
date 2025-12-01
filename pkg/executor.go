package pkg

import (
	"fmt"
	"os/exec"
)

func BuildCommand(args []string) []string {
	if len(args) == 1 {
		panic("No Command")
	}

	command := append([]string{}, args[1:]...)
	return command
}

func LookPath(command_name []string) string {
	path, err := exec.LookPath(command_name[0])
	if err != nil {
		panic(err)
	}
	return path
}

func OutputCommandName(command_name []string) {
	for _, arg := range command_name {
		fmt.Println(arg)
	}
}
