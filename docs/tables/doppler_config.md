---
title: "Steampipe Table: doppler_config - Query Doppler Configurations using SQL"
description: "Allows users to query Doppler Configurations, specifically the details of the secrets and their respective projects, providing insights into secret management and access control."
---

# Table: doppler_config - Query Doppler Configurations using SQL

Doppler is a universal secret manager that helps developers manage and securely access secrets such as database credentials, API keys, and tokens. It provides a centralized way to store, access, and manage secrets for various applications and services. Doppler helps ensure that secrets are securely stored and are only accessible to authorized users.

## Table Usage Guide

The `doppler_config` table provides insights into configurations within Doppler's secret management. As a developer or DevOps engineer, explore configuration-specific details through this table, including secrets, their values, and the projects they belong to. Utilize it to manage and control access to secrets, ensuring that they are only accessible to authorized users.

## Examples

### Basic info
This query allows users to gain insights into the configurations of their Doppler projects. It can be used to understand the creation date, environment, root, and lock status of each project, which can be useful for project management and security purposes.

```sql+postgres
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

```sql+sqlite
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
Analyze the settings to understand the relationship between environments and their associated configurations. This is useful for auditing purposes and to ensure proper setup of your environments.

```sql+postgres
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

```sql+sqlite
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
Explore which configurations are set as root in your projects to gain insights into your system's setup. This is useful for assessing the elements within your environment that may require additional security or permissions due to their root status.

```sql+postgres
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

```sql+sqlite
select
  project,
  name,
  created_at,
  environment,
  root
from
  doppler_config
where
  root = 1;
```

### List configs that are locked
Discover the configurations that are locked, to understand which project settings are currently unmodifiable. This can be useful in managing project changes and identifying potential bottlenecks in the workflow.

```sql+postgres
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

```sql+sqlite
select
  project,
  name,
  created_at,
  locked
from
  doppler_config
where
  locked = 1;
```

### List configs that are created in last 30 days
Explore recent configurations by identifying those created within the past month. This aids in keeping track of changes and updates, ensuring your system remains up-to-date and secure.

```sql+postgres
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

```sql+sqlite
select
  name,
  created_at,
  initial_fetch_at,
  last_fetch_at
from
  doppler_config
where
  created_at >= datetime('now', '-30 day');
```

### Get project details for each config
Discover the segments that provide an overview of each project, including the environment and configuration details. This is useful for understanding the structure and organization of your projects.

```sql+postgres
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

```sql+sqlite
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