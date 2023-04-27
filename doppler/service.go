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
	config := GetConfig(d)

	if *config.APIKey != "" {
		doppler.Key = *config.APIKey
		return &project.Client{
			Backend: doppler.GetBackend(),
			Key:     doppler.Key,
		}, nil
	}

	return nil, nil
}

// Config Client
func GetConfigClient(ctx context.Context, d *plugin.Connection) (*projectConfig.Client, error) {
	config := GetConfig(d)

	if *config.APIKey != "" {
		doppler.Key = *config.APIKey
		return &projectConfig.Client{
			Backend: doppler.GetBackend(),
			Key:     doppler.Key,
		}, nil
	}

	return nil, nil
}

// Secret Client
func GetSecretClient(ctx context.Context, d *plugin.Connection) (*secret.Client, error) {
	config := GetConfig(d)

	if *config.APIKey != "" {
		doppler.Key = *config.APIKey
		return &secret.Client{
			Backend: doppler.GetBackend(),
			Key:     doppler.Key,
		}, nil
	}

	return nil, nil
}

// Environment Client
func GetEnvironmentClient(ctx context.Context, d *plugin.Connection) (*environment.Client, error) {
	config := GetConfig(d)

	if *config.APIKey != "" {
		return &environment.Client{
			Backend: doppler.GetBackend(),
			Key:     *config.APIKey,
		}, nil
	}

	return nil, nil
}

// Service Token Client
func GetServiceTokenClient(ctx context.Context, d *plugin.Connection) (*servicetoken.Client, error) {
	config := GetConfig(d)

	if *config.APIKey != "" {
		return &servicetoken.Client{
			Backend: doppler.GetBackend(),
			Key:     *config.APIKey,
		}, nil
	}

	return nil, nil
}

// User Client
func GetUserClient(ctx context.Context, d *plugin.Connection) (*audit.Client, error) {
	config := GetConfig(d)

	if *config.APIKey != "" {
		return &audit.Client{
			Backend: doppler.GetBackend(),
			Key:     *config.APIKey,
		}, nil
	}
	return nil, nil
}

// Workplace Client
func GetWorkplaceClient(ctx context.Context, d *plugin.Connection) (*workplace.Client, error) {
	config := GetConfig(d)

	if *config.APIKey != "" {
		return &workplace.Client{
			Backend: doppler.GetBackend(),
			Key:     *config.APIKey,
		}, nil
	}

	return nil, nil
}

func GetActivityLogClient(ctx context.Context, d *plugin.Connection) (*activitylog.Client, error) {
	config := GetConfig(d)

	if *config.APIKey != "" {
		return &activitylog.Client{
			Backend: doppler.GetBackend(),
			Key:     *config.APIKey,
		}, nil
	}

	return nil, nil
}

