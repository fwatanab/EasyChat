import React, { useState, useEffect, useRef } from "react";
import MessageList from "./MessageList";
import MessageInput from "./MessageInput";

const ChatWindow = () => {

	const [messages, setMessages] = useState([]);
	const ws = useRef(null);

	useEffect(() => {
		// WebSocket接続を作成し、サーバーのエンドポイントに接続
		ws.current = new WebSocket("ws://localhost:8000/ws");

		ws.current.onopen = () => {
			console.log("Connected to WebSocket server");
		};

		// メッセージを受け取ったときの処理を設定
		ws.current.onmessage = (event) => {
			// 受信したメッセージをJSONにパース
			const message = JSON.parse(event.data);
			// 受信したメッセージをメッセージリストに追加
			setMessages((prevMessages) => [...prevMessages, message.Message]);
		};

		ws.current.onerror = (error) => {
			console.log("WebSocket error: ", error);
		};

		ws.current.onclose = (event) => {
			console.log(`WebSocket closed: ${event.code} (reason: ${event.reason})`);
		};

		// クリーンアップとしてコンポーネントのアンマウント時にWebSocket接続を閉じる
		return () => {
			ws.current.close();
		};
	}, []);

	const addMessage = (message) => {
		// メッセージをサーバーに送信
		ws.current.send(JSON.stringify({ message }));
	};

	return (
		<div className="chat-window">
			<MessageList messages={messages} />
			<MessageInput addMessage={addMessage}/>
		</div>
	);
};

export default ChatWindow;
