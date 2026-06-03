# gasmetrics

![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white)

Home gas consuption tracker. Built as three independent services.

## gasmetrics-server

Containerized HTTP server that manages home gas readings.

### API Endpoints

The server listens on port `1024` in the container and does **not** expose any port on the host machine.

#### Get a list of readings

```
GET /
```

This endpoint accepts a `limit` query parameter for specifying how many entries to return:

```
GET /?limit={n}
```

**Response body:**

```
[
  {
    "id": 12,
    "value": 678,
    "recorded_at": "2026-05-30T12:08:48.197787+02:00"
  },
  {
    "id": 5,
    "value": 673,
    "recorded_at": "2026-05-22T13:48:34.590414+02:00"
  },
  {
    "id": 4,
    "value": 672,
    "recorded_at": "2026-05-20T09:24:54.176378+02:00"
  },
]
```

#### Get the average consumption per day

```
GET /stats
```

**Response body:**

```
{
  "avg": 1.472222222222222
}
```


#### Submit a new gas reading

```
POST /
```

#### Delete a reading

```
DELETE /{id}
```

## Telegram bot

Telegram bot that polls for new messages and forwards readings to the server via HTTP

### How to create a Bot

In order to create a new bot:
1. Start a chat with `@BotFather`
2. Send `/newbot`
3. Choose a name and a username
4. Store the token

## How to run the containers

Create the files to store the Telegram token and the PostgreSQL password:

```bash
mkdir secrets
printf '<token>' > gasmetrics-bot/secrets/telegram_token.txt
printf '<password>' > gasmetrics-server/secrets/postgres_password.txt
```

Run the containers:

```bash
docker compose up
```

## Troubleshooting

### PostgreSQL password

Make sure that the password for connecting to PostgreSQL is the same as the one used in the **postgres** service.

### Telegram token

To retrieve the token:
1. Send `/mybots` to `@BotFather`
2. Choose the bot from the list
3. Go into `API Token`
