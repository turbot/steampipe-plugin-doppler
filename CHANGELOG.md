## v0.1.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters.
- Recompiled plugin with Go version `1.21`.

## v0.0.2 [2023-07-11]

_Bug fixes_

- Fixed the plugin's config argument to use `token` instead of `doppler_token` to align with the API documentation. ([#4](https://github.com/turbot/steampipe-plugin-doppler/pull/4))

## v0.0.1 [2023-06-22]

_What's new?_

- New tables added
  - [doppler_activity_log](https://hub.steampipe.io/plugins/turbot/doppler/tables/doppler_activity_log)
  - [doppler_config](https://hub.steampipe.io/plugins/turbot/doppler/tables/doppler_config)
  - [doppler_environment](https://hub.steampipe.io/plugins/turbot/doppler/tables/doppler_environment)
  - [doppler_project](https://hub.steampipe.io/plugins/turbot/doppler/tables/doppler_project)
  - [doppler_secret](https://hub.steampipe.io/plugins/turbot/doppler/tables/doppler_secret)
  - [doppler_service_token](https://hub.steampipe.io/plugins/turbot/doppler/tables/doppler_service_token)
  - [doppler_user](https://hub.steampipe.io/plugins/turbot/doppler/tables/doppler_user)
