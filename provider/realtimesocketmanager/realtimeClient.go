package realtimesocketmanager

//
//import (
//	"encoding/json"
//	"time"
//
//	"github.com/RemoteState/connect-up/models"
//	"github.com/google/uuid"
//	"github.com/gorilla/websocket"
//	"github.com/sirupsen/logrus"
//)
//
//const (
//
//	// Time allowed to write a message to the peer.
//	writeWait = 10 * time.Second
//
//	// Time allowed to read the next pong message from the peer.
//	pongWait = 60 * time.Second
//
//	// Send pings to peer with this period. Must be less than pongWait.
//	pingPeriod = (pongWait * 9) / 10
//)
//
//// RealtimeClient is a middleman between the websocket connection and the hub.
//type RealtimeClient struct {
//	hub *RealtimeHub
//
//	// The websocket connection.
//	conn *websocket.Conn
//
//	// Buffered channel of outbound messages.
//	send chan []byte
//
//	// userContext is the user details for which the RealtimeClient belongs to.
//	userContext *models.UserContext
//
//	// uuID is a unique identifier for differentiating each RealtimeClient
//	// connection.
//	uuID uuid.UUID
//
//	// connectionUnixNano is the unix time in nano when connection established
//	connectionUnixNano int64
//
//	UsersSubscribedForOnlineStatus map[string]models.UserOnlineStatus
//}
//
//func NewRealtimeClient(hub *RealtimeHub, conn *websocket.Conn, uc *models.UserContext) *RealtimeClient {
//	uuID, err := uuid.NewUUID()
//	if err != nil {
//		logrus.Error(err)
//	}
//	return &RealtimeClient{
//		hub:                            hub,
//		conn:                           conn,
//		send:                           make(chan []byte),
//		userContext:                    uc,
//		uuID:                           uuID,
//		connectionUnixNano:             time.Now().UnixNano(),
//		UsersSubscribedForOnlineStatus: make(map[string]models.UserOnlineStatus),
//	}
//}
//
//func (c *RealtimeClient) Register() {
//	c.hub.register <- c
//}
//
//func (c *RealtimeClient) close() {
//	close(c.send)
//	if err := c.conn.Close(); err != nil {
//		logrus.Error("closing client", err.Error())
//	}
//}
//
//// WritePump pumps messages from the hub to the websocket connection.
////
//// A goroutine running writePump is started for each connection. The
//// application ensures that there is at most one writer to a connection by
//// executing all writes from this goroutine.
//func (c *RealtimeClient) WritePump() {
//	ticker := time.NewTicker(pingPeriod)
//	defer func() {
//		ticker.Stop()
//		if err := c.conn.Close(); err != nil {
//			logrus.Error(err)
//		}
//		c.hub.unregister <- c
//	}()
//	for {
//		select {
//		case message, ok := <-c.send:
//			if err := c.conn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
//				logrus.Error(err)
//			}
//			if !ok {
//				// The hub closed the channel.
//				if err := c.conn.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
//					logrus.Error(err)
//				}
//				return
//			}
//
//			if err := c.conn.WriteMessage(websocket.BinaryMessage, message); err != nil {
//				logrus.Error(err)
//				return
//			}
//
//		case <-ticker.C:
//			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
//			if err := c.conn.WriteMessage(websocket.PingMessage, []byte("ping")); err != nil {
//				return
//			}
//		}
//	}
//}
//
//// ReadPump pumps messages from the websocket connection to the hub.
//func (c *RealtimeClient) ReadPump() {
//	defer func() {
//		c.hub.unregister <- c
//		_ = c.conn.Close()
//	}()
//
//	c.conn.SetReadLimit(models.MB)
//	_ = c.conn.SetReadDeadline(time.Now().Add(pongWait))
//	c.conn.SetPongHandler(func(string) error { _ = c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
//	for {
//		_, msg, err := c.conn.ReadMessage()
//		if err != nil {
//			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
//				logrus.Infof("error: %v", err)
//			}
//			break
//		}
//
//		var message models.Message
//		if err := json.Unmarshal(msg, &message); err != nil {
//			logrus.Errorf("Failed to Unmarshal message error: %v", err)
//			continue
//		}
//
//		switch message.Type {
//		case models.WSMessageTypeOnlineStatus:
//			c.getOnlineStatus(message)
//		case models.WSMessageTypeReadReceipt:
//			c.processReadReceipt(message)
//		case models.WSMessageTypeReceivedReceipt:
//			c.processReceivedReceipt(message)
//		case models.WSMessageTypePing:
//			c.processPing()
//		case models.WSMessageTypeChatGroupConfig:
//			c.setChatGroupConfig(message)
//		case models.WSMessageSubscribeUserOnlineStatus:
//			c.subscribeUserStatus(message)
//		case models.WsMessageUnsubscribeUserOnlineStatus:
//			c.unsubscribeUserStatus(message)
//		case models.WSMessageTypeMuteAudio:
//			c.muteAudioForUser(message)
//		case models.WSMessageTypeUnmuteAudio:
//			c.unmuteAudioForUser(message)
//		case models.WSMessageTypeDisableVideo:
//			c.disableVideoForUser(message)
//		case models.WSMessageTypeEnableVideo:
//			c.enableVideoForUser(message)
//		case models.WSMessageTypeEndCall:
//			c.endCallRequest(message)
//		case models.WSMessageTypeDeclineCall:
//			c.declineCallRequest(message)
//		default:
//			logrus.Errorf("invalid ws message type")
//		}
//	}
//}
