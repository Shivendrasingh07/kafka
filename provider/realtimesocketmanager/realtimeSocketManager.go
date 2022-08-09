package realtimesocketmanager

import (
	"context"
	"encoding/json"
	"example.com/m/models"
	"example.com/m/provider"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"net"
	"strconv"
)

// RealtimeHub maintains the set of active clients and broadcasts messages to the clients.
type RealtimeHub struct {
	// registered clients. Each RealtimeClient is identified by a user ID and a
	// uuID. There are a set of user IDs and each user ID has a set of uuIDs.
	// Each uuID identifies a unique RealtimeClient.
	//clients map[int]map[uuid.UUID]*RealtimeClient
	//
	//// register requests from the clients.
	//register chan *RealtimeClient
	//
	//// unregister clients from RealtimeHub
	//unregister chan *RealtimeClient
	//
	//// getClients retrieves a gl
	//getClients    chan int
	//outGetClients chan getClientsResp
	Messenger provider.KafkaProvider

	//Done chan bool
}

//type getClientsResp struct {
//	clients map[uuid.UUID]*RealtimeClient
//	err     error
//}

func NewRealtimeHub(messenger provider.KafkaProvider) provider.WebSocketHubProvider {
	return &RealtimeHub{
		//register:   make(chan *RealtimeClient, 1),
		//unregister: make(chan *RealtimeClient, 1),
		//
		//getClients:    make(chan int, 1),
		//outGetClients: make(chan getClientsResp, 1),
		//
		//Done:      make(chan bool, 1),
		//clients:   make(map[int]map[uuid.UUID]*RealtimeClient),
		Messenger: messenger,
	}
}

//func (h *RealtimeHub) Run() {
//	go h.SubscribeAllPartitions()
//	for {
//		select {
//		case <-h.Done:
//			h.Stop()
//			return
//		case client := <-h.register:
//			if _, ok := h.clients[client.userContext.ID]; !ok {
//				h.clients[client.userContext.ID] = make(map[uuid.UUID]*RealtimeClient)
//			}
//			h.clients[client.userContext.ID][client.uuID] = client
//
//		case client := <-h.unregister:
//			delete(h.clients[client.userContext.ID], client.uuID)
//
//		case userID := <-h.getClients:
//			clients, ok := h.clients[userID]
//			if !ok {
//				h.outGetClients <- getClientsResp{
//					clients: nil,
//					err:     fmt.Errorf("no Clients are associated with userID %v", userID),
//				}
//				break
//			}
//			h.outGetClients <- getClientsResp{
//				clients: clients,
//				err:     nil,
//			}
//		}
//	}
//}

func (h *RealtimeHub) Get() interface{} {
	return h
}

//func (h *RealtimeHub) Stop() {
//	for _, userClients := range h.clients {
//		for _, client := range userClients {
//			close(client.send)
//		}
//	}
//	h.Messenger.Close()
//}

func (h *RealtimeHub) Run() {
	go h.SubscribeAllPartitions()
}

func (h *RealtimeHub) SubscribeAllPartitions() {
	kafkaHost := "localhost:9092"

	conn, err := kafka.Dial("tcp", kafkaHost)
	if err != nil {
		logrus.Errorf("SubscribeAllPartitions: error kafka dialing %v", err)
		return
	}

	var controllerConn *kafka.Conn
	defer func() {
		_ = conn.Close()
		_ = controllerConn.Close()
	}()

	controller, err := conn.Controller()
	if err != nil {
		logrus.Errorf("startNotifier: error kafka dialing controller %v", err)
		return
	}

	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		panic(err.Error())
	}

	// This will have no effect if topic already existed
	err = controllerConn.CreateTopics(kafka.TopicConfig{
		Topic:             "testing101",
		NumPartitions:     5,
		ReplicationFactor: 1,
	})
	if err != nil {
		logrus.Errorf("SubscribeAllPartitions: error creating topic %v", err)
		return
	}

	h.Subscribe()
}

func (h *RealtimeHub) Subscribe() {
	kafkaHost := "localhost:9092"

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{kafkaHost},
		Topic:   "testing101",
		GroupID: "g1",
	})
	_ = r.SetOffset(kafka.LastOffset)

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}
		//logrus.Infof("RealtimeHub: time %v read message at offset %d: %s = %s\n", time.Now(), m.Offset, string(m.Key), string(m.Value))

		var message models.Message
		if err = json.Unmarshal(m.Value, &message); err != nil {
			logrus.Errorf("Failed to Unmarshal message from topic: %q error: %v", m.Topic, err)
			return
		}

		//messageBytes, err := json.Marshal(message.Message)
		//if err != nil {
		//	logrus.Errorf("error Marshaling chatMessage: %q", err.Error())
		//	return
		//}

		logrus.Println(message)
		fmt.Println(message)

		//for i := range message.ToUserID {
		//	h.getClients <- message.ToUserID[i]
		//
		//	out, ok := <-h.outGetClients
		//	if !ok || out.err != nil {
		//		continue
		//	}
		//
		//	for i := range out.clients {
		//		out.clients[i].send <- messageBytes
		//	}
		//}
	}
}

//func (h *RealtimeHub) Messengers() providers.KafkaProvider {
//	return h.Messenger
//}

//func (h *RealtimeHub) SendOnlineStatusReport() {
//	for {
//		for _, clients := range h.clients {
//			for _, c := range clients {
//				c.SendOnlineStatusReportFromClient(false)
//			}
//		}
//		time.Sleep(10 * time.Second)
//	}
//}
