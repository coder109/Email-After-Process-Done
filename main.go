package main

import (
	"fmt"
	"os"
	"os/exec"
	"profiler/pkg"
)

func main() {
	// Init
	var info pkg.Information
	pkg.GetTime(&info, "start")
	email_info := pkg.InitPrivateEmailInformation()

	// Prepare the Log
	raw_log_file := pkg.GetLogFile(info)
	log_file, err := os.OpenFile(raw_log_file, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer log_file.Close()

	// Get the Command Name
	command := pkg.BuildCommand(os.Args, info)
	pkg.GetCommandName(&info)

	// Start the Process
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdout = log_file
	cmd.Stderr = log_file
	err = cmd.Start()
	if err != nil {
		panic(err)
	}

	// Wait till end
	err = cmd.Wait()

	pkg.GetTime(&info, "end")
	content := pkg.InformationToString(info)
	fmt.Println(content)
	pkg.SendEmail(content, email_info.Sender, email_info.Recver, email_info.Port, email_info.Smtp_server, email_info.Password, raw_log_file, err != nil)
}
