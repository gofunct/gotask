package main

/**
 * This is the main file for the Task application
 * License: MIT
 **/
import (
	"flag"
	"log"
	"net/http"
	"strings"

	"github.com/gofunct/coleman/config"
	"github.com/gofunct/coleman/sessions"
	"github.com/gofunct/gotask/views"
)

func main() {
	values, err := config.ReadConfig("config.json")
	var port *string

	if err != nil {
		port = flag.String("port", "", "IP address")
		flag.Parse()

		//User is expected to give :8080 like input, if they give 8080
		//we'll append the required ':'
		if !strings.HasPrefix(*port, ":") {
			*port = ":" + *port
			log.Println("port is " + *port)
		}

		values.ServerPort = *port
	}

	views.PopulateTemplates()

	//Login logout
	http.HandleFunc("/login/", views.LoginFunc)
	http.HandleFunc("/logout/", sessions.RequiresLogin(sessions.LogoutFunc))
	http.HandleFunc("/signup/", sessions.SignUpFunc)

	http.HandleFunc("/add-category/", sessions.RequiresLogin(views.AddCategoryFunc))
	http.HandleFunc("/add-comment/", sessions.RequiresLogin(views.AddCommentFunc))
	http.HandleFunc("/add/", sessions.RequiresLogin(views.AddTaskFunc))

	//these handlers are used to delete
	http.HandleFunc("/del-comment/", sessions.RequiresLogin(views.DeleteCommentFunc))
	http.HandleFunc("/del-category/", sessions.RequiresLogin(views.DeleteCategoryFunc))
	http.HandleFunc("/delete/", sessions.RequiresLogin(views.DeleteTaskFunc))

	//these handlers update
	http.HandleFunc("/upd-category/", sessions.RequiresLogin(views.UpdateCategoryFunc))
	http.HandleFunc("/update/", sessions.RequiresLogin(views.UpdateTaskFunc))

	//these handlers are used for restoring tasks
	http.HandleFunc("/incomplete/", sessions.RequiresLogin(views.RestoreFromCompleteFunc))
	http.HandleFunc("/restore/", sessions.RequiresLogin(views.RestoreTaskFunc))

	//these handlers fetch set of tasks
	http.HandleFunc("/", sessions.RequiresLogin(views.ShowAllTasksFunc))
	http.HandleFunc("/category/", sessions.RequiresLogin(views.ShowCategoryFunc))
	http.HandleFunc("/deleted/", sessions.RequiresLogin(views.ShowTrashTaskFunc))
	http.HandleFunc("/completed/", sessions.RequiresLogin(views.ShowCompleteTasksFunc))

	//these handlers perform action like delete, mark as complete etc
	http.HandleFunc("/complete/", sessions.RequiresLogin(views.CompleteTaskFunc))
	http.HandleFunc("/files/", sessions.RequiresLogin(views.UploadedFileHandler))
	http.HandleFunc("/trash/", sessions.RequiresLogin(views.TrashTaskFunc))
	http.HandleFunc("/edit/", sessions.RequiresLogin(views.EditTaskFunc))
	http.HandleFunc("/search/", sessions.RequiresLogin(views.SearchTaskFunc))

	http.Handle("/static/", http.FileServer(http.Dir("public")))

	http.HandleFunc("/api/get-task/", views.GetTasksFuncAPI)
	http.HandleFunc("/api/get-deleted-task/", views.GetDeletedTaskFuncAPI)
	http.HandleFunc("/api/add-task/", views.AddTaskFuncAPI)
	http.HandleFunc("/api/update-task/", views.UpdateTaskFuncAPI)
	http.HandleFunc("/api/delete-task/", views.DeleteTaskFuncAPI)

	http.HandleFunc("/api/get-token/", views.GetTokenHandler)
	http.HandleFunc("/api/get-category/", views.GetCategoryFuncAPI)
	http.HandleFunc("/api/add-category/", views.AddCategoryFuncAPI)
	http.HandleFunc("/api/update-category/", views.UpdateCategoryFuncAPI)
	http.HandleFunc("/api/delete-category/", views.DeleteCategoryFuncAPI)

	log.Println("running server on ", values.ServerPort)
	log.Fatal(http.ListenAndServe(values.ServerPort, nil))
}
