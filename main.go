package main

import (
	"chat_service/domain"
	"chat_service/router"
	"chat_service/utils"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var sessions *domain.SessionSocket //[]domain.SessionSocket{}

func main() {
	server := gin.Default()
	publicRoute := server.Group("/main")
	{
		server.Group("/user", router.UserRoute(publicRoute))
	}

	upgrader := websocket.Upgrader{}
	server.GET("/ws", func(ctx *gin.Context) {
		id := ctx.Request.URL.Query().Get("userId")
		con, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			return 
		}

		socket := getSocket(con, id)

		for {
			_, msgByte, err := con.ReadMessage()
			if err != nil {
				fmt.Println("Error " + err.Error())
				continue
			}
			msg := string(msgByte)
			fmt.Println(msg)

			var chat domain.RealtimeChatEvent
			err = json.Unmarshal(msgByte, &chat)
			if err != nil {
				fmt.Println("Unmarshal error " + err.Error())
				continue
			}

			var newMsg domain.RealtimeChatEvent
			if (chat.User.Id == "responden") {
				newMsg = domain.RealtimeChatEvent {
					Id: utils.RandStringBytes(10),
					Text: chat.Text,
					User: chat.User,
				}
			} else {
				newMsg = domain.RealtimeChatEvent {
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
			
			err = socket.Ws.WriteMessage(websocket.TextMessage, jsonDat)
			fmt.Println("Send message ")
			if err != nil {
				fmt.Println("Error writing message " + err.Error())
			}
		}
	})

	server.Run("0.0.0.0:8003")
}

func getSocket(con *websocket.Conn, sessionId string) *domain.SessionSocket {
	var socket *domain.SessionSocket

	if sessions == nil || sessions.Sid == "" {
		socket = &domain.SessionSocket{
					Sid: sessionId,
					Ws: con,
				}
		sessions = socket
	} else {
		socket = sessions
	}
	// if i := slices.IndexFunc(sessions, func(s domain.SessionSocket) bool { return s.Sid == sessionId }); i != nil || i {
	// 	socket = sessions[i]
	// } else {
	// 	socket = domain.SessionSocket{
	// 		Sid: sessionId,
	// 		Ws: *con,
	// 	}
	// 	sessions = append(sessions, socket)
	// }
	return socket
}