# シンプルなチャットアプリケーション

## 概要

このプロジェクトは、WebSocketを使用したサーバーとクライアント間の通信、MySQLを使用したデータベースを学ぶためのチュートリアルです。  
任意のユーザー名でログインすることで、チャット機能を使うことができます。  
リアルタイムでメッセージが表示されるシンプルなチャットアプリケーションです。(ユーザー名、メッセージ、送信時刻の表示)  
７日間経過したメッセージはデータベースから削除されます。  

## データベース / フレームワーク / 言語

MySQL / React / HTML, CSS, JavaScript

## インストール方法

1. リポジトリをクローンします：

   ```bash
   git clone https://github.com/fwatanab/EasyChat.git

2. 自身のMySQLデータベースのユーザー名、パスワードに置き換える：  
main.go: 13行目
   ```bash
   database.InitDB("user:password@tcp(127.0.0.1:3306)/")

3. バックエンドサーバーを起動します：

   ```bash
   cd EasyChat/backend
   go mod tidy
   go run main.go
   
4. フロントエンドサーバーを起動します：

   ```bash
   cd EasyChat/frontend
   npm install
   npm start

5. ブラウザを開いて、http://localhost:3000 にアクセスしてチャットアプリケーションを使用します。
