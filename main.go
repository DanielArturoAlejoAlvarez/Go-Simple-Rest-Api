package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type task struct {
	ID      int    `json:'ID'`
	Name    string `json:'Name'`
	Content string `json:'Content'`
}

type allTasks []task

var tasks = allTasks{
	{
		ID:      1,
		Name:    "First Task",
		Content: "This is a content of our task",
	},
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(tasks)
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome a my REST API")
}

func main() {
	//fmt.Println("Hello World!")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/api/tasks", getTasks)

	log.Fatal(http.ListenAndServe(":9000", router))
}
