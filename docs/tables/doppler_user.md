---
title: "Steampipe Table: doppler_user - Query Doppler Users using SQL"
description: "Allows users to query Doppler Users, specifically the user information and settings, providing insights into user management and access control."
---

# Table: doppler_user - Query Doppler Users using SQL

Doppler is a universal secrets manager that helps you to manage and securely access secrets, such as database credentials, API keys, and tokens. It provides an intuitive interface for managing secrets across multiple environments and cloud platforms. Doppler ensures that secrets are automatically synced and updated across all applications, services, and users.

## Table Usage Guide

The `doppler_user` table provides insights into user accounts within Doppler. As a security engineer, explore user-specific details through this table, including user roles, access rights, and associated metadata. Utilize it to uncover information about users, such as their assigned roles, the access rights they have, and the verification of their access controls.

## Examples

### Basic info
Explore which users have access to your Doppler workspace, including when they were created and their associated email and username. This can help you manage user access and understand who has permissions within your workspace.

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
Explore which users have been added to your system in the past month. This can help you keep track of your growing user base and understand recent changes in your user demographics.

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
Discover the users who have owner access, which is crucial for understanding the distribution of administrative privileges in your workplace. This can be particularly useful in audits or when reevaluating permission structures.

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