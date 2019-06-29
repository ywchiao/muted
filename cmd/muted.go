package main

import (
	"encoding/json"
	//	"errors"
	//	"html/template"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

type Room struct {
	Tag         string
	Name        string
	Brief       string
	Description string
}

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

func (p *Room) save() error {
	filename := p.Tag + ".json"

	text, _ := json.Marshal(p)

	return ioutil.WriteFile(filename, []byte(text), 0600)
}

func loadRoom(title string) ([]byte, error) {
	filename := title + ".json"

	body, err := ioutil.ReadFile(filename)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadRoom(title)

	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)

		return
	}

	fmt.Fprintf(w, "%s", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	//	p, err := loadRoom(title)

	//	if err != nil {
	//		p = &Room{Tag: title}
	//	}
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	//	body := r.FormValue("body")

	//	json.NewDecoder(r.Body).Decode(p)

	//	p := &Room{Title: title, Body: []byte(body)}
	//	err := p.save()

	//	if err != nil {
	//		http.Error(w, err.Error(), http.StatusInternalServerError)

	//		return
	//	}

	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)

		if m == nil {
			http.NotFound(w, r)

			return
		}

		fn(w, r, m[2])
	}
}

func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	//	http.HandleFunc("/edit/", makeHandler(editHandler))
	//	http.HandleFunc("/save/", makeHandler(saveHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
