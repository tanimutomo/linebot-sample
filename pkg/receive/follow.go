package receive

import (
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/tanimutomo/linebot-sample/pkg/bot"
	"github.com/tanimutomo/linebot-sample/pkg/db"
)

func follow(e *linebot.Event) {
	bot, err := bot.New()
	if err != nil {
		return
	}
	uid := e.Source.UserID
	luser, err := bot.GetProfile(uid).Do()
	if err != nil {
		log.Println("follow fail")
		log.Println("err:", err)
		return
	}
	user := db.CreateUser(luser.UserID, luser.DisplayName)
	log.Println("follow success")
	log.Println("user:", user)
}
