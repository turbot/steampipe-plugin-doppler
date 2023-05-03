# Table: doppler_service_token

A Doppler Service Token provides read-only secrets access to a specific config within a project.

## Examples

### Basic info

```sql
select
  name,
  slug,
  access,
  expires_at,
  created_at,
  config,
  environment,
  project,
  workplace_name
from
  doppler_service_token;
```

### List service tokens that are expiring in the next 30 days

```sql
select
  name,
  slug,
  access,
  expires_at,
  created_at,
  config,
  environment,
  project,
  workplace_name
from
  doppler_service_token
where
  config = 'dev'
  and expires_at < now() + interval '30 day';
```

### List service tokens with read/write access

```sql
select
  name,
  slug,
  access,
  expires_at,
  created_at,
  config,
  environment,
  project,
  workplace_name
from
  doppler_service_token
where
  access = 'read/write';
```

### List service tokens that never expires

```sql
select
  name,
  slug,
  access,
  expires_at,
  created_at,
  config,
  environment,
  project,
  workplace_name
from
  doppler_service_token
where
  expires_at is null;
```
