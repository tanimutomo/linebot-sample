package bot

import (
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/tanimutomo/linebot-sample/pkg/config"
)

var bot *linebot.Client

func New() (*linebot.Client, error) {
	if bot != nil {
		return bot, nil
	}
	b, err := linebot.New(config.LINE.ChannelSecret, config.LINE.AccessToken)
	if err != nil {
		panic(err)
	}
	bot = b
	return bot, nil
}
