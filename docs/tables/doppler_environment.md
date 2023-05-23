# Table: doppler_environment

By default, each project has three default environments for defining configuration at the root level: Development, Staging, and Production.

## Examples

### Basic info

```sql
select
  id,
  slug,
  name,
  project,
  created_at,
  workplace_name
from
  doppler_environment;
```

### List environments that are created in the last 30 days

```sql
select
  id,
  slug,
  name,
  project,
  created_at,
  workplace_name
from
  doppler_environment
where
  created_at >= now() - interval '30 day';
```

### List environments from a specific project

```sql
select
  id,
  slug,
  name,
  project,
  created_at,
  workplace_name
from
  doppler_environment
where
  project = 'example-project';
```