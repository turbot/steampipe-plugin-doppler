package doppler

import (
	"context"

	"github.com/nikoksr/doppler-go"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDopplerUser(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "doppler_user",
		Description: "Doppler User",
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getUser,
		},
		List: &plugin.ListConfig{
			Hydrate: listUsers,
		},
		Columns: commonColumnsForAllResource([]*plugin.Column{
			{
				Name:        "id",
				Description: "ID is the unique identifier for the object.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "name",
				Description: "The user's name.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("User.Name"),
			},
			{
				Name:        "access",
				Description: "The description of the project.",
				Type:        proto.ColumnType_STRING,
			},
			{
				Name:        "created_at",
				Description: "The time when the project was created.",
				Type:        proto.ColumnType_TIMESTAMP,
			},
			{
				Name:        "email",
				Description: "The user's email address.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("User.Email"),
			},
			{
				Name:        "profile_image_url",
				Description: "The user's profile image URL.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("User.ProfileImageURL"),
			},
			{
				Name:        "username",
				Description: "The user's username.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("User.UserName"),
			},

			// Doppler standard column
			{
				Name:        "title",
				Description: ColumnDescriptionTitle,
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("User.Name"),
			},
		}),
	}
}

//// LIST FUNCTION

func listUsers(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	// Get client
	client, _, err := GetUserClient(ctx, d.Connection)
	if err != nil {
		plugin.Logger(ctx).Error("doppler_user.listUsers", "client_error", err)
		return nil, err
	}

	input := &doppler.AuditWorkplaceUserListOptions{}

	// The SDK does not support pagination till date(04/23).
	op, _, err := client.WorkplaceUserList(ctx, input)
	if err != nil {
		plugin.Logger(ctx).Error("doppler_user.listUsers", "api_error", err)
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

//// HYDRATED FUNCTIONS

func getUser(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	id := d.EqualsQualString("id")

	// Empty check
	if id == "" {
		return nil, nil
	}

	// Get client
	client, _, err := GetUserClient(ctx, d.Connection)
	if err != nil {
		plugin.Logger(ctx).Error("doppler_user.getProject", "client_error", err)
		return nil, err
	}

	input := &doppler.AuditWorkplaceUserGetOptions{
		UserID: id,
	}

	op, _, err := client.WorkplaceUserGet(ctx, input)
	if err != nil {
		plugin.Logger(ctx).Error("doppler_project.getProject", "api_error", err)
		return nil, err
	}

	return op, nil
}
