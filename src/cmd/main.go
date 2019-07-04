package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/revenue-hack/golang-google-tasks/src"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/tasks/v1"
)

func main() {

	srv := getService()

	showTODOList(srv)

	showTaskList(srv)
}

func showTODOList(srv *tasks.Service) {
	list := src.NewTODOOperation(srv.Tasklists).List()

	fmt.Println("TODO Lists:")
	if len(list) > 0 {
		for _, i := range list {
			fmt.Printf("%s (%s)\n", i.Title, i.Id)
		}
	} else {
		fmt.Print("No todo lists found.")
	}
}

func showTaskList(srv *tasks.Service) {
	todo := src.NewTODOOperation(srv.Tasklists).First()
	list := src.NewTaskOperation(srv.Tasks).ListByTODOID(todo.Id)

	fmt.Println("Task Lists:")
	if len(list) > 0 {
		for _, i := range list {
			fmt.Printf("%s\t%s\t%s\n", i.Title, i.Due, i.SelfLink)
		}
	} else {
		fmt.Print("No task lists found.")
	}
}

func getService() *tasks.Service {
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, tasks.TasksReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := src.GetClient(config)

	srv, err := tasks.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve tasks Client %v", err)
	}

	return srv
}
