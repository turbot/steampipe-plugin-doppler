package main

import (
	"context"
	"fmt"
	"log"

	"github.com/nikoksr/doppler-go"
	"github.com/nikoksr/doppler-go/project"
)

func main() {
	// Set your API key
	doppler.Key = ""

	// List all your secrets
	projects, _, err := project.List(context.Background(), &doppler.ProjectListOptions{})
	if err != nil {
		log.Fatalf("failed to list projects: %v", err)
	}

	for _, value := range projects {
		fmt.Printf("%s: \n", "Name", *value.Name)
		fmt.Printf("%s: \n", "ID", *value.ID)
		fmt.Printf("%s: \n", "created at", *value.CreatedAt)
		fmt.Printf("%s: \n", "Description", *value.Description)
	}
}
