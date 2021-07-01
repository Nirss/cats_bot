package main

import (
	"image/jpeg"
	"log"
	"os"
	"strings"
	"time"

	"github.com/disintegration/imaging"

	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	defaultCatsQuery = [6]string{"КОТИК", "КОТЯРА", "КОТИЩЕ", "КОТЭ", "CAT", "CATS"}
	littleCatsQuery  = [5]string{"КОТЕНОК", "КОТЁНОК", "КОТЕЙКА", "КОТЕНОЧЕК", "КОТЁНОЧЕК"}
	bigCatsQuery     = [2]string{"КОТЯРА", "КОТИЩЕ"}
	manulQuery       = [1]string{"МАНУЛ"}

	queue *MessageQueue
)

func telegramBot() {
	b, err := tb.NewBot(tb.Settings{
		Token:  "1411061889:AAEcTyodcCmaQOJMFIp49MVYwA1STHX8LK8",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	queue = NewMessageQueue(b, 100*time.Millisecond)

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle(tb.OnText, func(m *tb.Message) {
		tag := getTag(m.Text)
		if tag == "" {
			return
		}
		if m.Sender.ID == 607033288 {
			tag = "мейн-кун"
		}
		queue.MessageQueue <- Message{
			Chat: m.Chat,
			Tag:  tag,
		}
	})

	b.Start()
}

func cropImage(filename string) error {
	src, err := imaging.Open(filename)
	if err != nil {
		return err
	}
	var imageWith = src.Bounds().Dx()
	var imageHeight = src.Bounds().Dy()
	src = imaging.CropAnchor(src, imageWith, imageHeight-15, imaging.Top)
	image, err := os.Create(filename)

	if err != nil {
		panic(err)
	}
	defer image.Close()
	jpeg.Encode(image, src, nil)
	return nil
}

func getTag(query string) string {
	switch commandExist(query) {
	case "кот":
		return "котэ"
	case "котенок":
		return "личинка+котэ"
	case "котяра":
		return "большие+котэ"
	case "манул":
		return "манул"
	default:
		return ""
	}
}

func commandExist(query string) string {
	query = strings.ToUpper(query)
	for _, item := range defaultCatsQuery {
		if strings.Contains(query, item) {
			return "кот"
		}
	}
	for _, item := range bigCatsQuery {
		if strings.Contains(query, item) {
			return "котяра"
		}
	}
	for _, item := range littleCatsQuery {
		if strings.Contains(query, item) {
			return "котенок"
		}
	}
	for _, item := range manulQuery {
		if strings.Contains(query, item) {
			return "манул"
		}
	}
	return ""
}
