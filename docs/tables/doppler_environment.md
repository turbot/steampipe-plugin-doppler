---
title: "Steampipe Table: doppler_environment - Query Doppler Environments using SQL"
description: "Allows users to query Doppler Environments, specifically retrieving critical data such as environment ID, name, and the associated project ID."
---

# Table: doppler_environment - Query Doppler Environments using SQL

Doppler is a universal secrets manager that helps developers manage and securely access environment secrets during application development. It provides a secure and scalable solution for storing and retrieving environment-specific configuration data. Doppler Environment is a specific resource within Doppler that represents a unique environment for a project, containing all the environment variables for that specific environment.

## Table Usage Guide

The `doppler_environment` table provides insights into Doppler Environments within a project. As a DevOps engineer, explore environment-specific details through this table, including environment ID, name, and the associated project ID. Utilize it to uncover information about each environment, such as the specific variables stored in each environment, aiding in efficient configuration management and security audits.

## Examples

### Basic info
Explore which Doppler environments are linked to specific projects, when they were created, and their associated workplace names. This can help assess the organization and distribution of your resources within a project.

```sql
select
  id,
  slug,
  name,
  project,
  created_at,
  workplace_name
from
  doppler_environment;
```

### List environments that are created in the last 30 days
Discover the segments that were established in the past month. This can provide insights into recent project developments and their associated environments, aiding in project management and planning.

```sql
select
  id,
  slug,
  name,
  project,
  created_at,
  workplace_name
from
  doppler_environment
where
  created_at >= now() - interval '30 day';
```

### List environments for a specific project
Explore which environments are associated with a specific project. This is useful for understanding the various settings and configurations tied to a particular project, aiding in project management and oversight.

```sql
select
  id,
  slug,
  name,
  project,
  created_at,
  workplace_name
from
  doppler_environment
where
  project = 'example-project';
```