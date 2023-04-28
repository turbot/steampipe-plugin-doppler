package doppler

import (
	"context"

	"github.com/nikoksr/doppler-go"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDopplerActivityLog(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "doppler_activity_log",
		Description: "Doppler Activity Log",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getActivityLog,
		},
		List: &plugin.ListConfig{
			Hydrate: listActivityLogs,
		},
		Columns: commonColumnsForAllResource([]*plugin.Column{

			{
				Name:        "id",
				Description: "ID is the unique identifier for the activity log.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "text",
				Description: "Text describing the event.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "user_name",
				Description: "The user's name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("User.Name"),
			},
			{
				Name:        "config",
				Description: "Config is the config's name.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_at",
				Description: "The time when the project was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "environment",
				Description: "Environment is the environment's unique identifier.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "project",
				Description: "Project is the project that triggered the event.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "user_email",
				Description: "The user's email address.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("User.Email"),
			},
			{
				Name:        "user_username",
				Description: "The user's username.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("User.UserName"),
			},
			{
				Name:        "user_profile_image_url",
				Description: "The user's profile image URL.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("User.ProfileImageURL"),
			},
		}),
	}
}

//// LIST FUNCTION

func listActivityLogs(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Get client
	client, _, err := GetActivityLogClient(ctx, d.Connection)
	if err != nil {
		plugin.Logger(ctx).Error("doppler_activity_log.listActivityLogs", "client_error", err)
		return nil, err
	}

	input := &doppler.ActivityLogListOptions{}

	// The SDK does not support pagination till date(04/23).
	op, _, err := client.List(ctx, input)
	if err != nil {
		plugin.Logger(ctx).Error("doppler_activity_log.listActivityLogs", "api_error", err)
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

func getActivityLog(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	id := d.EqualsQualString("id")

	// Get client
	client, _, err := GetActivityLogClient(ctx, d.Connection)
	if err != nil {
		plugin.Logger(ctx).Error("doppler_user.getActivityLog", "client_error", err)
		return nil, err
	}

	input := &doppler.ActivityLogGetOptions{
		ID: id,
	}

	op, _, err := client.Get(ctx, input)
	if err != nil {
		plugin.Logger(ctx).Error("doppler_activity_log.getActivityLog", "api_error", err)
		return nil, err
	}

	return op, nil
}
