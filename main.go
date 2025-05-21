package main

import "context"

func main() {
	client, collection, err := ConnectDB()
	if err != nil {
		panic(err)
	}
	defer client.Disconnect((context.TODO()))

	println("collection name : ", collection.Name())
}