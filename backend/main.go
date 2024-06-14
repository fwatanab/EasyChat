package main

import (
	"log"
	"net/http"
	"easychat/handlers"
)

func main() {
	// '/ws'エンドポイントへのリクエストを'handleConnections'で処理
	http.HandleFunc("/ws", handlers.HandleConnections)
	// 別のゴルーチンで'handleMessages'を実行
	go handlers.HandleMessages()

	// サーバーをポート8080で開始
	log.Println("Server started on :8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
