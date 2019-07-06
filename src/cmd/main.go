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

	updateTODO(srv, "hogehgo", "aaaaa")
}

func showTODOList(srv *tasks.Service) {
	list, err := src.NewTODOOperation(src.NewTODOOOpWrap(srv.Tasklists)).List()
	if err != nil {
		log.Fatal(err)
	}

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
	todo, err := src.NewTODOOperation(src.NewTODOOOpWrap(srv.Tasklists)).First()
	if err != nil {
		log.Fatal(err)
	}
	list, err := src.NewTaskOperation(src.NewTaskOpWrap(srv.Tasks)).ListByTODOID(todo.Id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Task Lists:")
	if len(list) > 0 {
		for _, i := range list {
			fmt.Printf("%s\t%s\t%s\n", i.Title, i.Due, i.SelfLink)
		}
	} else {
		fmt.Print("No task lists found.")
	}
}

func createTODO(srv *tasks.Service, title string) {
	todo, err := src.NewTODOOperation(src.NewTODOOOpWrap(srv.Tasklists)).Create(title)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("create TODO: %s\n", todo.Title)
}

func updateTODO(srv *tasks.Service, targetTitle, updatedTitle string) {
	todo, err := src.NewTODOOperation(src.NewTODOOOpWrap(srv.Tasklists)).UpdateTitleByTODOID(targetTitle, updatedTitle)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("update TODO: %s", todo.Title)
}

func deleteTODO(srv *tasks.Service, title string) {
	if err := src.NewTODOOperation(src.NewTODOOOpWrap(srv.Tasklists)).DeleteByTODOID(title); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("delete TODO: %s", title)
}

func getService() *tasks.Service {
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, tasks.TasksScope)
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
