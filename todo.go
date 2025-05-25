package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Todo struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`        
	Task        string             `bson:"task"`                 
	Description string             `bson:"description"`          
	Completed   bool               `bson:"completed"`            
	CreatedAt   time.Time          `bson:"created_at"`           
	CompletedAt *time.Time         `bson:"completed_at,omitempty"`
}

func AddTodo(collection *mongo.Collection, task string, description string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()

	todo := Todo {
		Task: task,
		Description: description,
		Completed: false,
		CreatedAt: time.Now(),
	}
	_, err := collection.InsertOne(ctx, todo)
	return err 
}

func ListTodos(collection *mongo.Collection) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return err 	
	}
	defer cursor.Close(ctx)

	fmt.Println("Todo List: ")
	for cursor.Next(ctx) {
		var todo Todo 
		if err := cursor.Decode(&todo); err != nil {
			return err 
		}
		status := "❌"
		if todo.Completed {
			status = "✅"
		}
		fmt.Printf("%s %s - %s, %s\n", status, todo.Task, todo.Description, todo.CreatedAt)
	}
	return nil
}