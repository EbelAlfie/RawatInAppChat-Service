package main

import (
	"chat_service/domain/model"
	"chat_service/router"
	"chat_service/utils"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func main() {
	server := gin.Default()
	publicRoute := server.Group("/main")
	{
		server.Group("/user", router.UserRoute(publicRoute))
	}

	upgrader := websocket.Upgrader{}
	server.GET("/ws", func(ctx *gin.Context) {
		fmt.Println("Called")
		con, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			return 
		}

		for {
			_, msgByte, err := con.ReadMessage()
			if err != nil {
				fmt.Println("Error " + err.Error())
				continue
			}
			msg := string(msgByte)
			fmt.Println(msg)

			var chat model.RealtimeChatEvent
			err = json.Unmarshal(msgByte, &chat)
			if err != nil {
				fmt.Println("Unmarshal error " + err.Error())
				continue
			}

			var newMsg model.RealtimeChatEvent
			if (chat.User.Id == "responden") {
				newMsg = model.RealtimeChatEvent {
					Id: utils.RandStringBytes(10),
					Text: chat.Text,
					User: chat.User,
				}
			} else {
				newMsg = model.RealtimeChatEvent {
					Id: chat.Id,
					Text: chat.Text,
					User: chat.User,
				}
			}
			jsonDat, err := json.Marshal(newMsg)
			if err != nil {
				fmt.Println("Encode error " + err.Error())
				continue
			}
			err = con.WriteMessage(websocket.TextMessage, jsonDat)
			fmt.Println("Send message ")
			if err != nil {
				fmt.Println("Error writing message " + err.Error())
			}
		}
	})

	server.Run("127.0.0.1:8003")
}
