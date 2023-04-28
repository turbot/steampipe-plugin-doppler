# Table: doppler_user

Doppler user.

## Examples

### Basic info

```sql
select
  id,
  text,
  created_at,
  user_name,
  user_email,
  workplace_name 
from
  doppler_activity_log;
```

### List activity within the last 30 days

```sql
select
  id,
  text,
  created_at,
  user_name,
  user_email,
  workplace_name 
from
  doppler_activity_log 
where
  created_at > now() - interval '30 day';
```

### List activity from bots like Doppler Bot and GitHub bot

```sql
select
  id,
  text,
  created_at,
  user_name,
  workplace_name 
from
  doppler_activity_log 
where
  user_name like '%Bot%';
```