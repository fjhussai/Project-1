package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/fjhussai/project-1/ssh"
)

/* this struct is supposed to hold the information that's gunna be displayed in our html file
type Welcome struct {
	Name string
	Time string
}
*/

func main() {
	var sshUser = os.Args[2]
	var sshHost = os.Args[1]
	var sshPass = os.Args[3]

	output := ssh.ServerProc(sshHost, sshUser, sshPass, "ps")

	http.HandleFunc("/", HelloServer(http.ResponseWriter, *http.Request, output))
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request, output string) {

	fmt.Fprint(w, "%s", output)
}

// "My name is %s and I am %d years old", name, age
