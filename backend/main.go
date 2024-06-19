package main

import (
	"easychat/database"
	"easychat/handlers"
	"log"
	"net/http"
	"time"
)

func main() {
	// // データベース接続を初期化
	database.InitDB("root:Ftsg0601_348441@tcp(127.0.0.1:3306)/")

	// 定期的にメンテナンス処理を実行
	go func() {
		for {
			// // 指定した日数経過したメッセージを削除
			err := database.DeleteOldMessages(1)
			if err != nil {
				log.Printf("Failed to delete old messages: %v", err)
			}
			// 1日ごとに実行
			time.Sleep(24 * time.Hour)
		}
	}()

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
