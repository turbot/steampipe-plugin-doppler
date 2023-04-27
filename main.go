// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	"github.com/nikoksr/doppler-go"
// 	"github.com/nikoksr/doppler-go/project"
// 	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
// )

package main

import (
	"github.com/turbot/steampipe-plugin-doppler/doppler"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: doppler.Plugin})
}

// func main() {

// 	// Create a new Doppler client with your API key and API secret.
// 	// client, err := doppler.GetBackend()
// 	// client := doppler.GetBackend()
// 	// Set your API key
// 	// doppler.Key = "dp.pt.BBS2eoMYCQW6fLv2fgysWTfT3nkR2cBSaap887XeRfV"

// 	// List all your secrets
// 	client, _ := GetProjectClient(context.Background(), &plugin.Connection{})
// 	projects, _, err := project.List(context.Background(), &doppler.ProjectListOptions{})
// 	if err != nil {
// 		log.Fatalf("failed to list projects: %v", err)
// 	}

// 	for _, value := range projects {
// 		fmt.Printf("%s: \n", "Name", *value.Name)
// 		fmt.Printf("%s: \n", "ID", *value.ID)
// 		fmt.Printf("%s: \n", "created at", *value.CreatedAt)
// 		fmt.Printf("%s: \n", "Description", *value.Description)
// 	}
// }
