package main

import (
	"27-simple-web-app-extended/game"

	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/play", playRound)

	const port = ":8080"

	log.Println("Starting web server on port", port)

	err := http.ListenAndServe(port, nil)
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	check(err)

	err = t.Execute(w, nil)
	check(err)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	_ = r

	renderTemplate(w, "index.html")
}

func playRound(w http.ResponseWriter, r *http.Request) {
	// Get the value of "c" which is represented in the javascript
	// choose() function at line 40: 'fetch("/play?c=" + x)'.
	playerChoice, err := strconv.Atoi(r.URL.Query().Get("c"))
	check(err)

	result := game.PlayRound(playerChoice)

	out, err := json.MarshalIndent(result, "", "    ")
	check(err)

	w.Header().Set("Content-Type", "application(json)")

	_, err = w.Write(out)
	check(err)
}
