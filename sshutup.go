package main

import (
	"os/exec"
)

func letstrythis() {
	exec.Command("ssh", "ubuntu@192.168.56.103").Run()
	exec.Command("ubuntu").Run()

}
