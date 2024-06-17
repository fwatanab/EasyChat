import React, { useState, useEffect, useRef } from "react";
import { useLocation } from "react-router-dom";
import MessageList from "./MessageList";
import MessageInput from "./MessageInput";

const ChatWindow = () => {

	const [messages, setMessages] = useState([]);
	const ws = useRef(null);

	// 現在のURLを取得
	const location = useLocation();
	// URLのクエリパラメーターを解析
	const queryParams = new URLSearchParams(location.search);
	// クエリパラメーターからloginNameを取得
	const loginName = queryParams.get("name");

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

			// デバッグ用ログ
//			console.log(message);

			// 受信したメッセージをメッセージリストに追加
			setMessages((prevMessages) => [...prevMessages, message]);
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

	const addMessage = (name, inputMessage) => {
		const message = {name, inputMessage};

		// デバッグ用ログ
//		console.log(message);

		// メッセージをサーバーに送信
		ws.current.send(JSON.stringify(message));
	};

	return (
		<div className="chat-window">
			<p>{loginName}</p>
			<MessageList messages={messages} />
			<MessageInput addMessage={addMessage} loginName={loginName}/>
		</div>
	);
};

export default ChatWindow;
