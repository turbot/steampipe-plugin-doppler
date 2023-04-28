package doppler

import (
	"os"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type dopplerConfig struct {
	DOPPLER_TOKEN *string `cty:"doppler_token"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"doppler_token": {
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
	if config.DOPPLER_TOKEN == nil && dopplerToken == "" {
		return dopplerConfig{}
	} else if config.DOPPLER_TOKEN == nil && dopplerToken != "" {
		config.DOPPLER_TOKEN = &dopplerToken
	}

	return config
}
