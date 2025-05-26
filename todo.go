package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
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

	styles := []table.Style {
		table.StyleDefault,
		table.StyleLight,
		table.StyleColoredDark,
		table.StyleColoredBlueWhiteOnBlack,
	}

	t := table.NewWriter()
	t.SetColumnConfigs([]table.ColumnConfig{
		{Number: 1, Align: text.AlignCenter}, // Index
    {Number: 2, Align: text.AlignCenter}, // Task
    {Number: 3, Align: text.AlignCenter}, // Status
    {Number: 4, Align: text.AlignCenter}, // Description
	})
	
	t.SetCaption("Todo List")

	t.AppendHeader(table.Row{"#", "Task", "status", "Description", "Created At"})
	t.SetTitle("title")

	for _, style := range styles {
		idx := 1
		for cursor.Next(ctx) {
			var todo Todo 
			if err := cursor.Decode(&todo); err != nil {
				return err 
			}
			status := "❌"
			if todo.Completed {
				status = "✅"
			}
			fmt.Println(todo.CreatedAt)
			t.AppendRow(table.Row{idx, todo.Task, status, todo.Description, todo.CreatedAt.Local().Format("Jan 02, 2006,\n 03:04:05 PM")})
			idx++
		}
		
		t.SetAutoIndex(true)
		t.SetStyle(style)
		t.Style().Options.SeparateRows = true 
		t.Style().Format.Header = text.FormatTitle

		fmt.Println(t.Render())
		fmt.Println()

	}

	

	return nil
}