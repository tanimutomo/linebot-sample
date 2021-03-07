package send

import (
	"log"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/tanimutomo/linebot-sample/pkg/bot"
)

const thumbnailURL = "https://image.flaticon.com/icons/png/512/1791/1791327.png"

var choices = []choice{
	{
		id:   1,
		data: "action=choice&id=1",
		name: "definitely",
	},
	{
		id:   2,
		data: "action=choice&id=2",
		name: "yes",
	},
	{
		id:   3,
		data: "action=choice&id=3",
		name: "no",
	},
	{
		id:   4,
		data: "action=choice&id=4",
		name: "fu*k",
	},
}

type choice struct {
	id    int
	label string
	data  string
	name  string
}

func ButtonMessage(to, msg string) {
	bot, _ := bot.New()

	acts := make([]linebot.TemplateAction, len(choices))
	for i, c := range choices {
		acts[i] = linebot.NewPostbackAction(c.name, c.data, "", c.name)
	}
	btemp := linebot.NewButtonsTemplate(thumbnailURL, "Choose one", msg, acts...)
	tmsg := linebot.NewTemplateMessage("Hungry Or Not", btemp)

	_, err := bot.PushMessage(to, tmsg).Do()
	if err != nil {
		log.Println("button fail")
		log.Println("err:", err)
	}

	log.Println("button success")
}
