import React from "react";
import Message from "./Message";

const MessageList = ({ messages, loginName }) => {

	return (
		<div className="message-list">
			{messages.map((message, index) => (
				<Message key={index} message={message} loginName={loginName} />
			))}
		</div>
	);
};

export default MessageList;
