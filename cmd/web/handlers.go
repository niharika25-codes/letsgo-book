package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"snippetbox.niharika.net/internal/models"
)

// handler function
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("server", "go")

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	for _, snippet := range snippets {
		fmt.Fprintf(w, "%+v\n", snippet)
	}

	/*
	files := []string {
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
	}
	err = ts.ExecuteTemplate(w, "base", nil)
	//w.Write([]byte("Hello from Snippetbox"))
	if err != nil {
		app.serverError(w, r, err)
	}
		*/
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
	}
	/*files := []string{
        "./ui/html/base.tmpl",
        "./ui/html/partials/nav.tmpl",
        "./ui/html/pages/view.tmpl",
    }

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
	}
	err = ts.ExecuteTemplate(w, "base", snippet)*/
	//fmt.Fprintf(w, "Display a specific snippet ID %d ...", id)
	fmt.Fprintf(w, "%+v", snippet)

	
}

func (app *application)  snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	title := "O snail"
    content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
    expires := 7
	id, err := app.snippets.Insert(title, content, expires)
    if err != nil {
        app.serverError(w, r, err)
        return
    }
	 http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}