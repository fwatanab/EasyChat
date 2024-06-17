import React, { useRef } from "react";

const MessageInput = ({ addMessage, loginName }) => {

	const inputMessage = useRef();

	const handleSendMessage = (e) => {
		e.preventDefault();
		const message = inputMessage.current.value.trim();
		if (message) {
			addMessage(loginName, message);
			inputMessage.current.value = "";
		}
	};

	return (
		<form onSubmit={handleSendMessage} className="message-input">
			<input type="text" ref={inputMessage} />
			<button type="submit" onClick= {handleSendMessage}>送信</button>
		</form>
	);
};

export default MessageInput;
