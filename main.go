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
	sshUser := os.Args[2]
	sshHost := os.Args[1]
	sshPass := os.Args[3]

	// routing, request, response

	http.HandleFunc("/ps", func(w http.ResponseWriter, r *http.Request) {

		var ps = ssh.ServerProc(sshHost, sshUser, sshPass, "ps")

		fmt.Fprintf(w, "%s\n", "Ouput for ps command")
		fmt.Fprintf(w, "%s", ps)
		fmt.Fprintf(w, "%s", "Thank you")
	})

	http.HandleFunc("/uptime", func(w http.ResponseWriter, r *http.Request) {
		var uptime = ssh.ServerProc(sshHost, sshUser, sshPass, "uptime")
		fmt.Fprintf(w, "%s\n", "Ouput for uptime command")
		fmt.Fprintf(w, "%s", uptime)
		fmt.Fprintf(w, "%s", "Thank you")
	})

	http.HandleFunc("/df", func(w http.ResponseWriter, r *http.Request) {

		var dfOutput = ssh.ServerProc(sshHost, sshUser, sshPass, "df -h")
		fmt.Fprintf(w, "%s\n", "Ouput for df command")

		fmt.Fprintf(w, "%s", dfOutput)
		fmt.Fprintf(w, "%s", "Thank you")
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		var output = ssh.ServerProc(sshHost, sshUser, sshPass, "ps")
		var html = `<html>
				<style>
				h1 {
					text-align: center;
					font-family: 'Franklin Gothic Medium', 'Arial Narrow', Arial, sans-serif;
					color: #900C3F;
					font-size: 70 px;
				}
				
				h2 {
					text-align: center;
					font-family: 'Franklin Gothic Medium', 'Arial Narrow', Arial, sans-serif;
					color: #FFC300;
					font-size: 50 px;
				}

				p {
					text-align: center;
					color: #C70039;
				}

				ul {
					text-align: center;
					list-style-type: circle;
					color: #FF5733;
				}
				
				</style>
				<body>
					<div>
						<h1>Hello There!</h1>
						<h2>Let's check on your device!</h2>
						<ul>
					
							<li><a href="/ps">Check the running Processes with PS</a></li>
							<li><a href="/df">Check available disk sapace with DF</a></li>
							<li><a href="/uptime">Run UPTIME</a></li>
						</ul>


					</div>
					<div<%s</div>
				</body>
					</html>
			`

		fmt.Fprintf(w, html, output)

	})
	fmt.Println("listening")
	http.ListenAndServe(":8080", nil)
}

// "My name is %s and I am %d years old", name, age
