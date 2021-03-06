package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Model
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
	//Headers
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(tasks)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	//Get ID
	params := mux.Vars(r)
	taskID, err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Fprintf(w, "ID invalid!")
		return
	}

	//Get obj task from tasks array
	for _, t := range tasks {
		if t.ID == taskID {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(t)
		}
	}

}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask task

	//Add ioutil module
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a task valid")
	}

	json.Unmarshal(reqBody, &newTask)

	//Generate ID thought index
	newTask.ID = len(tasks) + 1
	tasks = append(tasks, newTask)

	//Headers
	w.Header().Set("Content-Type", "application/json")
	//Sending Status
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newTask)
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskID, err := strconv.Atoi(params["id"])

	var updTask task

	if err != nil {
		fmt.Fprintf(w, "ID invalid!")
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Please enter valid data!")
	}

	json.Unmarshal(reqBody, &updTask)

	for i, t := range tasks {
		if t.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			updTask.ID = taskID
			tasks = append(tasks, updTask)

			fmt.Fprintf(w, "The task with ID %v has been updated successfully!", taskID)
		}
	}

}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskID, err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Fprintf(w, "ID invalid!")
		return
	}

	for i, t := range tasks {
		if t.ID == taskID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Fprintf(w, "Task with ID %v has been remove successfully!", taskID)
		}
	}
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome a my REST API")
}

func main() {
	//fmt.Println("Hello World!")
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/api/tasks", getTasks).Methods("GET")
	router.HandleFunc("/api/tasks/{id}", getTask).Methods("GET")
	router.HandleFunc("/api/tasks", createTask).Methods("POST")
	router.HandleFunc("/api/tasks/{id}", updateTask).Methods("PUT")
	router.HandleFunc("/api/tasks/{id}", deleteTask).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", router))
}
