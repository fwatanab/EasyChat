import React from "react";

const Message = ({ message }) => {
	if (!message) {
		return null;
	}

	function formatTimestamp(timestamp) {
	const currentDate = new Date();
	const messageDate = new Date(timestamp);

	const timeDiff = currentDate - messageDate;
	const hoursDiff = Math.floor(timeDiff / (1000 * 60 * 60));

	if (hoursDiff < 24) {
		const hours = messageDate.getHours().toString().padStart(2, "0");
		const minutes = messageDate.getMinutes().toString().padStart(2, "0");
		return `${hours}:${minutes}`;
	} else {
		const daysDiff = Math.floor(hoursDiff / 24);
		return `${daysDiff} day(s)`;
	}
}

	return (
		<div className="message">
			<li>
				<strong>{message.name}: </strong>{message.inputMessage}
				<span className="timestamp">{formatTimestamp(message.timestamp)}</span>
			</li>
		</div>
	);
};

export default Message;

