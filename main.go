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
	email_info := pkg.InitPrivateEmailInformation()

	// Get the Command Name
	command := pkg.BuildCommand(os.Args)
	pkg.GetCommandName(&info)

	// Find the Path
	path := pkg.LookPath(command)

	// Start the Process
	pro, err := os.StartProcess(path, command, &os.ProcAttr{})
	if err != nil {
		panic(err)
	}

	// Wait
	pro.Wait()

	pkg.GetTime(&info, "end")
	content := pkg.InformationToString(info)
	fmt.Println(content)
	pkg.SendEmail(content, email_info.Sender, email_info.Recver, email_info.Port, email_info.Smtp_server, email_info.Password)
}
