# gasmetrics

![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-2496ED?style=flat&logo=docker&logoColor=white)

Telegram bot for tracking home gas readings, which are stored in a PostgreSQL database.

## How to create a Telegram bot

In order to create a new bot:
1. Start a chat with `@BotFather`
2. Send `/newbot`
3. Choose a name and a username
4. Store the token

## How to run

Create the files to store the Telegram token and the PostgreSQL password:

```bash
mkdir secrets
printf '<token>' > secrets/telegram_token.txt
printf '<password>' > secrets/postgres_password.txt
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
