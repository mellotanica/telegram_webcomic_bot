package bot

import (
	"gopkg.in/tucnak/telebot.v2"
	"telegram_webcomic_bot/sources"
)


func startTasks(bot *telebot.Bot) {

	go sources.KeepFeedsUpdated(bot)

}