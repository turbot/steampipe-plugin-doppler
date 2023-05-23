# Table: doppler_activity_log

From adding team members, to modifying secrets - Doppler will generate logs for every action your team makes. Activity Log table helps you to query the activity logs.

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
  created_at >= now() - interval '30 day';
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

### Get the most recent action

```sql
select
  id,
  text,
  created_at,
  environment,
  user_email
from
  doppler_activity_log
order by
  created_at desc limit 1;
```

### List most common actors

```sql
select
  project,
  user_name,
  user_email,
  count(*)
from
  doppler_activity_log
group by
  project,
  user_name,
  user_email
order by
  count desc;
```

### List out the most common activities

```sql
select
  id,
  text,
  created_at
from
  doppler_activity_log
order by
  created_at desc;
```