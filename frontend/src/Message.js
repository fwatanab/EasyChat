import React from "react";

const Message = ({ message }) => {
	return (
		<div className="message">
			<li>{message}</li>
		</div>
	);
};

export default Message;

