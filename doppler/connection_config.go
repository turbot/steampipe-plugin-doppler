package doppler

import (
	"fmt"
	"os"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type dopplerConfig struct {
	Token *string `cty:"token"`
	ProjectId    *string `cty:"project_id"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"token": {
		Type: schema.TypeString,
	},
	"project_id": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &dopplerConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) dopplerConfig {
	if connection == nil || connection.Config == nil {
		return dopplerConfig{}
	}
	config, _ := connection.Config.(dopplerConfig)
	return config
}

// GetConfigWithToken :: If the user has not provided a token and project ID in the .spc file, the application should check if the 'DOPPLER_TOKEN' and 'DOPPLER_PROJECT_ID' environment variables have values set. If they do have values set, then these environment variables should be used to authenticate the application.

func GetConfigWithToken(connection *plugin.Connection) dopplerConfig {
	config := GetConfig(connection)
	token := os.Getenv("DOPPLER_TOKEN")
	dopplerProjectId := os.Getenv("DOPPLER_PROJECT_ID")

	if config.Token == nil && token == "" {
		errorMessage := fmt.Sprintf("Connection %s config does not have a valid token, update the token in the configuration file or environment variable DOPPLER_TOKEN and restart the steampipe.", connection.Name)
		panic(errorMessage)
	} else if config.Token == nil && token != "" {
		config.Token = &token
	} else if config.ProjectId == nil && dopplerProjectId == "" {
		errorMessage := fmt.Sprintf("Connection %s config does not have a valid Project ID, update the Project ID in the configuration file or environment variable DOPPLER_PROJECT_ID and restart the steampipe.", connection.Name)
		panic(errorMessage)
	} else if config.ProjectId == nil && dopplerProjectId != "" {
		config.ProjectId = &dopplerProjectId
	}

	return config
}
