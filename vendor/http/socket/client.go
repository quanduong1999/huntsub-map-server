package socket

import (
	"http/web"
	"huntsub/huntsub-map-server/event"
	"huntsub/huntsub-map-server/o/org/user"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var obj = event.ObjectEventSource

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Connection is an middleman between the websocket Connection and the hub.
type Connection struct {
	// The websocket Connection.
	Ws *websocket.Conn

	// Buffered channel of outbound messages.
	Send chan []byte
}

// readPump pumps messages from the websocket Connection to the hub.
func (s Subscription) readPump(h *Hub) {
	c := s.Conn
	defer func() {
		h.Unregister <- s
		c.Ws.Close()
	}()
	c.Ws.SetReadLimit(maxMessageSize)
	c.Ws.SetReadDeadline(time.Now().Add(pongWait))
	c.Ws.SetPongHandler(func(string) error { c.Ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, msg, err := c.Ws.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}
		m := Message{msg, s.Room}
		h.Broadcast <- m
	}
}

// write writes a message with the given message type and payload.
func (c *Connection) write(mt int, payload []byte) error {
	c.Ws.SetWriteDeadline(time.Now().Add(writeWait))
	return c.Ws.WriteMessage(mt, payload)
}

// writePump pumps messages from the hub to the websocket Connection.
func (s *Subscription) writePump(u *user.User) {
	c := s.Conn
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Ws.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				println("CLOSE")
				var newUser = &user.User{}
				newUser.StatusActive.Online = false
				newUser.StatusActive.RoomID = ""
				newUser.StatusActive.TimeOut = time.Now().Local().Unix()
				res, err := u.Update(newUser)
				web.AssertNil(err)
				obj.EmitUpdate(res)
				return
			}
			if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func ServeWs(h *Hub, w http.ResponseWriter, r *http.Request, roomId string, u *user.User) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
	c := &Connection{Send: make(chan []byte, 256), Ws: ws}
	s := Subscription{c, roomId}
	h.Register <- s
	go s.writePump(u)
	go s.readPump(h)
}
