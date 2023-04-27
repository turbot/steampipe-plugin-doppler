package doppler

import (
	"context"

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
		// Get: &plugin.GetConfig{
		// 	KeyColumns: plugin.AllColumns([]string{"project_id", "config_name", "secret_name"}),
		// 	Hydrate:    getSecret,
		// },
		List: &plugin.ListConfig{
			ParentHydrate: listProjects,
			Hydrate:       listSecrets,
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
		Columns: []*plugin.Column{
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
				Description: "The title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("SecretKey"),
			},
		},
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
	if err != nil {
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

// //// HYDRATED FUNCTIONs

// func getSecret(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
// 	projectId := d.EqualsQualString("project_id")
// 	configName := d.EqualsQualString("config_name")
// 	secretName := d.EqualsQualString("secret_name")

// 	// Get client
// 	client, err := GetSecretClient(ctx, d.Connection)
// 	if err != nil {
// 		plugin.Logger(ctx).Error("doppler_secret.getSecret", "client_error", err)
// 		return nil, err
// 	}

// 	input := &doppler.SecretGetOptions{
// 		Project: projectId,
// 		Config:  configName,
// 		Name:    secretName,
// 	}

// 	op, _, err := client.Get(ctx, input)
// 	if err != nil {
// 		plugin.Logger(ctx).Error("doppler_secret.getSecret", "api_error", err)
// 		return nil, err
// 	}
// 	plugin.Logger(ctx).Error("OUTPUT ===>>>", op)
// 	if op != nil {
// 		return &SecretInfo{
// 			ProjectID:           projectId,
// 			ConfigName:          configName,
// 			SecretName:          *op.Name,
// 			SecretValueRaw:      *op.Value.Raw,
// 			SecretValueComputed: *op.Value.Computed,
// 		}, nil
// 	}

// 	return nil, nil
// }
