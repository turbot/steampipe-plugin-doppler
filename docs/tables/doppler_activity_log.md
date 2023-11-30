---
title: "Steampipe Table: doppler_activity_log - Query Doppler Activity Logs using SQL"
description: "Allows users to query Doppler Activity Logs, specifically the event details, providing insights into user activities and potential security incidents."
---

# Table: doppler_activity_log - Query Doppler Activity Logs using SQL

Doppler is a service that provides a secure and universal secret manager. It enables developers and organizations to manage and securely access sensitive information such as API keys, credentials, and other secrets. Doppler Activity Logs are a feature of Doppler that records user activities, providing an audit trail that can be used for security and compliance purposes.

## Table Usage Guide

The `doppler_activity_log` table provides insights into user activities within Doppler. As a security analyst, explore user-specific activity details through this table, including event types, timestamps, and associated metadata. Utilize it to uncover information about user actions, such as those related to secret access, changes to environment configurations, and potential security incidents.

## Examples

### Basic info
Explore the history of user activity in your workplace by identifying when specific actions were taken and by whom. This can be useful for auditing purposes, understanding user behavior, or investigating specific incidents.

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
Explore recent user activity to gain insights into user engagement and interactions within the past month. This can help assess the dynamics of your workplace and understand user behavior patterns.

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
Identify instances where activities have been carried out by bots such as Doppler Bot and GitHub bot. This can help in monitoring automated processes and ensure they are functioning as expected.

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
Explore the latest activity in your Doppler environment by identifying the most recent action taken, providing insights into the user's actions and the timing of those actions. This is particularly useful for auditing and tracking changes made in your environment.

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
Uncover the details of the most active users across different projects by analyzing the frequency of their activities. This can help in identifying key contributors and understanding user engagement within each project.

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
Analyze the settings to understand the most frequent activities recorded. This can be beneficial in identifying trends or patterns in system use, aiding in operational efficiency and proactive issue resolution.

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