package main

import (
	"fmt"
	"os"
	"profiler/pkg"
)

func main() {
	// Init
	var info pkg.Information
	pkg.GetTime(&info, "start")

	// Get the Command Name
	command := pkg.BuildCommand(os.Args)
	pkg.GetCommandName(&info)

	// Find the Path
	path := pkg.LookPath(command)

	// Start the Process
	_, err := os.StartProcess(path, command, &os.ProcAttr{})
	if err != nil {
		panic(err)
	} else {
		pkg.GetTime(&info, "end")
		content := pkg.InformationToString(info)
		fmt.Println(content)
		pkg.SendEmail(content)
	}
}
