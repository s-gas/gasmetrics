# gasmetrics

![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white)

Home gas consuption tracker. Built as three independent services.

## Services
- ### [gasmetrics-server](./gasmetrics-server)
  HTTP server that stores the readings in a PostgreSQL database.

- ### [gasmetrics-bot](./gasmetrics-bot)
  Telegram bot that polls for messages and sends the readings to the server via HTTP.

- ### [gasmetrics-cli](./gasmetrics-cli)
  CLI that sends the readings to the server via HTTP.
    

## How to create a Telegram bot

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

## How to run the CLI

```bash
gasmetrics-cli <value>
```

## Troubleshooting

### PostgreSQL password

Make sure that the password for connecting to PostgreSQL is the same as the one used in the **postgres** service.

### Telegram token

To retrieve the token:
1. Send `/mybots` to `@BotFather`
2. Choose the bot from the list
3. Go into `API Token`

### CLI Path

Make sure that `gasmetrics-cli` is in `PATH`. To add it:

```bash
sudo GOBIN=/usr/local/bin/ go install
```
