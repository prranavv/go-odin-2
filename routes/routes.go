package routes

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type Messages struct {
	Text  string
	User  string
	Added time.Time
}

func Mainpage(w http.ResponseWriter, r *http.Request) {
	messages := map[string][]Messages{
		"Messages": {
			{
				Text:  "Hello There",
				User:  "VK",
				Added: time.Now(),
			},
			{
				Text:  "Nice to See You",
				User:  "Mummy",
				Added: time.Now(),
			},
			{
				Text:  "My god",
				User:  "Kukku",
				Added: time.Now(),
			},
			{
				Text:  "Yes",
				User:  "Papa",
				Added: time.Now(),
			},
		},
	}
	tmpl := template.Must(template.ParseFiles("./Templates/index.html"))
	tmpl.Execute(w, messages)
}

func NewMessage(w http.ResponseWriter, r *http.Request) {
	text := r.PostFormValue("message")
	user := r.PostFormValue("username")
	var NMessage Messages
	NMessage.Text = text
	NMessage.User = user
	NMessage.Added = time.Now()
	htmlstr := fmt.Sprintf("<li class='messages'>%s <p></p>%s-%s</li>", user, text, NMessage.Added)
	tmpl, _ := template.New("p").Parse(htmlstr)
	tmpl.Execute(w, nil)
}
