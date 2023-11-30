---
title: "Steampipe Table: doppler_project - Query Doppler Projects using SQL"
description: "Allows users to query Doppler Projects, specifically the project ID, name, and associated environment variables."
---

# Table: doppler_project - Query Doppler Projects using SQL

Doppler is a universal secrets management platform that helps developers to securely manage and quickly access secrets, such as database credentials, API keys, and tokens, in any environment. It allows users to centralize and manage secrets across all applications and services, ensuring secure access and handling. Doppler Projects are a key resource within Doppler, acting as a container for environments and their associated secrets.

## Table Usage Guide

The `doppler_project` table provides insights into Projects within Doppler's secrets management platform. As a security engineer or developer, explore project-specific details through this table, including project ID, name, and associated environment variables. Utilize it to manage and monitor the security of secrets across different projects and environments.

## Examples

### Basic info
Explore the basic details of your projects, including their creation date and unique identifiers, to gain a better understanding of your project timeline and organization. This can be particularly useful for managing multiple projects and ensuring each one is progressing as expected.

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
Discover the projects that were initiated within a certain time frame, specifically between two given dates. This is useful for assessing the volume and pace of project creation during a specific period.

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
Explore particular project details to understand its creation time, description, and other key information. This is useful when you need a quick overview of a specific project without having to go through all the project data.

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