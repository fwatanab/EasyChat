import React from "react";

const Message = ({ message, loginName }) => {
	return (
		<div className="message">
			<li>
				<strong>{loginName}: </strong>{message}
			</li>
		</div>
	);
};

export default Message;

