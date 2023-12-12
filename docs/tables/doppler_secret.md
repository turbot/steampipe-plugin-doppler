---
title: "Steampipe Table: doppler_secret - Query Doppler Secrets using SQL"
description: "Allows users to query Doppler Secrets, specifically information regarding secret data and its associated metadata, providing insights into secret management and potential security risks."
---

# Table: doppler_secret - Query Doppler Secrets using SQL

Doppler is a universal secrets manager that helps developers and organizations securely manage and quickly access secrets, such as database credentials, API keys, and tokens. It provides a centralized system to store, access, and deploy secrets across applications, services, and infrastructure. Doppler ensures that sensitive data is encrypted at rest, in transit, and even in use.

## Table Usage Guide

The `doppler_secret` table provides insights into secrets within Doppler. As a security engineer, explore secret-specific details through this table, including secret values, associated metadata, and secret versions. Utilize it to uncover information about secrets, such as those that are outdated, unused, or potentially exposed, and to verify the security of your secret management practices.

## Examples

### Basic info
Explore the configuration and security details of a project by identifying the raw and computed values of its secrets. This can be beneficial in assessing the project's security parameters and ensuring the integrity of the sensitive data it holds.

```sql+postgres
select
  project,
  config_name,
  secret_name,
  secret_value_raw,
  secret_value_computed
from
  doppler_secret;
```

```sql+sqlite
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
Explore configuration details associated with each secret to gain insights into their creation time and deployment environment, which can assist in auditing and managing application settings.

```sql+postgres
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

```sql+sqlite
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
Gain insights into the number of secrets tied to each configuration in your Doppler setup. This is useful for understanding the security scope of each configuration.

```sql+postgres
select
  s.config_name,
  count(s.secret_name)
from
  doppler_secret s
  inner join doppler_config c on s.config_name = c.name
group by
  s.config_name;
```

```sql+sqlite
select
  s.config_name,
  count(s.secret_name)
from
  doppler_secret s
  inner join doppler_config c on s.config_name = c.name
group by
  s.config_name;
```

### Get environment details of each secret
Discover the segments that provide insight into the specific environment details associated with each secret. This is beneficial for understanding the environmental context and timing of your secrets, which can aid in security and management tasks.

```sql+postgres
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
  inner join doppler_config c on s.config_name = c.name
  inner join doppler_environment e on e.slug = c.environment;
```

```sql+sqlite
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
  inner join doppler_config c on s.config_name = c.name
  inner join doppler_environment e on e.slug = c.environment;
```