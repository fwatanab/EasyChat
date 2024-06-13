# シンプルなチャットアプリケーション

## 概要

このプロジェクトは、WebSocketを使用したサーバーとクライアント間の通信を学ぶためのチュートリアルです。  
リアルタイムでメッセージが表示されるシンプルなチャットアプリケーションを開発します。  
ただし、データベースは実装されていないため、ページのリロードを行うとメッセージがリセットされます。

## フレームワーク / 言語

React / HTML, JavaScript

## インストール方法

1. リポジトリをクローンします：

   ```bash
   git clone https://github.com/fwatanab/EasyChat.git

2. バックエンドサーバーを起動します：

   ```bash
   cd EasyChat/backend
   go mod tidy
   go run main.go
   
3. フロントエンドサーバーを起動します：

   ```bash
   cd EasyChat/frontend
   npm install
   npm start

4. ブラウザを開いて、http://localhost:3000 にアクセスしてチャットアプリケーションを使用します。
