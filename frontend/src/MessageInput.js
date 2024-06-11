import React, { useRef } from "react";

const MessageInput = ({ addMessage }) => {

	const inputMessage = useRef();

	const sendMessage = () => {
		if (inputMessage.current.value.trim() !== "") {
			addMessage(inputMessage.current.value);
			inputMessage.current.value = "";
		}
	};

	return (
		<div className="message-input">
			<input type="text" ref={inputMessage}/>
			<button onClick={sendMessage}>送信</button>
		</div>
	);
};

export default MessageInput;
