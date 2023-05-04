package doppler

import (
	"fmt"
	"os"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type dopplerConfig struct {
	DOPPLER_TOKEN *string `cty:"doppler_token"`
	PROJECT_ID    *string `cty:"project_id"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"doppler_token": {
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
	dopplerToken := os.Getenv("DOPPLER_TOKEN")
	dopplerProjectId := os.Getenv("DOPPLER_PROJECT_ID")

	if config.DOPPLER_TOKEN == nil && dopplerToken == "" {
		errorMessage := fmt.Sprintf("Connection %s config does not have a valid token, update the token in the configuration file or environment variable DOPPLER_TOKEN and restart the steampipe.", connection.Name)
		panic(errorMessage)
	} else if config.DOPPLER_TOKEN == nil && dopplerToken != "" {
		config.DOPPLER_TOKEN = &dopplerToken
	} else if config.PROJECT_ID == nil && dopplerProjectId == "" {
		errorMessage := fmt.Sprintf("Connection %s config does not have a valid Project ID, update the Project ID in the configuration file or environment variable DOPPLER_PROJECT_ID and restart the steampipe.", connection.Name)
		panic(errorMessage)
	} else if config.PROJECT_ID == nil && dopplerProjectId != "" {
		config.PROJECT_ID = &dopplerProjectId
	}

	return config
}
