---
organization: Turbot
category: ["security"]
icon_url: "/images/plugins/turbot/doppler.svg"
brand_color: ""
display_name: "Doppler"
short_name: "doppler"
description: "Steampipe plugin to query projects, environments, secrets and more from Doppler."
og_description: "Query Doppler with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/doppler-social-graphic.png"
---

# Doppler + Steampipe

[Doppler](https://www.doppler.com/) is a cloud-native secrets management platform that securely centralizes and automates the management of sensitive data across teams and applications.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

List your Doppler projects:

```sql
select
  name,
  id,
  description,
  created_at,
  workplace_name
from
  doppler_project;
```

```
+---------------------+---------------------+----------------------------------------------+---------------------------+-----------------+
| name                | id                  | description                                  | created_at                | workplace_name  |
+---------------------+---------------------+----------------------------------------------+---------------------------+-----------------+
| example-project     | example-project     | An example project with some sample secrets. | 2023-04-26T17:59:48+05:30 | steampipeplugin |
| plugin-test-project | plugin-test-project | This is my first fancy project               | 2023-04-26T18:14:58+05:30 | steampipeplugin |
+---------------------+---------------------+----------------------------------------------+---------------------------+-----------------+
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

| Item        | Description                                                                                                                                                                                                                                                                                                      |
| ----------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Credentials | Doppler requires an [Access Token](https://docs.doppler.com/reference/auth-token-formats)                                                                                   |
| Permissions | The permission scope of access tokens is limited to the projects or environments based on token type.                                                                       |
| Radius      | Each connection represents a single doppler Installation.                                                                                                                   |
| Resolution  | 1. Credentials explicitly set in a steampipe config file (`~/.steampipe/config/doppler.spc`)<br />2. Credentials specified in environment variables, e.g., `DOPPLER_TOKEN`. |

### Configuration

Installing the latest doppler plugin will create a config file (`~/.steampipe/config/doppler.spc`) with a single connection named `doppler`:

Configure your account details in `~/.steampipe/config/doppler.spc`:

```hcl
connection "doppler" {
  plugin = "doppler"

  # `doppler_token` (required) - To create an access token, refer to https://docs.doppler.com/docs/service-tokens
  # Can also be set with the DOPPLER_TOKEN environment variable.
  # doppler_token = "dp.pt.abcdVDI7jCoV92ylJS9yXYZO5CZRiGm0WWWnZgsZZih"

}
```

## Configuring Doppler Credentials

### Access Token Credentials

You may specify the Access Token to authenticate:

- `doppler_token`: Specify the access token, either a personal or service token.

```hcl
connection "doppler" {
  plugin = "doppler"
  doppler_token = "dp.pt.abcdVDI7jCoV92ylJS9yXYZO5CZRiGm0WWWnZgsZZih"
}
```

### Credentials from Environment Variables

The doppler plugin will use the doppler environment variable to obtain credentials in the connection:

```sh
export DOPPLER_TOKEN="dp.pt.abcdVDI7jCoV92ylJS9yXYZO5CZRiGm0WWWnZgsZZih"
```

## Get involved

- Open source: https://github.com/turbot/steampipe-plugin-doppler
- Community: [Slack Channel](https://steampipe.io/community/join)
