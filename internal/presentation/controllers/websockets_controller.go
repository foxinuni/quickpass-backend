package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/foxinuni/quickpass-backend/internal/domain/entities"
	"github.com/foxinuni/quickpass-backend/internal/domain/services"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true // You may want to add origin checking for security
    },
}

type MessageStruct struct{
	log *entities.LogHistory
	id int
}

var clients = make(map[*websocket.Conn]int)
var eventsChannel = make(chan MessageStruct)

type WebSocketsController struct {
	actionsService services.ActionsService
}

func NewWebSocketsController(actionsService services.ActionsService) *WebSocketsController {
	return &WebSocketsController{
		actionsService: actionsService,
	}
}

func (sc *WebSocketsController) NewEventLog(occasionId int){
	log, eventId, _, err :=sc.actionsService.GetLastLog(occasionId)
	if err != nil{
		return
	}
	if eventId != nil{
		eventsChannel <- MessageStruct{
			log: log,
			id: *eventId,
		}
	}
}

func (sc *WebSocketsController) EventsWebSocketHanlder(c echo.Context) error {
	eventId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
        return err
    }
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
    if err != nil {
        return err
    }
    defer ws.Close()
	
    clients[ws] = eventId
    for {
        _, _, err := ws.ReadMessage()
        if err != nil {
            delete(clients, ws)
            break
        }
    }
    return nil
}

func EventBroadcaster() {
    for {
        msg := <-eventsChannel
        for client, eventId := range clients {
			if eventId != msg.id{
				continue
			}
			message, _ := json.Marshal(msg.log)
            err := client.WriteMessage(websocket.TextMessage, []byte(message))
            if err != nil {
                client.Close()
                delete(clients, client)
            }
        }
    }
}
