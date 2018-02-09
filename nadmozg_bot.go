package main

import (
	"log"
	"github.com/Syfaro/telegram-bot-api"
    //"strings"
	"strings"
    "regexp"
	//"time"
	//"math/rand"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("519585678:AAF5Eg-_ODEaGTgcIKhsbgPVKt4vfUb_sF4")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	//bot.GetUp

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		//sr := rand.NewSource(time.Now().UnixNano())
		//r1 := rand.New(sr)

        chatMessage := update.Message.Text
        result, err := regexp.MatchString("/^нет$/i", chatMessage)
        log.Println(result)
        if result && (err != nil) && probably(20){
 			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Пидора ответ!")
			//msg.ReplyToMessageID = update.Message.MessageID

			bot.Send(msg)
        	continue
		}

		if (len(strings.Split(update.Message.Text, " ")) == 1) && probably(10) {
			lex_redupl, err := huifma(update.Message.Text)
			if err == nil {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, lex_redupl)
				//msg.ReplyToMessageID = update.Message.MessageID

				bot.Send(msg)
			}
			continue
		}
		if probably(7 ) {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, gopnic_answer())
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}
	}
}