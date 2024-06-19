package handlers

import (
	"easychat/database"
	"easychat/models"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// HTTP接続をWebSocket接続にアップグレードするための設定
var upgrader = websocket.Upgrader{
	// バッファサイズを設定
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// オリジンチェックを無効にします（どのオリジンからでも接続を受け入れます）
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	conn                *websocket.Conn
	initialMessagesSent bool
}

// 接続中のクライアントを保持するマップ
var clients = make(map[*websocket.Conn]*Client)

// メッセージをブロードキャストするためのチャネル
var broadcast = make(chan models.Message)

// 新しいWebSocket接続を処理
func HandleConnections(w http.ResponseWriter, r *http.Request) {
	// HTTP接続をWebSocket接続にアップグレード
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	// 接続をclientsマップに追加
	client := &Client{conn: ws, initialMessagesSent: false}
	clients[ws] = client

	if !client.initialMessagesSent {
		initialMessages, err := database.GetMessages()
		if err != nil {
			log.Printf("Failed to get initial messages: %v\n", err)
		} else {
			err := ws.WriteJSON(&initialMessages)
			if err != nil {
				log.Printf("Failed to send initial messages: %v\n", err)
			} else {
				client.initialMessagesSent = true
			}
		}
	}

	for {
		var msg models.Message
		// クライアントからメッセージを受け取る
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}

		// デバッグ用プリント
		//		log.Printf("Received message: %+v", msg)
		//		log.Printf("name: %s message: %s\n", msg.Name, msg.InputMessage)

		err = database.AddUser(msg.Name)
		if err != nil {
			log.Printf("Failed to add user: %v", err)
			continue
		}

		userID, err := database.GetUserID(msg.Name)
		if err != nil {
			log.Printf("Failed to get user ID: %v", err)
		}

		err = database.AddMessage(userID, msg.InputMessage)
		if err != nil {
			log.Printf("Failed to add message: %v", err)
		}

		time, err := database.GetTime(userID)
		if err != nil {
			log.Printf("Failed to get time: %v", err)
		}
		msg.Timestamp = time

		// 受け取ったメッセージをbroadcastチャネルに送信
		broadcast <- msg
	}
	delete(clients, ws)
}

// ブロードキャストされたメッセージをすべてのクライアントに送信
func HandleMessages() {
	for {
		// チャネルからメッセージを受け取る
		msg := <-broadcast

		for client := range clients {
			// クライアントにメッセージを送信
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				// エラーが発生した場合、クライアントを閉じる
				client.Close()
				// クライアントをclientsマップから削除
				delete(clients, client)
			}
		}
	}
}
