package main

import (
	createmenu "create-rich-menu/create-menu"
	"fmt"
	"log"

	"github.com/line/line-bot-sdk-go/v8/linebot"
)

func main() {

	channelSecret := ""
	channelToken := ""

	linebotClient, err := linebot.New(channelSecret, channelToken)
	if err != nil {
		log.Panicf("unable initial linebot client: %v", err)
	}

	menus := createmenu.GetMenu("uat")

	for _, v := range menus {
		res, err := linebotClient.CreateRichMenu(v).Do()
		if err != nil {
			panic(err)
		}
		fmt.Println(v.Name, res.RichMenuID)
	}
}
