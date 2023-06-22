# Table: doppler_project

A project in Doppler is where you define the app config and secrets for a single service or application.

## Examples

### Basic info

```sql
select
  id,
  name,
  created_at,
  slug,
  description
from
  doppler_project;
```

### List projects created between two specific dates

```sql
select
  name,
  id,
  created_at
from
  doppler_project
where
  created_at between '2022-01-01' and '2022-12-31';
```

### Get particular project details

```sql
select
  id,
  name,
  created_at,
  description,
  slug
from
  doppler_project
where
  name = 'my_first_project';
```
