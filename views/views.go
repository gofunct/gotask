package views

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gofunct/coleman/sessions"
	"github.com/gofunct/gotask/db"
)

var homeTemplate *template.Template
var deletedTemplate *template.Template
var completedTemplate *template.Template
var editTemplate *template.Template
var searchTemplate *template.Template
var templates *template.Template
var loginTemplate *template.Template

var message string //notification
var err error

func ShowAllTasksFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		username := sessions.GetCurrentUserName(r)
		context, err := db.GetTasks(username, "pending", "")
		log.Println(context)
		categories := db.GetCategories(username)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusInternalServerError)
		} else {
			if message != "" {
				context.Message = message
			}
			context.CSRFToken = "abcd"
			context.Categories = categories
			message = ""
			expiration := time.Now().Add(365 * 24 * time.Hour)
			cookie := http.Cookie{Name: "csrftoken", Value: "abcd", Expires: expiration}
			http.SetCookie(w, &cookie)
			homeTemplate.Execute(w, context)
		}
	}
}

// Handler for /trash
func ShowTrashTaskFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		username := sessions.GetCurrentUserName(r)
		categories := db.GetCategories(username)
		context, err := db.GetTasks(username, "deleted", "")
		context.Categories = categories
		if err != nil {
			http.Redirect(w, r, "/trash", http.StatusInternalServerError)
		}
		if message != "" {
			context.Message = message
			message = ""
		}
		err = deletedTemplate.Execute(w, context)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func ShowCompleteTasksFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		username := sessions.GetCurrentUserName(r)
		categories := db.GetCategories(username)
		context, err := db.GetTasks(username, "completed", "")
		context.Categories = categories
		if err != nil {
			http.Redirect(w, r, "/completed", http.StatusInternalServerError)
		}
		completedTemplate.Execute(w, context)
	}
}

func ShowCategoryFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" && sessions.IsLoggedIn(r) {
		category := r.URL.Path[len("category/"):]
		username := sessions.GetCurrentUserName(r)
		categories := db.GetCategories(username)
		context, err := db.GetTasks(username, "", category)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusInternalServerError)
		}
		if message != "" {
			context.Message = message
		}
		context.CSRFToken = "abcd"
		context.Categories = categories
		message = ""
		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{Name: "csrftoken", Value: "abcd", Expires: expiration}
		http.SetCookie(w, &cookie)
		homeTemplate.Execute(w, context)
	}
}

//LoginFunc implements the login functionality, will add a cookie to the cookie store for managing authentication
func LoginFunc(w http.ResponseWriter, r *http.Request) {
	session, err := sessions.Store.Get(r, "session")

	if err != nil {
		log.Println("error identifying session")
		loginTemplate.Execute(w, nil)
		return
	}

	switch r.Method {
	case "GET":
		loginTemplate.Execute(w, nil)
	case "POST":
		log.Print("Inside POST")
		r.ParseForm()
		username := r.Form.Get("username")
		password := r.Form.Get("password")

		if (username != "" && password != "") && db.ValidUser(username, password) {
			session.Values["loggedin"] = "true"
			session.Values["username"] = username
			session.Save(r, w)
			log.Print("user ", username, " is authenticated")
			http.Redirect(w, r, "/", 302)
			return
		}
		log.Print("Invalid user " + username)
		loginTemplate.Execute(w, nil)
	default:
		http.Redirect(w, r, "/login/", http.StatusUnauthorized)
	}
}
