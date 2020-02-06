package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"

	_ "github.com/joho/godotenv/autoload"
)

func init() {
	// TODO: use gohowdy
	cmd := exec.Command("figlet")
	var result bytes.Buffer
	cmd.Stdin = strings.NewReader("go_mail")
	cmd.Stdout = &result
	_ = cmd.Run()
	fmt.Println(result.String())
	log.Println("gomail Ready!")
}

func main() {
	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		var msg Payload
		if r.Method == "GET" {
			msg.Sender = r.URL.Query().Get("sender")
			msg.Title = r.URL.Query().Get("title")
			msg.Message = r.URL.Query().Get("message")
		}
		if r.Method == "POST" {
			decoder := json.NewDecoder(r.Body)
			decoder.Decode(&msg)
		}
		msg.UserIP = readUserIP(r)
		msg.UserAgent = readUserAgent(r)

		err := SendMail(msg)

		w.Header().Add("Content-type", "application/json; charset=utf-8")
		if err != nil {
			fmt.Fprintf(w, "{\"result\": \"%+v\"}", err)
		} else {
			fmt.Fprintf(w, "{\"result\": \"ok\"}")
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "gomail\n%s %s", r.Method, r.URL.Path[1:])
	})
	http.ListenAndServe("0.0.0.0:25587", nil)
}
