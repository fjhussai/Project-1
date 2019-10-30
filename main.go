package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

// this struct is supposed to hold the information that's gunna be displayed in our html file
type Welcome struct {
	Name string
	Time string
}

func main() {
	welcome := Welcome{"Whoever", time.Now().Format(time.Stamp)}

	//let's tell go where to find our html file
	templates := template.Must(template.ParseFiles("templates/welcome-template.html"))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if name := r.FormValue("name"); name != "" {
			welcome.Name = name
		}
		if err := templates.ExecuteTemplate(w, "welcome-template.html", welcome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	})

	fmt.Println("Listening")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
