package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Usage: todo <command> [arguments]")
		return
	}

	command := os.Args[1]

	client, collection, err := ConnectDB()
	
	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}
	defer client.Disconnect(context.TODO())
	

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo add \"task\" [\"Description\"]")
			return
		}
		
		task := os.Args[2]
		description := ""
		if len(os.Args) >= 4 {
			description = strings.Join(os.Args[3:], " ")
		}

		err := AddTodo(collection, task, description)
		if err != nil {
			log.Fatalf("Falied to add todo : %v", err)
		} 

		fmt.Println("Todo added successfully!")
	
		case "list":
			err := ListTodos(collection)
			if err != nil {
				log.Fatalf("Falied to list todos: %v", err)
			}

	default:
		fmt.Println("Unknown command:", command)
		fmt.Println("Available commands: add")
	}
	


}