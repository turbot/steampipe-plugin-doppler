package main

import (
	"github.com/turbot/steampipe-plugin-doppler/doppler"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: doppler.Plugin})
}
