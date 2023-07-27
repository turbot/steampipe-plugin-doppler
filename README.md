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

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-doppler/blob/main/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Doppler Plugin](https://github.com/turbot/steampipe-plugin-doppler/labels/help%20wanted)
