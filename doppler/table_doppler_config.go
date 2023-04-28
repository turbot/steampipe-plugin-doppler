package doppler

import (
	"context"

	"github.com/nikoksr/doppler-go"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDopplerConfig(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "doppler_config",
		Description: "Doppler config refer to the configuration files that define the settings, parameters, and environment variables used by your application or service",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"project", "name"}),
			IgnoreConfig: &plugin.IgnoreConfig{
				ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"Could not find requested project", "Could not find requested config"}),
			},
			Hydrate: getConfig,
		},
		List: &plugin.ListConfig{
			ParentHydrate: listProjects,
			Hydrate:       listConfigs,
			KeyColumns: plugin.KeyColumnSlice{
				{
					Name:    "project",
					Require: plugin.Optional,
				},
			},
		},
		Columns: commonColumnsForAllResource([]*plugin.Column{
			{
				Name:        "name",
				Description: "The name of the config.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "project",
				Description: "The ID of the project.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_at",
				Description: "Date and time of the object's creation.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "environment",
				Description: "Identifier of the environment that the config belongs to.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "initial_fetch_at",
				Description: "Date and time of the first secrets fetch.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "locked",
				Description: "Whether the config can be renamed and/or deleted.",
				Type:        proto.ColumnType_BOOL,
			},
			{
				Name:        "last_fetch_at",
				Description: "Date and time of the last secrets fetch.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "root",
				Description: "Whether the config is the root of the environment.",
				Type:        proto.ColumnType_BOOL,
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

func listConfigs(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	project := h.Item.(*doppler.Project)
	projectId := d.EqualsQualString("project")

	// Reduce the numbers of API call if the project id is provided in the where clause.
	if projectId != "" {
		if projectId != *project.ID {
			return nil, nil
		}
	}

	// Get client
	client, err := GetConfigClient(ctx, d.Connection)
	if err != nil {
		plugin.Logger(ctx).Error("doppler_config.listConfigs", "client_error", err)
		return nil, err
	}

	input := &doppler.ConfigListOptions{
		Project: *project.ID,
	}

	// The SDK does not support pagination till date(04/23).
	op, _, err := client.List(ctx, input)
	if err != nil {
		plugin.Logger(ctx).Error("doppler_config.listConfigs", "api_error", err)
		return nil, err
	}

	for _, config := range op {
		d.StreamListItem(ctx, config)

		// Context may get cancelled due to manual cancellation or if the limit has been reached.
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

//// HYDRATED FUNCTIONS

func getConfig(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	projectId := d.EqualsQualString("project")
	configName := d.EqualsQualString("name")

	// Empty Check
	if projectId == "" || configName == "" {
		return nil, nil
	}

	// Get client
	client, err := GetConfigClient(ctx, d.Connection)
	if err != nil {
		plugin.Logger(ctx).Error("doppler_config.getConfig", "client_error", err)
		return nil, err
	}

	input := &doppler.ConfigGetOptions{
		Project: projectId,
		Config:  configName,
	}

	op, _, err := client.Get(ctx, input)
	if err != nil {
		plugin.Logger(ctx).Error("doppler_config.getConfig", "api_error", err)
		return nil, err
	}

	if op != nil {
		return op, nil
	}

	return nil, nil
}
