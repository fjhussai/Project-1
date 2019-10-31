package ssh

import (
	"fmt"

	"github.com/sfreiberg/simplessh"
)

func ServerProc(sshHost string, sshUser string, sshPass string, command string) []byte {

	var client *simplessh.Client
	var err error

	fmt.Println(sshUser)
	fmt.Println(sshPass)
	fmt.Println(sshHost)

	if client, err = simplessh.ConnectWithPassword(sshHost, sshUser, sshPass); err != nil {
		fmt.Print(err)
	}

	if err != nil {
		panic(err)
	}
	defer client.Close()

	output, err := client.Exec("ps")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Uptime: %s\n", output)

	return output
}
