---
organization: Turbot
category: ["security"]
icon_url: "/images/plugins/turbot/doppler.svg"
brand_color: "#5571EF"
display_name: "Doppler"
short_name: "doppler"
description: "Steampipe plugin to query projects, environments, secrets and more from Doppler."
og_description: "Query Doppler with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/doppler-social-graphic.png"
---

# Doppler + Steampipe

[Doppler](https://www.doppler.com/) is a cloud-native secrets management platform that securely centralizes and automates the management of sensitive data across teams and applications.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

List your Doppler secrets:

```sql
select
  project,
  config_name,
  secret_name,
  secret_value_computed
from
  doppler_secret;
```

```
+---------------------+-------------+---------------------+-----------------------+
| project             | config_name | secret_name         | secret_value_computed |
+---------------------+-------------+---------------------+-----------------------+
| plugin-test-project | dev_aws     | DOPPLER_PROJECT     | plugin-test-project   |
| plugin-test-project | stg_aws     | DOPPLER_CONFIG      | stg_aws               |
| plugin-test-project | sandbox     | DOPPLER_CONFIG      | sandbox               |
| plugin-test-project | sandbox     | DOPPLER_ENVIRONMENT | sandbox               |
| plugin-test-project | prd         | DOPPLER_CONFIG      | prd                   |
| plugin-test-project | prd         | DOPPLER_PROJECT     | plugin-test-project   |
| plugin-test-project | prd         | DOPPLER_ENVIRONMENT | prd                   |
+---------------------+-------------+---------------------+-----------------------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/doppler/tables)**

## Quick start

### Install

Download and install the latest Doppler plugin:

```sh
steampipe plugin install doppler
```

### Credentials

| Item        | Description                                                                                                                                                                 |
| ----------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Credentials | Doppler requires a `project ID` and an [Access Token](https://docs.doppler.com/reference/auth-token-formats)                                                                                   |
| Permissions | The permission scope of access tokens is limited to the projects or environments based on token type.                                                                       |
| Radius      | Each connection represents a single doppler Installation.                                                                                                                   |
| Resolution  | 1. Credentials explicitly set in a steampipe config file (`~/.steampipe/config/doppler.spc`)<br />2. Credentials specified in environment variables, e.g., `DOPPLER_TOKEN` and `DOPPLER_PROJECT_ID`. |

### Configuration

Installing the latest doppler plugin will create a config file (`~/.steampipe/config/doppler.spc`) with a single connection named `doppler`:

Configure your account details in `~/.steampipe/config/doppler.spc`:

```hcl
connection "doppler" {
  plugin = "doppler"

  # `doppler_token` (required) - To create an access token, refer to https://docs.doppler.com/docs/service-tokens.
  # Can also be set with the DOPPLER_TOKEN environment variable.
  # doppler_token = "dp.pt.abcdVDI7jCoV92ylJS9yXYZO5CZRiGm0WWWnZgsZZih"

  # `project_id` (required) - The doppler Project ID within a workplace.
  # Can also be set with the DOPPLER_PROJECT_ID environment variable.
  # project_id = "plugin-test-project"
}
```

Alternatively, you can also use the standard Doppler environment variables to obtain credentials **only if other arguments (`doppler_token` and `project_id`) are not specified** in the connection:

```sh
export DOPPLER_TOKEN=dp.pt.abcdVDI7jCoV92ylJS9yXYZO5CZRiGm0WWWnZgsZZih
export DOPPLER_PROJECT_ID=plugin-test-project
```
## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-doppler
- Community: [Slack Channel](https://steampipe.io/community/join)
