// Copyright 2016 LINE Corporation
//
// LINE Corporation licenses this file to you under the Apache License,
// version 2.0 (the "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at:
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package main

import (
	"fmt"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"log"
	"net/http"
	"os"
)

func main() {
	bot, err := linebot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTP Server for receiving requests from LINE platform
	http.HandleFunc("/callback", func(w http.ResponseWriter, req *http.Request) {
		events, err := bot.ParseRequest(req)
		if err != nil {
			if err == linebot.ErrInvalidSignature {
				w.WriteHeader(400)
			} else {
				w.WriteHeader(500)
			}
			return
		}
		for _, event := range events {
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					if message.Text == "時間與地點" {
						replyTextMessage := linebot.NewTextMessage(fmt.Sprintf("Hi hi, Elaine 與 豆哥的婚禮將在 2022/12/17 於寒舍艾麗酒店\n11073台灣台北市信義區松高路18號舉辦\n歡迎一同共襄盛舉！"))
						replyLocationMessage := linebot.NewLocationMessage(
							"寒舍艾麗酒店",
							"11073台灣台北市信義區松高路18號",
							25.038710596389304,
							121.56735086587742,
						)
						if _, err = bot.ReplyMessage(event.ReplyToken, replyTextMessage, replyLocationMessage).Do(); err != nil {
							log.Print(err)
						}
					} else if message.Text == "大聲說出想對我們說的話吧！" {
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("收到你想對我們說的話囉🥸\n祝平安順心。")).Do(); err != nil {
							log.Print(err)
						}
					} else if message.Text == "婚紗照" {
						picture1 := linebot.NewImageMessage("https://unsplash.com/photos/8vaQKYnawHw", "https://unsplash.com/photos/8vaQKYnawHw")
						picture2 := linebot.NewImageMessage("https://unsplash.com/photos/M2T1j-6Fn8w", "https://unsplash.com/photos/M2T1j-6Fn8w")
						picture3 := linebot.NewImageMessage("https://unsplash.com/photos/K8KiCHh4WU4", "https://unsplash.com/photos/K8KiCHh4WU4")
						picture4 := linebot.NewImageMessage("https://unsplash.com/photos/PuXtB1B4zL8", "https://unsplash.com/photos/PuXtB1B4zL8")
						picture5 := linebot.NewImageMessage("https://unsplash.com/photos/NO8Sj4dKE8k", "https://unsplash.com/photos/NO8Sj4dKE8k")
						if _, err = bot.ReplyMessage(event.ReplyToken, picture1, picture2, picture3, picture4, picture5, linebot.NewTextMessage("更多照片請看 https://unsplash.com/")).Do(); err != nil {
							log.Print(err)
						}
					} else {
						if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
							log.Print(err)
						}
					}
				case *linebot.StickerMessage:
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewStickerMessage("1", "1")).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	})
	// This is just sample code.
	// For actual use, you must support HTTPS by using `ListenAndServeTLS`, a reverse proxy or something else.
	log.Printf("About to listen on %s. Go to https://127.0.0.1:%s/", os.Getenv("PORT"), os.Getenv("PORT"))
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), nil); err != nil {
		log.Fatal(err)
	}
}
