package createmenu

import (
	"fmt"
	"strings"

	"github.com/line/line-bot-sdk-go/v8/linebot"
)

func GetMenu(env string) []linebot.RichMenu {

	var (
		linkageName string
		nonPortName string
		portName    string
		chatBarText string
		lineLiffURL string
	)

	envUpper := strings.ToUpper(env)
	linkageName = fmt.Sprintf("LINKAGE %s", envUpper)
	nonPortName = fmt.Sprintf("NON PORTFOLIO %s", envUpper)
	portName = fmt.Sprintf("PORTFOLIO %s", envUpper)
	chatBarText = fmt.Sprintf("Menu %s", envUpper)

	switch env {
	case "dev":
		lineLiffURL = "https://liff.line.me/2006474280-Xx6D8vKL"
	case "sit":
		lineLiffURL = "https://liff.line.me/2006570555-r54YW6NZ"
	case "uat":
		lineLiffURL = "https://liff.line.me/2006570573-4mzkPBLq"
	}

	size := linebot.RichMenuSize{
		Width:  2500,
		Height: 1686,
	}

	areasLinkage := []linebot.AreaDetail{
		{
			Bounds: linebot.RichMenuBounds{
				X:      834,
				Y:      843,
				Width:  834,
				Height: 843,
			},
			Action: linebot.RichMenuAction{
				Type: linebot.RichMenuActionTypeURI,
				URI:  lineLiffURL,
			},
		},
		{
			Bounds: linebot.RichMenuBounds{
				X:      1668,
				Y:      843,
				Width:  834,
				Height: 843,
			},
			Action: linebot.RichMenuAction{
				Type: linebot.RichMenuActionTypeURI,
				URI:  "https://centralfoodwholesale.co.th/",
			},
		},
	}

	areasNonPort := []linebot.AreaDetail{
		{
			Bounds: linebot.RichMenuBounds{
				X:      834,
				Y:      843,
				Width:  834,
				Height: 843,
			},
			Action: linebot.RichMenuAction{
				Type: linebot.RichMenuActionTypeURI,
				URI:  fmt.Sprintf("%s?mode=selectstore", lineLiffURL), // "https://liff.line.me/2006474280-Xx6D8vKL?mode=selectstore",
			},
		},
		{
			Bounds: linebot.RichMenuBounds{
				X:      1668,
				Y:      843,
				Width:  834,
				Height: 843,
			},
			Action: linebot.RichMenuAction{
				Type: linebot.RichMenuActionTypeURI,
				URI:  "https://centralfoodwholesale.co.th/",
			},
		},
	}

	areasPort := []linebot.AreaDetail{
		{
			Bounds: linebot.RichMenuBounds{
				X:      1250,
				Y:      843,
				Width:  1250,
				Height: 843,
			},
			Action: linebot.RichMenuAction{
				Type: linebot.RichMenuActionTypeURI,
				URI:  "https://centralfoodwholesale.co.th/",
			},
		},
	}

	return []linebot.RichMenu{
		// LINKAGE
		{
			Size:        size,
			Selected:    true,
			Name:        linkageName,
			ChatBarText: chatBarText,
			Areas:       areasLinkage,
		},
		// NON_PROD
		{
			Size:        size,
			Selected:    true,
			Name:        nonPortName,
			ChatBarText: chatBarText,
			Areas:       areasNonPort,
		},
		// PROD
		{
			Size:        size,
			Selected:    true,
			Name:        portName,
			ChatBarText: chatBarText,
			Areas:       areasPort,
		},
	}
}
