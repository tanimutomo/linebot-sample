package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/tanimutomo/linebot-sample/pkg/db"
	"github.com/tanimutomo/linebot-sample/pkg/receive"
	"github.com/tanimutomo/linebot-sample/pkg/send"
)

func main() {
	go func() {
		receive.RunServer()
	}()
	for {
		log.Println("pooling...")
		id := 1
		for {
			user := db.FindUser(id)
			if user == nil {
				break
			}
			log.Println("user found")
			sendRandomMessage(user)
			time.Sleep(10 * time.Second)
			id++
		}
		time.Sleep(3 * time.Second)
	}
}

func sendRandomMessage(user *db.User) {
	if rand.NormFloat64() > 0.0 {
		send.TextMessage(user.LINEUserID, fmt.Sprintf("Hello, %s", user.Name))
	} else {
		send.ButtonMessage(user.LINEUserID, "Are you hungry?")
	}
}
