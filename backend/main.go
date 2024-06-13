package main

import (
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

// 接続中のクライアントを保持するマップ
var clients = make(map[*websocket.Conn]bool)

// メッセージをブロードキャストするためのチャネル
var broadcast = make(chan Message)

// WebSocketで送受信されるメッセージの構造体
type Message struct {
	Message string "json: message"
}

func main() {
	// '/ws'エンドポイントへのリクエストを'handleConnections'で処理
	http.HandleFunc("/ws", handleConnections)
	// 別のゴルーチンで'handleMessages'を実行
	go handleMessages()

	// サーバーをポート8080で開始
	log.Println("Server started on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// 新しいWebSocket接続を処理
func handleConnections(w http.ResponseWriter, r *http.Request) {
	// HTTP接続をWebSocket接続にアップグレード
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	// 接続をclientsマップに追加
	clients[ws] = true

	for {
		var msg Message
		// クライアントからメッセージを受け取る
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		// 受け取ったメッセージをbroadcastチャネルに送信
		broadcast <- msg
	}
}

// ブロードキャストされたメッセージをすべてのクライアントに送信
func handleMessages() {
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
