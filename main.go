package main

import (
	"os"

	"github.com/inferablehq/inferable-go"
	"github.com/joho/godotenv"
)

func main() {
	// Load vars from .env file
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// Instantiate the Inferable client.
	client, err := inferable.New(inferable.InferableOptions{
		// To get a new key, run:
		// npx @inferable/cli auth keys create 'My New Machine Key' --type='cluster_machine'
		APISecret:   os.Getenv("INFERABLE_API_SECRET"),
		APIEndpoint: os.Getenv("INFERABLE_API_ENDPOINT"),
	})

	if err != nil {
		panic(err)
	}

	service, err := client.RegisterService("MyService")
	if err != nil {
		// Handle error
	}

	type MyInput struct {
		Message string `json:"message"`
	}

	myFunc := func(input MyInput) (string, error) {
		return "Hello, " + input.Message, nil
	}

	// Register a demo function
	err = service.RegisterFunc(inferable.Function{
		Func:        myFunc,
		Name:        "MyFunction",
		Description: "A simple greeting function",
	})

	if err != nil {
		panic(err)
	}

	service.Start()

	err = service.Start()
	if err != nil {
		panic(err)
	}

	defer service.Stop()

	// Wait for CTRL+C
	<-make(chan struct{})
}
