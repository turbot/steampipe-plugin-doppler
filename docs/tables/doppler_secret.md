# Table: doppler_secret

Secrets in Doppler work very similarly to how they would on any other platform, with a few exceptions. The secrets engine has a couple of built-in perks!

- You **_must_** specify `config_name` in a `where` clause in order to use this table.

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
  doppler_secret
where
  config_name = 'dev';
```

### List secrets for a project

```sql
select
  s.project,
  s.config_name,
  s.secret_name,
  s.secret_value_raw
from
  doppler_secret s
  inner join
    doppler_config c
    on s.config_name = c.name
where
  s.project = 'example-project';
```

### List config details for each secret

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
  inner join
    doppler_config c
    on s.config_name = c.name;
```

### Count number of secrets by project and config

```sql
select
  s.project,
  s.config_name,
  count(s.secret_name)
from
  doppler_secret s
  inner join
    doppler_config c
    on s.config_name = c.name
group by
  s.project,
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
  e.initial_fetch_at as initial_fetch_at
from
  doppler_secret s
  inner join
    doppler_config c
    on s.config_name = c.name
  inner join
    doppler_environment e
    on e.slug = c.environment;
```
