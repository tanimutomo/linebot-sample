package receive

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/tanimutomo/linebot-sample/pkg/bot"
)

func RunServer() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Group(func(r chi.Router) {
		r.Use(validateSignature)
		r.Post("/webhook", webhook)
	})
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})
	http.ListenAndServe(":3000", r)
}

func webhook(w http.ResponseWriter, r *http.Request) {
	bot, err := bot.New()
	if err != nil {
		return
	}

	events, err := bot.ParseRequest(r)
	if err != nil {
		log.Println("follow fail")
		log.Println("err:", err)
	}

	for _, e := range events {
		switch e.Type {
		case linebot.EventTypeFollow:
			follow(e)
		}
	}
	log.Println("webhook success")
}
