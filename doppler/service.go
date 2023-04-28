package doppler

import (
	"context"

	"github.com/nikoksr/doppler-go"
	activitylog "github.com/nikoksr/doppler-go/activity_log"
	"github.com/nikoksr/doppler-go/audit"
	projectConfig "github.com/nikoksr/doppler-go/config"
	"github.com/nikoksr/doppler-go/environment"
	"github.com/nikoksr/doppler-go/project"
	"github.com/nikoksr/doppler-go/secret"
	servicetoken "github.com/nikoksr/doppler-go/service_token"
	"github.com/nikoksr/doppler-go/workplace"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// Project Client
func GetProjectClient(ctx context.Context, d *plugin.Connection) (*project.Client, error) {
	config := GetConfigWithToken(d)

	doppler.Key = *config.DOPPLER_TOKEN
	return &project.Client{
		Backend: doppler.GetBackend(),
		Key:     doppler.Key,
	}, nil

}

// Config Client
func GetConfigClient(ctx context.Context, d *plugin.Connection) (*projectConfig.Client, *string, error) {
	config := GetConfigWithToken(d)

	if *config.DOPPLER_TOKEN != "" {
		doppler.Key = *config.DOPPLER_TOKEN
		return &projectConfig.Client{
			Backend: doppler.GetBackend(),
			Key:     doppler.Key,
		}, config.PROJECT_ID, nil
	}

	return nil, nil, nil
}

// Secret Client
func GetSecretClient(ctx context.Context, d *plugin.Connection) (*secret.Client, *string, error) {
	config := GetConfigWithToken(d)

	if *config.DOPPLER_TOKEN != "" {
		doppler.Key = *config.DOPPLER_TOKEN
		return &secret.Client{
			Backend: doppler.GetBackend(),
			Key:     doppler.Key,
		}, config.PROJECT_ID, nil
	}

	return nil, nil, nil
}

// Environment Client
func GetEnvironmentClient(ctx context.Context, d *plugin.Connection) (*environment.Client, *string, error) {
	config := GetConfigWithToken(d)

	if *config.DOPPLER_TOKEN != "" {
		return &environment.Client{
			Backend: doppler.GetBackend(),
			Key:     *config.DOPPLER_TOKEN,
		}, config.PROJECT_ID, nil
	}

	return nil, nil, nil
}

// Service Token Client
func GetServiceTokenClient(ctx context.Context, d *plugin.Connection) (*servicetoken.Client, *string, error) {
	config := GetConfigWithToken(d)

	if *config.DOPPLER_TOKEN != "" {
		return &servicetoken.Client{
			Backend: doppler.GetBackend(),
			Key:     *config.DOPPLER_TOKEN,
		}, config.PROJECT_ID, nil
	}

	return nil, nil, nil
}

// User Client
func GetUserClient(ctx context.Context, d *plugin.Connection) (*audit.Client, *string, error) {
	config := GetConfigWithToken(d)

	if *config.DOPPLER_TOKEN != "" {
		return &audit.Client{
			Backend: doppler.GetBackend(),
			Key:     *config.DOPPLER_TOKEN,
		}, config.PROJECT_ID, nil
	}
	return nil, nil, nil
}

// Workplace Client
func GetWorkplaceClient(ctx context.Context, d *plugin.Connection) (*workplace.Client, error) {
	config := GetConfigWithToken(d)

	if *config.DOPPLER_TOKEN != "" {
		return &workplace.Client{
			Backend: doppler.GetBackend(),
			Key:     *config.DOPPLER_TOKEN,
		}, nil
	}

	return nil, nil
}

func GetActivityLogClient(ctx context.Context, d *plugin.Connection) (*activitylog.Client, *string, error) {
	config := GetConfigWithToken(d)

	if *config.DOPPLER_TOKEN != "" {
		return &activitylog.Client{
			Backend: doppler.GetBackend(),
			Key:     *config.DOPPLER_TOKEN,
		}, config.PROJECT_ID, nil
	}

	return nil, nil, nil
}
