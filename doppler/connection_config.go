package doppler

import (
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
