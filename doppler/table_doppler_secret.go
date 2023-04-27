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

func tableDopplerSecret(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "doppler_secret",
		Description: "Doppler Secret",
		List: &plugin.ListConfig{
			ParentHydrate: listProjects,
			Hydrate:       listSecrets,
			// TODO: Uncomment the ignore config once the ignore config started working with parent hydrate.
			// IgnoreConfig: &plugin.IgnoreConfig{
			// 	ShouldIgnoreErrorFunc: shouldIgnoreErrors([]string{"Could not find requested config"}),
			// },
			KeyColumns: plugin.KeyColumnSlice{
				{
					Name:    "project_id",
					Require: plugin.Optional,
				},
				{
					Name:    "config_name",
					Require: plugin.Required,
				},
			},
		},
		Columns: commonColumnsForAllResource([]*plugin.Column{
			{
				Name:        "project_id",
				Description: "The ID of the project",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("ProjectID"),
			},
			{
				Name:        "config_name",
				Description: "The name of the config.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "secret_name",
				Description: "The secrect name.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "secret_value_raw",
				Description: "The raw value of the secret.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "secret_value_computed",
				Description: "The computed value of the secret.",
				Type:        proto.ColumnType_STRING,
			},

			// Doppler standard column
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("SecretName"),
			},
		}),
	}
}

type SecretInfo struct {
	ProjectID           string
	ConfigName          string
	SecretName          string
	SecretValueRaw      string
	SecretValueComputed string
}

//// LIST FUNCTION

func listSecrets(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	project := h.Item.(*doppler.Project)
	projectId := d.EqualsQualString("project_id")
	configName := d.EqualsQualString("config_name")

	// Reduce the numbers of API call if the project id is provided in the where clause.
	if projectId != "" {
		if projectId != *project.ID {
			return nil, nil
		}
	}

	// Empty check
	if configName == "" {
		return nil, nil
	}

	// Get client
	client, err := GetSecretClient(ctx, d.Connection)
	if err != nil {
		plugin.Logger(ctx).Error("doppler_secret.listSecrets", "client_error", err)
		return nil, err
	}

	input := &doppler.SecretListOptions{
		Project: *project.ID,
		Config:  configName,
	}

	// The SDK does not support pagination till date(04/23).
	op, _, err := client.List(ctx, input)
	// In the case of parent hydrate the ignore config is not behaving properly so we need to handle the not found error code here.
	if err != nil {
		if strings.Contains(err.Error(), "Could not find requested config") {
			return nil, nil
		}
		plugin.Logger(ctx).Error("doppler_secret.listSecrets", "api_error", err)
		return nil, err
	}

	for k, v := range op {
		d.StreamListItem(ctx, &SecretInfo{
			ProjectID:           *project.ID,
			ConfigName:          configName,
			SecretName:          k,
			SecretValueRaw:      *v.Raw,
			SecretValueComputed: *v.Computed,
		})

		// Context may get cancelled due to manual cancellation or if the limit has been reached.
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}
