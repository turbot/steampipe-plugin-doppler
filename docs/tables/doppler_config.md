# Table: doppler_config

In Doppler, configs refer to the configuration files that define the settings, parameters, and environment variables used by your application or service. These files typically include information about database connections, API keys, authentication settings, and other runtime parameters.

## Example

### Basic info

```sql
select
  project,
  name,
  created_at,
  environment,
  root,
  locked
from
  doppler_config;
```

### Get environment details of each config

```sql
select
  c.name as config_name,
  c.created_at as config_created_at,
  c.root,
  e.name as environment_name,
  e.created_at as environment_created_at,
  e.slug
from
  doppler_config as c,
  doppler_environment as e
where
  e.slug = c.environment;
```

### List root configs

```sql
select
  project,
  name,
  created_at,
  environment,
  root
from
  doppler_config
where
  root;
```

### List locked configs

```sql
select
  project,
  name,
  created_at,
  locked
from
  doppler_config
where
  locked;
```

### List configs that are created in last 30 days

```sql
select
  name,
  created_at,
  initial_fetch_at,
  last_fetch_at
from
  doppler_config
where
  created_at >= now() - interval '30' day;
```

### Get project details for each config

```sql
select
  c.name config_name,
  c.project,
  c.environment,
  p.id as project_id,
  p.created_at as project_created_at,
  p.slug as project_slug
from
  doppler_config as c,
  doppler_project as p
where
  c.project = p.name;
```
