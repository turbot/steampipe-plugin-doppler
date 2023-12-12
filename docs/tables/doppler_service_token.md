---
title: "Steampipe Table: doppler_service_token - Query Doppler Service Tokens using SQL"
description: "Allows users to query Doppler Service Tokens, specifically the access and refresh tokens, providing insights into token expiry and usage patterns."
---

# Table: doppler_service_token - Query Doppler Service Tokens using SQL

Doppler is a universal secret manager that helps you easily access and manage your application's secrets. It provides a centralized way to store, retrieve, and manage secrets for various applications, environments, and infrastructures. Doppler helps you maintain the security of your secrets and ensures that they are available to your applications when needed.

## Table Usage Guide

The `doppler_service_token` table provides insights into the service tokens within Doppler's universal secret manager. As a DevOps engineer, explore token-specific details through this table, including token expiry and usage patterns. Utilize it to uncover information about tokens, such as those nearing expiry, the frequency of token refresh, and the overall usage patterns.

## Examples

### Basic info
Explore the specifics of your service tokens, such as their accessibility, expiration, and creation dates, as well as their configuration and associated projects. This can help you manage your tokens effectively, ensuring they are up-to-date and used in the correct environments.

```sql+postgres
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

```sql+sqlite
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
Identify instances where service tokens are set to expire within the next month. This is useful for staying ahead of potential access issues and ensuring uninterrupted service.

```sql+postgres
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
  and expires_at <= now() + interval '30 day';
```

```sql+sqlite
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
  and expires_at <= datetime('now', '+30 day');
```

### List service tokens with read/write access
Discover the service tokens that have read/write access. This is useful for identifying potential security risks, as these tokens can modify your data.

```sql+postgres
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

```sql+sqlite
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
Explore which service tokens within your project are set to never expire. This can help identify potential security risks and enforce best practices for token management.

```sql+postgres
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

```sql+sqlite
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