# Table: doppler_user

Represents a user who has accepted membership in a project.

## Examples

### Basic info

```sql
select
  id,
  name,
  email,
  username,
  created_at,
  access,
  workplace_name
from
  doppler_user;
```

### List users that are created in the last 30 days

```sql
select
  id,
  name,
  email,
  username,
  created_at,
  access,
  workplace_name
from
  doppler_user
where
  created_at >= now() - interval '30 day';
```

### List users with owner access

```sql
select
  id,
  name,
  email,
  username,
  created_at,
  access,
  workplace_name
from
  doppler_user
where
  access = 'owner';
```
