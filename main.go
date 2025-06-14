package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
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
		if len(os.Args) == 2 {
			err := ListTodos(collection)
			if err != nil {
				log.Fatalf("Falied to list todos: %v", err)
			}
		} else if len(os.Args) == 3 {
			status := os.Args[2]
			err := ListTodosOfStatus(collection, status)
			if err != nil {
				log.Fatalf("Error: %v", err)
			}
		}

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: update \"id\" \"description\" ")
			return
		}
		taskId, err := strconv.Atoi(os.Args[2])
		des := os.Args[3]
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		
		err = UpdateTodo(collection, taskId, des)
		if err != nil {
			log.Fatalf("Error %v", err)
		}

		fmt.Println("Todo update successfully")

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: delete \"id\"")
			return
		}
		taskId, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("Failed to delete todo: %v", err)
		}
		err = DeleteTodo(collection, taskId)
		if err != nil {
			log.Fatalf("Error %v", err)
		}
		fmt.Println("Todo deleted successfully")

	case "mark":
		if len(os.Args) < 4 {
			fmt.Println("Usage: mark id ")
		}
		taskId, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatalf("Failed to mark the todo: %v", err)
		}
		status := os.Args[3]
		err = UpdateStatus(collection, taskId, status)
		if err != nil {
			log.Fatal("failed to delete todo")
		}
		fmt.Println("Update status successfull")

	default:
		fmt.Println("Unknown command:", command)
		fmt.Println("Available commands: add")
	}
}