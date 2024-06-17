import React from "react";

const Message = ({ message }) => {

	return (
		<div className="message">
			<li>
				<strong>{message.name}: </strong>{message.inputMessage}
			</li>
		</div>
	);
};

export default Message;

