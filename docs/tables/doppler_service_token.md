# Table: doppler_service_token

Doppler Service Token.

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
  doppler_service_token 
where
  config = 'dev_aws';
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
  config = 'dev' 
  and access = 'read/write';
```

### List service tokens that never expire

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
  and expires_at is null;
```