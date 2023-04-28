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

// GetConfigWithToken :: if user doesn't specify the token in the .spc file, we should rely on the environment variable 'DOPPLER_TOKEN' if any value is set.

func GetConfigWithToken(connection *plugin.Connection) dopplerConfig {
	config := GetConfig(connection)
	dopplerToken := os.Getenv("DOPPLER_TOKEN")
	dopplerProjectId := os.Getenv("DOPPLER_PROJECT_ID")
	if config.DOPPLER_TOKEN == nil && dopplerToken == "" {
		errorMessage := fmt.Sprintf("Connection %s config does not have token, or does not have a valid token set in environment variable DOPPLER_TOKEN, please add the the token and restart the seampipe.", connection.Name)
		panic(errorMessage)
	} else if config.DOPPLER_TOKEN == nil && dopplerToken != "" {
		config.DOPPLER_TOKEN = &dopplerToken
	} else if config.PROJECT_ID == nil && dopplerProjectId == "" {
		errorMessage := fmt.Sprintf("Connection %s config does not have priject ID, or does not have a valid project ID set in environment variable DOPPLER_PROJECT_ID, please add the the project ID and restart the seampipe.", connection.Name)
		panic(errorMessage)
	} else if config.PROJECT_ID == nil && dopplerProjectId != "" {
		config.PROJECT_ID = &dopplerProjectId
	}

	return config
}
