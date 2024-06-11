import React, { useState } from "react";
import MessageList from "./MessageList";
import MessageInput from "./MessageInput";

const ChatWindow = () => {

	const [messages, setMessages] = useState([]);

	const addMessage = (message) => {
		return setMessages([...messages, message]);
	};

	return (
		<div className="chat-window">
			<MessageList messages={messages} />
			<MessageInput addMessage={addMessage}/>
		</div>
	);
};

export default ChatWindow;
