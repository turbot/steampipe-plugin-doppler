package doppler

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-doppler",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		ConnectionKeyColumns: []plugin.ConnectionKeyColumn{
			{
				Name:    "workplace_id",
				Hydrate: getWorkspaceIdForConnection,
			},
		},
		DefaultTransform: transform.FromGo().NullIfZero(),
		TableMap: map[string]*plugin.Table{
			"doppler_activity_log":  tableDopplerActivityLog(ctx),
			"doppler_config":        tableDopplerConfig(ctx),
			"doppler_environment":   tableDopplerEnvironment(ctx),
			"doppler_project":       tableDopplerProject(ctx),
			"doppler_secret":        tableDopplerSecret(ctx),
			"doppler_service_token": tableDopplerServiceToken(ctx),
			"doppler_user":          tableDopplerUser(ctx),
		},
	}
	return p
}
