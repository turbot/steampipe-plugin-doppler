package doppler

import (
	"context"
	"strings"

	"github.com/nikoksr/doppler-go"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDopplerServiceToken(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "doppler_service_token",
		Description: "Doppler Service Token",
		List: &plugin.ListConfig{
			// ParentHydrate: listProjects,
			ParentHydrate: listConfigs,
			Hydrate:       listServiceTokens,
			// TODO: Uncomment the ignore config once the ignore config started working with parent hydrate.
			// IgnoreConfig: &plugin.IgnoreConfig{
			// 	ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"Could not find requested config"}),
			// },
			KeyColumns: plugin.KeyColumnSlice{
				// {
				// 	Name:    "project",
				// 	Require: plugin.Optional,
				// },
				{
					Name:    "config",
					Require: plugin.Optional,
				},
			},
		},
		Columns: commonColumnsForAllResource([]*plugin.Column{
			{
				Name:        "name",
				Description: "Name of the service token.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "slug",
				Description: "A unique identifier of the service token.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "access",
				Description: "The access level of the service token. One of read, read/write.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_at",
				Description: "Date and time of the object's creation.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "config",
				Description: "The name of the config.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "environment",
				Description: "Unique identifier for the environment object.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "expires_at",
				Description: "Date and time of the token's expiration, or null if token does not auto-expire.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "key",
				Description: "An API key that is used for authentication. Only available when creating the token.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "project",
				Description: "Unique identifier for the project object.",
				Type:        proto.ColumnType_STRING,
			},

			// Doppler standard column
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Name"),
			},
		}),
	}
}

//// LIST FUNCTION

func listServiceTokens(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	config := h.Item.(*doppler.Config)
	// projectId := d.EqualsQualString("project")
	configName := d.EqualsQualString("config")

	// Reduce the numbers of API call if the project id is provided in the where clause.
	if configName != "" {
		if configName != *config.Name {
			return nil, nil
		}
	}

	// Empty check
	if configName == "" {
		return nil, nil
	}

	// Get client
	client, projectId, err := GetServiceTokenClient(ctx, d.Connection)
	if err != nil {
		plugin.Logger(ctx).Error("doppler_service_token.listServiceTokens", "client_error", err)
		return nil, err
	}

	input := &doppler.ServiceTokenListOptions{
		Project: *projectId,
		Config:  *config.Name,
	}

	// The SDK does not support pagination till date(04/23).
	op, _, err := client.List(ctx, input)
	if err != nil {
		if strings.Contains(err.Error(), "Could not find requested config") {
			return nil, nil
		}
		plugin.Logger(ctx).Error("doppler_service_token.listServiceTokens", "api_error", err)
		return nil, err
	}

	for _, item := range op {
		d.StreamListItem(ctx, item)

		// Context may get cancelled due to manual cancellation or if the limit has been reached.
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
