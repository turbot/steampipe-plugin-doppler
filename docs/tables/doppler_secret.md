# Table: doppler_secret

Secrets in Doppler work very similarly to how they would on any other platform, with a few exceptions. The secrets engine has a couple of built-in perks!

## Example

### Basic info

```sql
select
  project,
  config_name,
  secret_name,
  secret_value_raw,
  secret_value_computed
from
  doppler_secret;
```

### Get config details for each secret

```sql
select
  s.project,
  s.secret_name,
  s.config_name,
  c.created_at as config_created_at,
  c.environment as config_environment,
  c.root as config_root
from
  doppler_secret s
  inner join doppler_config c on s.config_name = c.name;
```

### Count the number of secrets by config

```sql
select
  s.config_name,
  count(s.secret_name)
from
  doppler_secret s
  inner join
    doppler_config c
    on s.config_name = c.name
group by
  s.config_name;
```

### Get environment details of each secret

```sql
select
  s.project,
  s.secret_name,
  s.config_name,
  e.id as environment_id,
  c.environment as environment_name,
  e.created_at as environment_created_at,
  e.initial_fetch_at as environmant_initial_fetch_at
from
  doppler_secret s
  inner join
    doppler_config c
    on s.config_name = c.name
  inner join doppler_environment e on e.slug = c.environment;
```
