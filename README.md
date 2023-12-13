![image](https://hub.steampipe.io/images/plugins/turbot/doppler-social-graphic.png)

# Doppler Plugin for Steampipe

Use SQL to query projects, environments, secrets and more from Doppler.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/doppler)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/doppler/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-doppler/issues)

## Quick start

### Install

Download and install the latest Doppler plugin:

```bash
steampipe plugin install doppler
```

Configure your [credentials](https://hub.steampipe.io/plugins/turbot/doppler#credentials) and [config file](https://hub.steampipe.io/plugins/turbot/doppler#configuration).

Configure your account details in `~/.steampipe/config/doppler.spc`:

```hcl
connection "doppler" {
  plugin = "doppler"

  # Authentication information
  token = "dp.pt.abcdVDI7jCoV92ylJS9yXYZO5CZRiGm0WWWnZgsZZih"
  project_id = "plugin-test-project"
}
```

Or through environment variables:

```sh
export DOPPLER_TOKEN=dp.pt.abcdVDI7jCoV92ylJS9yXYZO5CZRiGm0WWWnZgsZZih
export DOPPLER_PROJECT_ID=plugin-test-project
```

Run steampipe:

```shell
steampipe query
```

Get your Doppler project:

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
| plugin-test-project | plugin-test-project | This is my first fancy project               | 2023-04-26T18:14:58+05:30 | steampipeplugin |
+---------------------+---------------------+----------------------------------------------+---------------------------+-----------------+
```

## Engines

This plugin is available for the following engines:

| Engine        | Description
|---------------|------------------------------------------
| [Steampipe](https://steampipe.io/docs) | The Steampipe CLI exposes APIs and services as a high-performance relational database, giving you the ability to write SQL-based queries to explore dynamic data. Mods extend Steampipe's capabilities with dashboards, reports, and controls built with simple HCL. The Steampipe CLI is a turnkey solution that includes its own Postgres database, plugin management, and mod support.
| [Postgres FDW](https://steampipe.io/docs/steampipe_postgres/overview) | Steampipe Postgres FDWs are native Postgres Foreign Data Wrappers that translate APIs to foreign tables. Unlike Steampipe CLI, which ships with its own Postgres server instance, the Steampipe Postgres FDWs can be installed in any supported Postgres database version.
| [SQLite Extension](https://steampipe.io/docs/steampipe_sqlite/overview) | Steampipe SQLite Extensions provide SQLite virtual tables that translate your queries into API calls, transparently fetching information from your API or service as you request it.
| [Export](https://steampipe.io/docs/steampipe_export/overview) | Steampipe Plugin Exporters provide a flexible mechanism for exporting information from cloud services and APIs. Each exporter is a stand-alone binary that allows you to extract data using Steampipe plugins without a database.
| [Turbot Pipes](https://turbot.com/pipes/docs) | Turbot Pipes is the only intelligence, automation & security platform built specifically for DevOps. Pipes provide hosted Steampipe database instances, shared dashboards, snapshots, and more.

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-doppler.git
cd steampipe-plugin-doppler
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/doppler.spc
```

Try it!

```
steampipe query
> .inspect doppler
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Open Source & Contributing

This repository is published under the [Apache 2.0](https://www.apache.org/licenses/LICENSE-2.0) (source code) and [CC BY-NC-ND](https://creativecommons.org/licenses/by-nc-nd/2.0/) (docs) licenses. Please see our [code of conduct](https://github.com/turbot/.github/blob/main/CODE_OF_CONDUCT.md). We look forward to collaborating with you!

[Steampipe](https://steampipe.io) is a product produced from this open source software, exclusively by [Turbot HQ, Inc](https://turbot.com). It is distributed under our commercial terms. Others are allowed to make their own distribution of the software, but cannot use any of the Turbot trademarks, cloud services, etc. You can learn more in our [Open Source FAQ](https://turbot.com/open-source).

## Get Involved

**[Join #steampipe on Slack →](https://turbot.com/community/join)**

Want to help but don't know where to start? Pick up one of the `help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Doppler Plugin](https://github.com/turbot/steampipe-plugin-doppler/labels/help%20wanted)
