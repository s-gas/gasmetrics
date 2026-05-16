package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func Reply(s string, bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, s)
	msg.ReplyToMessageID = update.Message.MessageID
	bot.Send(msg)
}
