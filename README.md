# Dashboard Cronjob Backend

This project is the backend for a dashboard system that interacts with cronjobs.

## DB schema
https://dbdiagram.io/d/cronjob-6798e227263d6cf9a04c14d9

## Setup

To get started with the development setup:

1. Clone this repository:
```bash
  git clone https://github.com/krisnaadi/dashboard-cronjob-be
  cd dashboard-cronjob-be
```
2. import database/cronjob.sql

3. Copy .env
```bash
cp .env.example .env
```
4. Build the application

```bash
make build
```

5. Run the application
```bash
make run
```

To deploy using docker:

1. Clone this repository:
```bash
  git clone https://github.com/krisnaadi/dashboard-cronjob-be
  cd dashboard-cronjob-be
```
2. import database/cronjob.sql
3. Copy .env
```bash
cp .env.example .env
```
4. configure .env file
5. build docker image
```bash
docker compose build
```
6. run docker container
```bash
docker compose up -d
```