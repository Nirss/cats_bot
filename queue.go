package main

import (
	"log"
	"os"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

type Message struct {
	Chat *tb.Chat
	Tag  string
}

type MessageQueue struct {
	bot          *tb.Bot
	timeout      time.Duration
	MessageQueue chan Message
}

func NewMessageQueue(bot *tb.Bot, timout time.Duration) *MessageQueue {
	queue := &MessageQueue{
		bot:          bot,
		timeout:      timout,
		MessageQueue: make(chan Message, 100),
	}
	go queue.QueueWorker()
	return queue
}

func (m *MessageQueue) QueueWorker() {
	for msg := range m.MessageQueue {
		time.Sleep(m.timeout)
		joyUrl, err := GetRandomCats(msg.Tag)
		if err != nil {
			log.Println(err)
			_, err = m.bot.Send(msg.Chat, "Ошибка получения изображения")
			if err != nil {
				log.Println(err)
			}
			continue
		}
		filename, err := DownloadFile(joyUrl)
		if err != nil {
			log.Println(err)
			_, err = m.bot.Send(msg.Chat, "Ошибка распознавания изображения")
			if err != nil {
				log.Println(err)
			}
			continue
		}
		err = cropImage(filename)
		if err != nil {
			log.Println(err)
			continue
		}
		defer os.Remove(filename)
		image := &tb.Photo{File: tb.FromDisk(filename)}
		_, err = m.bot.Send(msg.Chat, image)
		if err != nil {
			log.Println(err)
			continue
		}
	}
}
