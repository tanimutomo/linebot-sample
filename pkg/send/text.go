package send

import (
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/tanimutomo/linebot-sample/pkg/bot"
)

func TextMessage(to, msg string) error {
	bot, err := bot.New()
	if err != nil {
		return err
	}

	text := linebot.NewTextMessage(msg)

	res, err := bot.PushMessage(to, text).Do()
	if err != nil {
		log.Println("send fail")
		log.Println("err:", err)
		return err
	}

	log.Println("send success")
	log.Println("res:", res)
	return nil
}
